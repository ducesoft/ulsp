package lsp

import (
	"context"
	"log"

	"github.com/ducesoft/ulsp/jsonrpc2"
)

type Messenger interface {
	ShowLog(context.Context, string) error
	ShowInfo(context.Context, string) error
	ShowWarning(context.Context, string) error
	ShowError(context.Context, string) error
}

type lspMessenger struct {
	conn *jsonrpc2.Conn
}

func NewLspMessenger(conn *jsonrpc2.Conn) Messenger {
	return &lspMessenger{
		conn: conn,
	}
}

func (that *lspMessenger) ShowLog(ctx context.Context, message string) error {
	log.Println("Send Message:", message)
	params := &ShowMessageParams{
		Type:    Log,
		Message: message,
	}
	return that.conn.Notify(ctx, "window/showMessage", params)
}

func (that *lspMessenger) ShowInfo(ctx context.Context, message string) error {
	log.Println("Send Message:", message)
	params := &ShowMessageParams{
		Type:    Info,
		Message: message,
	}
	return that.conn.Notify(ctx, "window/showMessage", params)
}

func (that *lspMessenger) ShowWarning(ctx context.Context, message string) error {
	log.Println("Send Message:", message)
	params := &ShowMessageParams{
		Type:    Warning,
		Message: message,
	}
	return that.conn.Notify(ctx, "window/showMessage", params)
}

func (that *lspMessenger) ShowError(ctx context.Context, message string) error {
	log.Println("Send Message:", message)
	params := &ShowMessageParams{
		Type:    Error,
		Message: message,
	}
	return that.conn.Notify(ctx, "window/showMessage", params)
}
