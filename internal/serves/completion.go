package serves

import (
	"context"
	"fmt"
	"github.com/ducesoft/ulsp/internal/completer"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
)

func (that *Server) Completion(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.CompletionParams) (*lsp.CompletionList, error) {
	f, ok := that.files[params.TextDocument.URI]
	if !ok {
		return nil, fmt.Errorf("document not found: %s", params.TextDocument.URI)
	}

	c := completer.NewCompleter(that.worker.Cache())
	if that.dbConn != nil {
		c.Driver = that.dbConn.Driver
	} else {
		c.Driver = ""
	}
	completionItems, err := c.Complete(f.Text, params, that.Config().LowercaseKeywords)
	if err != nil {
		return nil, err
	}
	return &lsp.CompletionList{
		IsIncomplete: true,
		ItemDefaults: nil,
		Items:        completionItems,
	}, nil
}
