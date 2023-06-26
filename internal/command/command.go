/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package command

import (
	"context"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/lsp"
	"github.com/sourcegraph/jsonrpc2"
)

func Load(command string) Command {
	return commands[command]
}

func Provide(command Command) {
	commands[command.Name()] = command
}

func Commands() []string {
	var cmds []string
	for k, _ := range commands {
		cmds = append(cmds, k)
	}
	return cmds
}

var commands = map[string]Command{}

type ExchangeKind string

const (
	DB         ExchangeKind = "database"
	Connection ExchangeKind = "connection"
)

type File interface {
	LID() string
	LText() string
}

type LS interface {
	Conn() *database.DBConnection
	Open(uri lsp.DocumentURI) File
	Repository(ctx context.Context) (database.DBRepository, error)
	Config() *config.Config
	Reconnection(ctx context.Context) error
	Exchange(kind ExchangeKind, name string) error
}

type Command interface {
	Name() string
	Attr(ctx context.Context, params *lsp.CodeActionParams) *lsp.CodeAction
	Exec(ctx context.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams, ls LS) (any, error)
}
