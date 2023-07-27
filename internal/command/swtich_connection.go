/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package command

import (
	"encoding/json"
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

func (that *switchConnection) Attr(ctx lsp.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
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

func (that *switchConnection) Exec(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (any, error) {
	return nil, nil
}
