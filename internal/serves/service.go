package serves

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/rs/zerolog/log"
	"github.com/sourcegraph/jsonrpc2"
)

type File struct {
	LanguageID string
	Text       string
}

func (that *Server) Initialize(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ParamInitialize) (*lsp.InitializeResult, error) {
	var dbc config.DBConfig
	b, err := json.Marshal(params.InitializationOptions)
	if nil != err {
		return nil, err
	}
	if err = json.Unmarshal(b, &dbc); nil != err {
		return nil, err
	}

	result := &lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync:   lsp.Full,
			HoverProvider:      &lsp.Or_ServerCapabilities_hoverProvider{Value: true},
			CodeActionProvider: true,
			CompletionProvider: &lsp.CompletionOptions{
				TriggerCharacters: []string{"(", "."},
			},
			SignatureHelpProvider: &lsp.SignatureHelpOptions{
				TriggerCharacters:   []string{"(", ","},
				RetriggerCharacters: []string{"(", ","},
				WorkDoneProgressOptions: lsp.WorkDoneProgressOptions{
					WorkDoneProgress: false,
				},
			},
			DefinitionProvider:              &lsp.Or_ServerCapabilities_definitionProvider{Value: true},
			DocumentFormattingProvider:      &lsp.Or_ServerCapabilities_documentFormattingProvider{Value: true},
			DocumentRangeFormattingProvider: &lsp.Or_ServerCapabilities_documentRangeFormattingProvider{Value: true},
			RenameProvider:                  true,
		},
	}

	that.initOptionDBConfig = &dbc

	// Initialize database database connection
	// NOTE: If no connection is found at this point, it is possible that the connection settings are sent to workspace config, so don't make an error
	messenger := lsp.NewLspMessenger(conn)
	if err = that.reconnectionDB(ctx); err != nil {
		if !errors.Is(ErrNoConnection, err) {
			if err = messenger.ShowInfo(ctx, err.Error()); err != nil {
				log.Error().Msgf("send info, %s", err.Error())
				return nil, err
			}
		} else {
			log.Error().Msgf("send err, %s", err.Error())
			if err = messenger.ShowError(ctx, err.Error()); err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}

func (that *Server) Initialized(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.InitializedParams) error {
	return nil
}

func (that *Server) Shutdown(ctx context.Context, conn *jsonrpc2.Conn) error {
	if that.dbConn != nil {
		return that.dbConn.Close()
	}
	return nil
}

func (that *Server) Exit(ctx context.Context, conn *jsonrpc2.Conn) error {
	if that.dbConn != nil {
		that.dbConn.Close()
	}
	return that.Stop()
}

func (that *Server) DidOpen(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DidOpenTextDocumentParams) error {
	if err := that.openFile(params.TextDocument.URI, params.TextDocument.LanguageID); err != nil {
		return err
	}
	if err := that.updateFile(params.TextDocument.URI, params.TextDocument.Text); err != nil {
		return err
	}
	return nil
}

func (that *Server) DidChange(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DidChangeTextDocumentParams) error {
	return that.updateFile(params.TextDocument.URI, params.ContentChanges[0].Text)
}

func (that *Server) DidSave(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DidSaveTextDocumentParams) error {
	if params.Text != "" {
		return that.updateFile(params.TextDocument.URI, params.Text)
	} else {
		return that.saveFile(params.TextDocument.URI)
	}
}

func (that *Server) DidClose(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DidCloseTextDocumentParams) error {
	return that.closeFile(params.TextDocument.URI)
}

func (that *Server) openFile(uri lsp.DocumentURI, languageID string) error {
	f := &File{
		Text:       "",
		LanguageID: languageID,
	}
	that.files[uri] = f
	return nil
}

func (that *Server) closeFile(uri lsp.DocumentURI) error {
	delete(that.files, uri)
	return nil
}

func (that *Server) updateFile(uri lsp.DocumentURI, text string) error {
	f, ok := that.files[uri]
	if !ok {
		return fmt.Errorf("document not found: %v", uri)
	}
	f.Text = text
	return nil
}

func (that *Server) saveFile(uri lsp.DocumentURI) error {
	return nil
}

func (that *Server) DidChangeConfiguration(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DidChangeConfigurationParams) error {
	// Update changed configration
	// that.WSCfg = params.Settings.SQLS

	// Skip database connection
	if that.dbConn != nil {
		return nil
	}

	// Initialize database database connection
	messenger := lsp.NewLspMessenger(conn)
	if err := that.reconnectionDB(ctx); err != nil {
		if !errors.Is(ErrNoConnection, err) {
			if err = messenger.ShowInfo(ctx, err.Error()); err != nil {
				log.Error().Msgf("send info, %s", err.Error())
				return err
			}
		} else {
			log.Error().Msgf("send err, %s", err.Error())
			if err = messenger.ShowError(ctx, err.Error()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (that *Server) reconnectionDB(ctx context.Context) error {
	if err := that.dbConn.Close(); err != nil {
		return err
	}

	dbConn, err := that.newDBConnection(ctx)
	if err != nil {
		return err
	}
	that.dbConn = dbConn
	dbRepo, err := that.newDBRepository(ctx)
	if err != nil {
		return err
	}
	if err := that.worker.ReCache(ctx, dbRepo); err != nil {
		return err
	}
	return nil
}

func (that *Server) newDBConnection(ctx context.Context) (*database.DBConnection, error) {
	// Get the most preferred DB connection settings
	connCfg := that.topConnection()
	if connCfg == nil {
		return nil, ErrNoConnection
	}
	if that.curConnectionIndex != 0 {
		connCfg = that.getConnection(that.curConnectionIndex)
	}
	if connCfg == nil {
		return nil, fmt.Errorf("not found database connection config, index %d", that.curConnectionIndex+1)
	}
	if that.curDBName != "" {
		connCfg.DBName = that.curDBName
	}
	that.curDBCfg = connCfg

	// Connect database
	conn, err := database.Open(connCfg)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (that *Server) newDBRepository(ctx context.Context) (database.DBRepository, error) {
	repo, err := database.CreateRepository(that.curDBCfg.Driver, that.dbConn.Conn)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (that *Server) topConnection() *config.DBConfig {
	// if the init config is set, ignore all other connection configs
	if that.initOptionDBConfig != nil {
		return that.initOptionDBConfig
	}

	cfg := that.getConfig()
	if cfg == nil || len(cfg.Connections) == 0 {
		return nil
	}
	return cfg.Connections[0]
}

func (that *Server) getConnection(index int) *config.DBConfig {
	cfg := that.getConfig()
	if cfg == nil || (index < 0 && len(cfg.Connections) <= index) {
		return nil
	}
	return cfg.Connections[index]
}

func (that *Server) getConfig() *config.Config {
	var cfg *config.Config
	switch {
	case validConfig(that.SpecificFileCfg):
		cfg = that.SpecificFileCfg
	case validConfig(that.WSCfg):
		cfg = that.WSCfg
	case validConfig(that.DefaultFileCfg):
		cfg = that.DefaultFileCfg
	default:
		cfg = config.NewConfig()
	}
	return cfg
}

func validConfig(cfg *config.Config) bool {
	// if cfg != nil && len(cfg.Connections) > 0 {
	if cfg != nil {
		return true
	}
	return false
}
