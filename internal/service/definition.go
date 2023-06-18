package service

import (
	"context"
	"fmt"
	"github.com/ducesoft/ulsp/ast"
	"github.com/ducesoft/ulsp/ast/astutil"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/ducesoft/ulsp/parser"
	"github.com/ducesoft/ulsp/parser/parseutil"
	"github.com/ducesoft/ulsp/token"
	"github.com/sourcegraph/jsonrpc2"
)

func (that *Server) Definition(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.DefinitionParams) ([]lsp.Location, error) {
	f, ok := that.files[params.TextDocument.URI]
	if !ok {
		return nil, fmt.Errorf("document not found: %s", params.TextDocument.URI)
	}

	return definition(params.TextDocument.URI, f.Text, params, that.worker.Cache())
}

func (that *Server) TypeDefinition(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.TypeDefinitionParams) ([]lsp.Location, error) {
	return that.Definition(ctx, conn, &lsp.DefinitionParams{
		TextDocumentPositionParams: params.TextDocumentPositionParams,
		WorkDoneProgressParams:     params.WorkDoneProgressParams,
		PartialResultParams:        params.PartialResultParams,
	})
}

func definition(url lsp.DocumentURI, text string, params *lsp.DefinitionParams, dbCache *database.DBCache) ([]lsp.Location, error) {
	pos := token.Pos{
		Line: params.Position.Line,
		Col:  params.Position.Character + 1,
	}
	parsed, err := parser.Parse(text)
	if err != nil {
		return nil, err
	}

	nodeWalker := parseutil.NewNodeWalker(parsed, pos)
	m := astutil.NodeMatcher{
		NodeTypes: []ast.NodeType{ast.TypeIdentifer},
	}
	currentVariable := nodeWalker.CurNodeButtomMatched(m)
	if currentVariable == nil {
		return nil, nil
	}

	aliases := parseutil.ExtractAliased(parsed)
	if len(aliases) == 0 {
		return nil, nil
	}

	var define ast.Node
	for _, v := range aliases {
		alias, _ := v.(*ast.Aliased)
		if alias.AliasedName.String() == currentVariable.String() {
			define = alias.AliasedName
			break
		}
	}

	if define == nil {
		return nil, nil
	}

	res := []lsp.Location{
		{
			URI: url,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      define.Pos().Line,
					Character: define.Pos().Col,
				},
				End: lsp.Position{
					Line:      define.End().Line,
					Character: define.End().Col,
				},
			},
		},
	}

	return res, nil
}