package serves

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/command"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/log"
	"github.com/ducesoft/ulsp/lsp"
)

type File struct {
	LanguageID string
	Text       string
}

func (that *File) LID() string {
	return that.LanguageID
}

func (that *File) LText() string {
	return that.Text
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
			TextDocumentSync: lsp.Full,
			CompletionProvider: &lsp.CompletionOptions{
				TriggerCharacters: []string{"(", "."},
			},
			HoverProvider: &lsp.Or_ServerCapabilities_hoverProvider{Value: true},
			SignatureHelpProvider: &lsp.SignatureHelpOptions{
				TriggerCharacters:   []string{"(", ","},
				RetriggerCharacters: []string{"(", ","},
				WorkDoneProgressOptions: lsp.WorkDoneProgressOptions{
					WorkDoneProgress: false,
				},
			},
			DefinitionProvider:              &lsp.Or_ServerCapabilities_definitionProvider{Value: true},
			DocumentHighlightProvider:       &lsp.Or_ServerCapabilities_documentHighlightProvider{Value: true},
			DocumentSymbolProvider:          &lsp.Or_ServerCapabilities_documentSymbolProvider{Value: true},
			CodeActionProvider:              true,
			ColorProvider:                   &lsp.Or_ServerCapabilities_colorProvider{Value: true},
			DocumentFormattingProvider:      &lsp.Or_ServerCapabilities_documentFormattingProvider{Value: true},
			DocumentRangeFormattingProvider: &lsp.Or_ServerCapabilities_documentRangeFormattingProvider{Value: true},
			RenameProvider:                  true,
			FoldingRangeProvider:            &lsp.Or_ServerCapabilities_foldingRangeProvider{Value: true},
			SelectionRangeProvider:          &lsp.Or_ServerCapabilities_selectionRangeProvider{Value: true},
			ExecuteCommandProvider: &lsp.ExecuteCommandOptions{
				Commands:                command.Commands(),
				WorkDoneProgressOptions: lsp.WorkDoneProgressOptions{},
			},
		},
		ServerInfo: &lsp.PServerInfoMsg_initialize{
			Name:    "LSP",
			Version: "0.0.1",
		},
	}

	that.initOptionDBConfig = &dbc

	// Initialize database database connection
	// NOTE: If no connection is found at this point, it is possible that the connection settings are sent to workspace config, so don't make an error
	messenger := lsp.NewLspMessenger(conn)
	if err = that.Reconnection(ctx); err != nil {
		if !errors.Is(ErrNoConnection, err) {
			if err = messenger.ShowInfo(ctx, err.Error()); err != nil {
				log.Error(ctx, "send info, %s", err.Error())
				return nil, err
			}
		} else {
			log.Error(ctx, "send err, %s", err.Error())
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
	if err := that.Reconnection(ctx); err != nil {
		if !errors.Is(ErrNoConnection, err) {
			if err = messenger.ShowInfo(ctx, err.Error()); err != nil {
				log.Error(ctx, "send info, %s", err.Error())
				return err
			}
		} else {
			log.Error(ctx, "send err, %s", err.Error())
			if err = messenger.ShowError(ctx, err.Error()); err != nil {
				return err
			}
		}
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

func (that *Server) topConnection() *config.DBConfig {
	// if the init config is set, ignore all other connection configs
	if that.initOptionDBConfig != nil {
		return that.initOptionDBConfig
	}

	cfg := that.Config()
	if cfg == nil || len(cfg.Connections) == 0 {
		return nil
	}
	return cfg.Connections[0]
}

func (that *Server) getConnection(index int) *config.DBConfig {
	cfg := that.Config()
	if cfg == nil || (index < 0 && len(cfg.Connections) <= index) {
		return nil
	}
	return cfg.Connections[index]
}

func validConfig(cfg *config.Config) bool {
	// if cfg != nil && len(cfg.Connections) > 0 {
	if cfg != nil {
		return true
	}
	return false
}
