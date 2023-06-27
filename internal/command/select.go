/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package command

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ducesoft/ulsp/ast"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/internal/i18n"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/ducesoft/ulsp/parser"
	"github.com/olekukonko/tablewriter"
	"io"
	"strconv"
	"strings"
)

func init() {
	Provide(new(executor))
}

type executor struct {
}

func (that *executor) Name() string {
	return "code/execute"
}

func (that *executor) Attr(ctx context.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
	return &lsp.CodeAction{
		Title: i18n.Sprintf(ctx, "Execute Query"),
		Kind:  lsp.Empty,
		Command: &lsp.Command{
			Title:     "Execute Query",
			Command:   that.Name(),
			Arguments: []json.RawMessage{[]byte(fmt.Sprintf("\"%s\"", params.TextDocument.URI))},
		},
	}
}

func (that *executor) Exec(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams, ls LS) (any, error) {
	// parse execute command arguments
	if nil == ls.Conn() {
		return nil, errors.New("database connection is not open")
	}
	if len(params.Arguments) == 0 {
		return nil, fmt.Errorf("required arguments were not provided: <File URI>")
	}
	uri := lsp.DocumentURI(params.Arguments[0])
	f := ls.Open(uri)
	if nil == f {
		return nil, fmt.Errorf("document not found, %q", uri)
	}
	showVertical := false
	if len(params.Arguments) > 1 {
		showVerticalFlag := string(params.Arguments[1])
		if showVerticalFlag == "-show-vertical" {
			showVertical = true
		}
	}

	// extract target query
	text := f.LText()
	// TODO XXX
	//if params.Range != nil {
	//	text = extractRangeText(
	//		text,
	//		params.Range.Start.Line,
	//		params.Range.Start.Character,
	//		params.Range.End.Line,
	//		params.Range.End.Character,
	//	)
	//}
	stmts, err := getStatements(text)
	if err != nil {
		return nil, err
	}

	// execute statements
	buf := new(bytes.Buffer)
	for _, stmt := range stmts {
		query := strings.TrimSpace(stmt.String())
		if query == "" {
			continue
		}

		if _, isQuery := database.QueryExecType(query, ""); isQuery {
			res, err := that.query(ctx, query, showVertical, ls)
			if err != nil {
				return nil, err
			}
			fmt.Fprintln(buf, res)
		} else {
			res, err := that.exec(ctx, query, showVertical, ls)
			if err != nil {
				return nil, err
			}
			fmt.Fprintln(buf, res)
		}
	}
	return buf.String(), nil
}

func getStatements(text string) ([]*ast.Statement, error) {
	parsed, err := parser.Parse(text)
	if err != nil {
		return nil, err
	}

	var stmts []*ast.Statement
	for _, node := range parsed.GetTokens() {
		stmt, ok := node.(*ast.Statement)
		if !ok {
			return nil, fmt.Errorf("invalid type want Statement parsed %T", stmt)
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
}

func (that *executor) query(ctx context.Context, query string, vertical bool, ls LS) (string, error) {
	repo, err := ls.Repository(ctx)
	if err != nil {
		return "", err
	}
	rows, err := repo.Query(ctx, query)
	if err != nil {
		return "", err
	}
	columns, err := database.Columns(rows)
	if err != nil {
		return "", err
	}
	stringRows, err := database.ScanRows(rows, len(columns))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if vertical {
		table := newVerticalTableWriter(buf)
		table.setHeaders(columns)
		for _, stringRow := range stringRows {
			table.appendRow(stringRow)
		}
		table.render()
	} else {
		table := tablewriter.NewWriter(buf)
		table.SetHeader(columns)
		for _, stringRow := range stringRows {
			table.Append(stringRow)
		}
		table.Render()
	}
	fmt.Fprintf(buf, "%d rows in set", len(stringRows))
	fmt.Fprintln(buf, "")
	fmt.Fprintln(buf, "")
	return buf.String(), nil
}

func (that *executor) exec(ctx context.Context, query string, vertical bool, ls LS) (string, error) {
	repo, err := ls.Repository(ctx)
	if err != nil {
		return "", err
	}
	result, err := repo.Exec(ctx, query)
	if err != nil {
		return "", err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "Query OK, %d row affected", rowsAffected)
	fmt.Fprintln(buf, "")
	fmt.Fprintln(buf, "")
	return buf.String(), nil
}

type verticalTableWriter struct {
	writer       io.Writer
	headers      []string
	rows         [][]string
	headerMaxLen int
}

func newVerticalTableWriter(writer io.Writer) *verticalTableWriter {
	return &verticalTableWriter{
		writer: writer,
	}
}

func (vtw *verticalTableWriter) setHeaders(headers []string) {
	vtw.headers = headers
	for _, h := range headers {
		length := len(h)
		if vtw.headerMaxLen < length {
			vtw.headerMaxLen = length
		}
	}
}

func (vtw *verticalTableWriter) appendRow(row []string) {
	vtw.rows = append(vtw.rows, row)
}

func (vtw *verticalTableWriter) render() {
	for rowNum, row := range vtw.rows {
		fmt.Fprintf(vtw.writer, "***************************[ %d. row ]***************************", rowNum+1)
		fmt.Fprintln(vtw.writer, "")
		for colNum, col := range row {
			header := vtw.headers[colNum]

			padHeader := fmt.Sprintf("%"+strconv.Itoa(vtw.headerMaxLen)+"s", header)
			fmt.Fprintf(vtw.writer, "%s | %s", padHeader, col)
			fmt.Fprintln(vtw.writer, "")
		}
	}
}
