/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package command

import (
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
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

type Command interface {
	Name() string
	Attr(ctx lsp.Context, params *lsp.CodeActionParams) *lsp.CodeAction
	Exec(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ExecuteCommandParams) (any, error)
}
