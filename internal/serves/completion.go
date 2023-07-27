package serves

import (
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/internal/completer"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
)

func (that *Server) Completion(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.CompletionParams) (*lsp.CompletionList, error) {
	f, err := ctx.Open(params.TextDocument.URI)
	if nil != err {
		return nil, cause.Error(err)
	}
	dfg, err := ctx.DBConfig()
	if nil != err {
		return nil, cause.Error(err)
	}
	c := completer.NewCompleter(ctx.DB())
	c.Driver = dfg.Driver
	completionItems, err := c.Complete(f.Text, params, ctx.Config().LowercaseKeywords)
	if err != nil {
		return nil, err
	}
	return &lsp.CompletionList{
		IsIncomplete: true,
		ItemDefaults: nil,
		Items:        completionItems,
	}, nil
}
