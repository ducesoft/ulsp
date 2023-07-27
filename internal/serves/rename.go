package serves

import (
	"github.com/ducesoft/ulsp/ast"
	"github.com/ducesoft/ulsp/ast/astutil"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/ducesoft/ulsp/parser"
	"github.com/ducesoft/ulsp/parser/parseutil"
	"github.com/ducesoft/ulsp/token"
)

func (that *Server) Rename(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.RenameParams) (*lsp.WorkspaceEdit, error) {
	f, err := ctx.Open(params.TextDocument.URI)
	if nil != err {
		return nil, cause.Errors(err)
	}
	parsed, err := parser.Parse(f.Text)
	if err != nil {
		return nil, err
	}

	pos := token.Pos{
		Line: params.Position.Line,
		Col:  params.Position.Character,
	}

	// Get the identifer on focus
	nodeWalker := parseutil.NewNodeWalker(parsed, pos)
	m := astutil.NodeMatcher{
		NodeTypes: []ast.NodeType{ast.TypeIdentifer},
	}
	currentVariable := nodeWalker.CurNodeButtomMatched(m)
	if currentVariable == nil {
		return nil, nil
	}

	// Get all identifiers in the statement
	idents, err := parseutil.ExtractIdenfiers(parsed, pos)
	if err != nil {
		return nil, err
	}

	// Extract only those with matching names
	renameTarget := []ast.Node{}
	for _, ident := range idents {
		if ident.String() == currentVariable.String() {
			renameTarget = append(renameTarget, ident)
		}
	}
	if len(renameTarget) == 0 {
		return nil, nil
	}

	edits := make([]lsp.TextEdit, len(renameTarget))
	for i, target := range renameTarget {
		edit := lsp.TextEdit{
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      target.Pos().Line,
					Character: target.Pos().Col,
				},
				End: lsp.Position{
					Line:      target.End().Line,
					Character: target.End().Col,
				},
			},
			NewText: params.NewName,
		}
		edits[i] = edit
	}

	res := &lsp.WorkspaceEdit{
		DocumentChanges: []lsp.DocumentChanges{
			{
				TextDocumentEdit: &lsp.TextDocumentEdit{
					TextDocument: lsp.OptionalVersionedTextDocumentIdentifier{
						Version: 0,
						TextDocumentIdentifier: lsp.TextDocumentIdentifier{
							URI: params.TextDocument.URI,
						},
					},
					Edits: edits,
				},
			},
		},
	}
	return res, nil
}
