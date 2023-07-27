/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package command

import (
	"encoding/json"
	"fmt"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/i18n"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
	"strings"
)

func init() {
	Provide(new(showConnections))
}

type showConnections struct {
}

func (that *showConnections) Name() string {
	return "code/switchConnection"
}

func (that *showConnections) Attr(ctx lsp.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
	return &lsp.CodeAction{
		Title: i18n.Sprintf(ctx, "Show Connections"),
		Kind:  lsp.Empty,
		Command: &lsp.Command{
			Title:     "Show Connections",
			Command:   that.Name(),
			Arguments: []json.RawMessage{},
		},
	}
}

func (that *showConnections) Exec(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (any, error) {
	var results []string
	conns := ctx.Config().Connections
	for i, cf := range conns {
		var desc string
		if cf.DataSourceName != "" {
			desc = cf.DataSourceName
		} else {
			switch cf.Proto {
			case config.ProtoTCP:
				desc = fmt.Sprintf("tcp(%s:%d)/%s", cf.Host, cf.Port, cf.DBName)
			case config.ProtoUDP:
				desc = fmt.Sprintf("udp(%s:%d)/%s", cf.Host, cf.Port, cf.DBName)
			case config.ProtoUnix:
				desc = fmt.Sprintf("unix(%s)/%s", cf.Path, cf.DBName)
			}
		}
		res := fmt.Sprintf("%d %s %s %s", i+1, cf.Driver, cf.Alias, desc)
		results = append(results, res)
	}
	return strings.Join(results, "\n"), nil
}
