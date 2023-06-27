package serves

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/command"
	"github.com/ducesoft/ulsp/lsp"
	"strconv"
	"strings"

	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/jsonrpc2"
)

func (that *Server) CodeAction(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.CodeActionParams) ([]lsp.CodeAction, error) {
	return nil, nil
}

func (that *Server) ExecuteCommand(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (interface{}, error) {
	if c := command.Load(params.Command); nil != c {
		return c.Exec(ctx, conn, params, that)
	}
	return nil, fmt.Errorf("unsupported command: %v", params.Command)
}

func extractRangeText(text string, startLine, startChar, endLine, endChar int) string {
	writer := bytes.NewBufferString("")
	scanner := bufio.NewScanner(strings.NewReader(text))

	i := 0
	for scanner.Scan() {
		t := scanner.Text()
		if i >= startLine && i <= endLine {
			st, en := 0, len(t)

			if i == startLine {
				st = startChar
			}
			if i == endLine {
				en = endChar
			}

			writer.Write([]byte(t[st:en]))
			if i != endLine {
				writer.Write([]byte("\n"))
			}
		}
		i++
	}
	return writer.String()
}

func (that *Server) Conn() *database.DBConnection {
	return that.dbConn
}

func (that *Server) Open(uri lsp.DocumentURI) command.File {
	return that.files[uri]
}

func (that *Server) Repository(ctx context.Context) (database.DBRepository, error) {
	repo, err := database.CreateRepository(that.curDBCfg.Driver, that.dbConn.Conn)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (that *Server) Config() *config.Config {
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

func (that *Server) Reconnection(ctx context.Context) error {
	if err := that.dbConn.Close(); err != nil {
		return err
	}

	dbConn, err := that.newDBConnection(ctx)
	if err != nil {
		return err
	}
	that.dbConn = dbConn
	dbRepo, err := that.Repository(ctx)
	if err != nil {
		return err
	}
	if err := that.worker.ReCache(ctx, dbRepo); err != nil {
		return err
	}
	return nil
}

func (that *Server) Exchange(kind command.ExchangeKind, name string) error {
	switch kind {
	case command.DB:
		that.curDBName = name
	case command.Connection:
		var index int

		cfg := that.Config()
		if cfg != nil {
			for i, conn := range cfg.Connections {
				if conn.Alias == name {
					index = i + 1
					break
				}
			}
		} else {
			index, _ = strconv.Atoi(name)
		}

		if index <= 0 {
			return fmt.Errorf("specify the connection index as a numbe")
		}
		index = index - 1

		// Reconnect database
		that.curConnectionIndex = index
	}
	return nil
}
