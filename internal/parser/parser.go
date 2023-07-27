/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package parser

import (
	"context"
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/ducesoft/ulsp/internal/parser/mysql"
	"github.com/ducesoft/ulsp/lsp"
	"strings"
)

func init() {
	parser.MySqlLexerInit()
	for _, name := range parser.MySqlLexerLexerStaticData.SymbolicNames {
		symbolicNames[name] = true
	}
}

var symbolicNames = map[string]bool{}

type keywordListener struct {
	*parser.BaseMySqlParserListener
	keywords []lsp.Range
}

func (that *keywordListener) VisitTerminal(node antlr.TerminalNode) {
	s := node.GetSymbol()
	keyword := strings.ToUpper(s.GetText())
	if symbolicNames[keyword] {
		that.keywords = append(that.keywords, lsp.Range{
			Start: lsp.Position{
				Line:      s.GetLine(),
				Character: s.GetStart(),
			},
			End: lsp.Position{
				Line:      s.GetLine(),
				Character: s.GetStop(),
			},
		})
	}
}

func ParseKeywords(ctx context.Context, text string) []lsp.Range {
	input := antlr.NewInputStream(text)
	lexer := parser.NewMySqlLexer(input)
	token := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	tree := parser.NewMySqlParser(token)
	// tree.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// tree.BuildParseTrees = true
	listener := &keywordListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree.Root())
	return listener.keywords
}
