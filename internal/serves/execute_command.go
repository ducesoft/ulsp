package serves

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/lsp"
	"io"
	"strconv"
	"strings"

	"github.com/ducesoft/ulsp/ast"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/parser"
	"github.com/olekukonko/tablewriter"
	"github.com/sourcegraph/jsonrpc2"
)

const (
	CommandExecuteQuery     = "executeQuery"
	CommandShowDatabases    = "showDatabases"
	CommandShowSchemas      = "showSchemas"
	CommandShowConnections  = "showConnections"
	CommandSwitchDatabase   = "switchDatabase"
	CommandSwitchConnection = "switchConnections"
	CommandShowTables       = "showTables"
)

func (that *Server) CodeAction(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.CodeActionParams) ([]lsp.CodeAction, error) {
	commands := []lsp.CodeAction{
		{
			Command: &lsp.Command{
				Title:     "Execute Query",
				Command:   CommandExecuteQuery,
				Arguments: []json.RawMessage{[]byte(fmt.Sprintf("\"%s\"", params.TextDocument.URI))},
			},
		},
		{
			Command: &lsp.Command{
				Title:     "Show Databases",
				Command:   CommandShowDatabases,
				Arguments: []json.RawMessage{},
			},
		},
		{
			Command: &lsp.Command{
				Title:     "Show Schemas",
				Command:   CommandShowSchemas,
				Arguments: []json.RawMessage{},
			},
		},
		{
			Command: &lsp.Command{
				Title:     "Show Connections",
				Command:   CommandShowConnections,
				Arguments: []json.RawMessage{},
			},
		},
		{
			Command: &lsp.Command{
				Title:     "Switch Database",
				Command:   CommandSwitchDatabase,
				Arguments: []json.RawMessage{},
			},
		},
		{
			Command: &lsp.Command{
				Title:     "Switch Connections",
				Command:   CommandSwitchConnection,
				Arguments: []json.RawMessage{},
			},
		},
		{
			Command: &lsp.Command{
				Title:     "Show Tables",
				Command:   CommandShowTables,
				Arguments: []json.RawMessage{},
			},
		},
	}
	return commands, nil
}

func (that *Server) ExecuteCommand(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (interface{}, error) {
	switch params.Command {
	case CommandExecuteQuery:
		return that.executeQuery(ctx, params)
	case CommandShowDatabases:
		return that.showDatabases(ctx, params)
	case CommandShowSchemas:
		return that.showSchemas(ctx, params)
	case CommandShowConnections:
		return that.showConnections(ctx, params)
	case CommandSwitchDatabase:
		return that.switchDatabase(ctx, params)
	case CommandSwitchConnection:
		return that.switchConnections(ctx, params)
	case CommandShowTables:
		return that.showTables(ctx, params)
	}
	return nil, fmt.Errorf("unsupported command: %v", params.Command)
}

func (that *Server) executeQuery(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	// parse execute command arguments
	if that.dbConn == nil {
		return nil, errors.New("database connection is not open")
	}
	if len(params.Arguments) == 0 {
		return nil, fmt.Errorf("required arguments were not provided: <File URI>")
	}
	uri := lsp.DocumentURI(params.Arguments[0])
	f, ok := that.files[uri]
	if !ok {
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
	text := f.Text
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
			res, err := that.query(ctx, query, showVertical)
			if err != nil {
				return nil, err
			}
			fmt.Fprintln(buf, res)
		} else {
			res, err := that.exec(ctx, query, showVertical)
			if err != nil {
				return nil, err
			}
			fmt.Fprintln(buf, res)
		}
	}
	return buf.String(), nil
}

