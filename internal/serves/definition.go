package serves

import (
	"github.com/ducesoft/ulsp/ast"
	"github.com/ducesoft/ulsp/ast/astutil"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/ducesoft/ulsp/parser"
	"github.com/ducesoft/ulsp/parser/parseutil"
	"github.com/ducesoft/ulsp/token"
)

func (that *Server) Definition(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DefinitionParams) ([]lsp.Location, error) {
	f, err := ctx.Open(params.TextDocument.URI)
	if nil != err {
		return nil, cause.Error(err)
	}
	return definition(params.TextDocument.URI, f.Text, params, ctx.DB())
}

func (that *Server) TypeDefinition(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.TypeDefinitionParams) ([]lsp.Location, error) {
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
