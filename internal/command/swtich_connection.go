/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package command

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ducesoft/ulsp/internal/i18n"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
)

func init() {
	Provide(new(switchConnection))
}

type switchConnection struct {
}

func (that *switchConnection) Name() string {
	return "code/switchConnection"
}

func (that *switchConnection) Attr(ctx context.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
	return &lsp.CodeAction{
		Title: i18n.Sprintf(ctx, "Switch Connection"),
		Kind:  lsp.Empty,
		Command: &lsp.Command{
			Title:     "Switch Connection",
			Command:   that.Name(),
			Arguments: []json.RawMessage{},
		},
	}
}

func (that *switchConnection) Exec(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams, ls LS) (any, error) {
	if len(params.Arguments) != 1 {
		return nil, fmt.Errorf("required arguments were not provided: <Connection Index>")
	}
	indexStr := string(params.Arguments[0])
	// Reconnect database
	if err := ls.Exchange(Connection, indexStr); nil != err {
		return nil, err
	}

	// close and reconnection to database
	if err := ls.Reconnection(ctx); err != nil {
		return nil, err
	}

	return nil, nil
}
