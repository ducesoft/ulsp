package serves

import (
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
)

var (
	ColorKeyword = lsp.Color{Red: 198, Green: 144, Blue: 114, Alpha: 1}
)

func (that *Server) DocumentColor(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DocumentColorParams) ([]lsp.ColorInformation, error) {
	return nil, nil
}

func (that *Server) DocumentHighlight(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DocumentHighlightParams) ([]lsp.DocumentHighlight, error) {
	var hls []lsp.DocumentHighlight
	return hls, nil
}
func (that *Server) DocumentLink(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DocumentLinkParams) ([]lsp.DocumentLink, error) {
	var links []lsp.DocumentLink
	return links, nil
}
func (that *Server) DocumentSymbol(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DocumentSymbolParams) ([]any, error) {
	var rs []any
	return rs, nil
}

func (that *Server) FoldingRange(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.FoldingRangeParams) ([]lsp.FoldingRange, error) {
	var frs []lsp.FoldingRange
	return frs, nil
}

func (that *Server) InlayHint(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.InlayHintParams) ([]lsp.InlayHint, error) {
	return nil, nil
}

func (that *Server) CodeLens(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.CodeLensParams) ([]lsp.CodeLens, error) {
	return nil, nil
}

func (that *Server) ColorPresentation(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ColorPresentationParams) ([]lsp.ColorPresentation, error) {
	return nil, nil
}