func extractRangeText(text string, startLine, startChar, endLine, endChar int) string {
	writer := bytes.NewBufferString("")
	scanner := bufio.NewScanner(strings.NewReader(text))

	i := 0
	for scanner.Scan() {
		t := scanner.Text()
		if i >= startLine && i <= endLine {
			st, en := 0, len(t)

			if i == startLine {
				st = startChar
			}
			if i == endLine {
				en = endChar
			}

			writer.Write([]byte(t[st:en]))
			if i != endLine {
				writer.Write([]byte("\n"))
			}
		}
		i++
	}
	return writer.String()
}

func (that *Server) query(ctx context.Context, query string, vertical bool) (string, error) {
	repo, err := that.newDBRepository(ctx)
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

func (that *Server) exec(ctx context.Context, query string, vertical bool) (string, error) {
	repo, err := that.newDBRepository(ctx)
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

func (that *Server) showDatabases(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	repo, err := that.newDBRepository(ctx)
	if err != nil {
		return "", err
	}
	databases, err := repo.Databases(ctx)
	if err != nil {
		return nil, err
	}
	return strings.Join(databases, "\n"), nil
}

func (that *Server) showSchemas(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	repo, err := that.newDBRepository(ctx)
	if err != nil {
		return "", err
	}
	schemas, err := repo.Schemas(ctx)
	if err != nil {
		return nil, err
	}
	return strings.Join(schemas, "\n"), nil
}

func (that *Server) switchDatabase(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	if len(params.Arguments) != 1 {
		return nil, fmt.Errorf("required arguments were not provided: <DB Name>")
	}
	dbName := string(params.Arguments[0])
	// Change current database
	that.curDBName = dbName

	// close and reconnection to database
	if err := that.reconnectionDB(ctx); err != nil {
		return nil, err
	}

	return nil, nil
}

func (that *Server) showConnections(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	var results []string
	conns := that.getConfig().Connections
	for i, conn := range conns {
		var desc string
		if conn.DataSourceName != "" {
			desc = conn.DataSourceName
		} else {
			switch conn.Proto {
			case config.ProtoTCP:
				desc = fmt.Sprintf("tcp(%s:%d)/%s", conn.Host, conn.Port, conn.DBName)
			case config.ProtoUDP:
				desc = fmt.Sprintf("udp(%s:%d)/%s", conn.Host, conn.Port, conn.DBName)
			case config.ProtoUnix:
				desc = fmt.Sprintf("unix(%s)/%s", conn.Path, conn.DBName)
			}
		}
		res := fmt.Sprintf("%d %s %s %s", i+1, conn.Driver, conn.Alias, desc)
		results = append(results, res)
	}
	return strings.Join(results, "\n"), nil
}

func (that *Server) switchConnections(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	if len(params.Arguments) != 1 {
		return nil, fmt.Errorf("required arguments were not provided: <Connection Index>")
	}
	indexStr := string(params.Arguments[0])
	var index int

	cfg := that.getConfig()
	if cfg != nil {
		for i, conn := range cfg.Connections {
			if conn.Alias == indexStr {
				index = i + 1
				break
			}
		}
	} else {
		index, _ = strconv.Atoi(indexStr)
	}

	if index <= 0 {
		return nil, fmt.Errorf("specify the connection index as a number, %w", err)
	}
	index = index - 1

	// Reconnect database
	that.curConnectionIndex = index

	// close and reconnection to database
	if err := that.reconnectionDB(ctx); err != nil {
		return nil, err
	}

	return nil, nil
}

func (that *Server) showTables(ctx context.Context, params *lsp.ExecuteCommandParams) (result interface{}, err error) {
	repo, err := that.newDBRepository(ctx)
	if err != nil {
		return "", err
	}
	m, err := repo.SchemaTables(ctx)
	if err != nil {
		return nil, err
	}
	schema, err := repo.CurrentSchema(ctx)
	if err != nil {
		return nil, err
	}
	results := []string{}
	for k, vv := range m {
		for _, v := range vv {
			if k != "" {
				if schema != k {
					continue
				}
				results = append(results, k+"."+v)
			} else {
				results = append(results, v)
			}
		}
	}
	return strings.Join(results, "\n"), nil
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
