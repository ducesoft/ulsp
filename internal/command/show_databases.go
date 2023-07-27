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
	Provide(new(showDatabases))
}

type showDatabases struct {
}

func (that *showDatabases) Name() string {
	return "code/showDatabases"
}

func (that *showDatabases) Attr(ctx lsp.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
	return &lsp.CodeAction{
		Title: i18n.Sprintf(ctx, "Show DataSources"),
		Kind:  lsp.Empty,
		Command: &lsp.Command{
			Title:     "Show DataSources",
			Command:   that.Name(),
			Arguments: []json.RawMessage{},
		},
	}
}

func (that *showDatabases) Exec(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (any, error) {
	repo, err := ctx.Repository()
	if err != nil {
		return "", err
	}
	databases, err := repo.Databases(ctx)
	if err != nil {
		return nil, err
	}
	return strings.Join(databases, "\n"), nil
}
