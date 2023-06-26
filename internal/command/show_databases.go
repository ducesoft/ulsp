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
	Provide(new(showDatabases))
}

type showDatabases struct {
}

func (that *showDatabases) Name() string {
	return "code/showDatabases"
}

func (that *showDatabases) Attr(ctx context.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
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

func (that *showDatabases) Exec(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams, ls LS) (any, error) {
	repo, err := ls.Repository(ctx)
	if err != nil {
		return "", err
	}
	databases, err := repo.Databases(ctx)
	if err != nil {
		return nil, err
	}
	return strings.Join(databases, "\n"), nil
}
