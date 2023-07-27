package serves

import (
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/internal/formatter"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
)

func (that *Server) Formatting(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DocumentFormattingParams) ([]lsp.TextEdit, error) {
	f, err := ctx.Open(params.TextDocument.URI)
	if nil != err {
		return nil, cause.Errors(err)
	}
	textEdits, err := formatter.Format(f.Text, params, ctx.Config())
	if err != nil {
		return nil, err
	}
	if len(textEdits) > 0 {
		return textEdits, nil
	}
	return nil, nil
}

func (that *Server) RangeFormatting(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DocumentRangeFormattingParams) ([]lsp.TextEdit, error) {
	_, err := ctx.Open(params.TextDocument.URI)
	if nil != err {
		return nil, cause.Errors(err)
	}
	textEdits := []lsp.TextEdit{}
	if len(textEdits) > 0 {
		return textEdits, nil
	}
	return nil, nil
}
