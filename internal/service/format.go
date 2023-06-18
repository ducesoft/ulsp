package service

import (
	"context"
	"fmt"
	"github.com/ducesoft/ulsp/internal/formatter"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/sourcegraph/jsonrpc2"
)

func (that *Server) Formatting(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DocumentFormattingParams) ([]lsp.TextEdit, error) {
	f, ok := that.files[params.TextDocument.URI]
	if !ok {
		return nil, fmt.Errorf("document not found: %s", params.TextDocument.URI)
	}

	textEdits, err := formatter.Format(f.Text, params, that.getConfig())
	if err != nil {
		return nil, err
	}
	if len(textEdits) > 0 {
		return textEdits, nil
	}
	return nil, nil
}

func (that *Server) RangeFormatting(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DocumentRangeFormattingParams) ([]lsp.TextEdit, error) {
	_, ok := that.files[params.TextDocument.URI]
	if !ok {
		return nil, fmt.Errorf("document not found: %s", params.TextDocument.URI)
	}

	textEdits := []lsp.TextEdit{}
	if len(textEdits) > 0 {
		return textEdits, nil
	}
	return nil, nil
}
