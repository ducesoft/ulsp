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
	Provide(new(switchDatabase))
}

type switchDatabase struct {
}

func (that *switchDatabase) Name() string {
	return "code/switchDatabase"
}

func (that *switchDatabase) Attr(ctx context.Context, params *lsp.CodeActionParams) *lsp.CodeAction {
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

func (that *switchDatabase) Exec(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams, ls LS) (any, error) {
	if len(params.Arguments) != 1 {
		return nil, fmt.Errorf("required arguments were not provided: <DB Name>")
	}
	dbName := string(params.Arguments[0])
	// Change current database
	if err := ls.Exchange(DB, dbName); nil != err {
		return nil, err
	}

	// close and reconnection to database
	if err := ls.Reconnection(ctx); err != nil {
		return nil, err
	}

	return nil, nil
}
