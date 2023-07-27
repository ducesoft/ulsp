package serves

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ducesoft/ulsp/internal/command"
	"github.com/ducesoft/ulsp/lsp"
	"strings"

	"github.com/ducesoft/ulsp/jsonrpc2"
)

func (that *Server) CodeAction(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.CodeActionParams) ([]lsp.CodeAction, error) {
	return nil, nil
}

func (that *Server) ExecuteCommand(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (interface{}, error) {
	if c := command.Load(params.Command); nil != c {
		return c.Exec(ctx, conn, params)
	}
	return nil, fmt.Errorf("unsupported command: %v", params.Command)
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
