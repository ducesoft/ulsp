package serves

import (
	"fmt"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/ducesoft/ulsp/parser"
	"github.com/ducesoft/ulsp/parser/parseutil"
	"github.com/ducesoft/ulsp/token"
)

func (that *Server) SignatureHelp(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.SignatureHelpParams) (*lsp.SignatureHelp, error) {
	f, err := ctx.Open(params.TextDocument.URI)
	if nil != err {
		return nil, cause.Error(err)
	}
	parsed, err := parser.Parse(f.Text)
	if err != nil {
		return nil, err
	}
	pos := token.Pos{
		Line: params.Position.Line,
		Col:  params.Position.Character,
	}
	nodeWalker := parseutil.NewNodeWalker(parsed, pos)
	types := getSignatureHelpTypes(nodeWalker)

	switch {
	case signatureHelpIs(types, SignatureHelpTypeInsertValue):
		insert, err := parseutil.ExtractInsert(parsed, pos)
		if err != nil {
			return nil, err
		}
		if !insert.Enable() {
			return nil, err
		}

		table := insert.GetTable()
		cols := insert.GetColumns()
		paramIdx := insert.GetValues().GetIndex(pos)
		tableName := table.Name

		var infos []lsp.ParameterInformation
		for _, col := range cols.GetIdentifers() {
			colName := col.String()
			colDoc := ""
			colDesc, ok := ctx.DB().Column(tableName, colName)
			if ok {
				colDoc = colDesc.OnelineDesc()
			}
			p := lsp.ParameterInformation{
				Label:         colName,
				Documentation: colDoc,
			}
			infos = append(infos, p)
		}

		documentation := &lsp.Or_SignatureInformation_documentation{}
		if err = documentation.UnmarshalJSON([]byte(fmt.Sprintf("%s table columns", tableName))); nil != err {
			return nil, err
		}
		signatureLabel := fmt.Sprintf("%s (%s)", tableName, cols.String())
		sh := &lsp.SignatureHelp{
			Signatures: []lsp.SignatureInformation{
				{
					Label:         signatureLabel,
					Documentation: documentation,
					Parameters:    infos,
				},
			},
			ActiveSignature: 0.0,
			ActiveParameter: paramIdx,
		}
		return sh, nil
	default:
		// pass
		return nil, nil
	}
}

type signatureHelpType int

const (
	_ signatureHelpType = iota
	SignatureHelpTypeInsertValue
	SignatureHelpTypeUnknown = 99
)

func (sht signatureHelpType) String() string {
	switch sht {
	case SignatureHelpTypeInsertValue:
		return "InsertValue"
	default:
		return ""
	}
}

func getSignatureHelpTypes(nw *parseutil.NodeWalker) []signatureHelpType {
	syntaxPos := parseutil.CheckSyntaxPosition(nw)
	types := []signatureHelpType{}
	switch {
	case syntaxPos == parseutil.InsertValue:
		types = []signatureHelpType{
			SignatureHelpTypeInsertValue,
		}
	default:
		// pass
	}
	return types
}

func signatureHelpIs(types []signatureHelpType, expect signatureHelpType) bool {
	for _, t := range types {
		if t == expect {
			return true
		}
	}
	return false
}
