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
	Provide(new(switchDatabase))
}

type switchDatabase struct {
}

func (that *switchDatabase) Name() string {
	return "code/switchDatabase"
}

func (that *switchDatabase) Attr(ctx lsp.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
	return &lsp.CodeAction{
		Title: i18n.Sprintf(ctx, "Switch Database"),
		Kind:  lsp.Empty,
		Command: &lsp.Command{
			Title:     "Switch Database",
			Command:   that.Name(),
			Arguments: []json.RawMessage{},
		},
	}
}

func (that *switchDatabase) Exec(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (any, error) {
	return nil, nil
}
