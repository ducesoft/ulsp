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
	"strings"
)

func init() {
	Provide(new(showSchemas))
}

type showSchemas struct {
}

func (that *showSchemas) Name() string {
	return "code/switchConnection"
}

func (that *showSchemas) Attr(ctx lsp.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
	return &lsp.CodeAction{
		Title: i18n.Sprintf(ctx, "Show Schemas"),
		Kind:  lsp.Empty,
		Command: &lsp.Command{
			Title:     "Show Schemas",
			Command:   that.Name(),
			Arguments: []json.RawMessage{},
		},
	}
}

func (that *showSchemas) Exec(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (any, error) {
	repo, err := ctx.Repository()
	if err != nil {
		return "", err
	}
	schemas, err := repo.Schemas(ctx)
	if err != nil {
		return nil, err
	}
	return strings.Join(schemas, "\n"), nil
}
