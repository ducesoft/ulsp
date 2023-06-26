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
	"github.com/ducesoft/ulsp/internal/i18n"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/sourcegraph/jsonrpc2"
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

func (that *showSchemas) Attr(ctx context.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
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

func (that *showSchemas) Exec(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams, ls LS) (any, error) {
	repo, err := ls.Repository(ctx)
	if err != nil {
		return "", err
	}
	schemas, err := repo.Schemas(ctx)
	if err != nil {
		return nil, err
	}
	return strings.Join(schemas, "\n"), nil
}
