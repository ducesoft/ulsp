package jsonrpc2

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
)

func Example() {
	ctx := context.Background()

	// Create an in-memory network connection. This connection is used below to
	// transport the JSON-RPC messages. However, any io.ReadWriteCloser may be
	// used to send/receive JSON-RPC messages.
	connA, connB := net.Pipe()

	// The following JSON-RPC connection is both a client and a server. It can
	// send requests as well as receive requests. The incoming requests are
	// handled by myHandler.
	jsonrpcConnA := NewConn(ctx, NewPlainObjectStream(connA), &myHandler{})
	defer jsonrpcConnA.Close()

	// The following JSON-RPC connection has no handler, meaning that it is
	// configured to only be a client. It can send requests and receive the
	// responses to those requests, but it will ignore any incoming requests.
	jsonrpcConnB := NewConn(ctx, NewPlainObjectStream(connB), nil)
	defer jsonrpcConnB.Close()

	// Send a request from jsonrpcConnB to jsonrpcConnA. The result of a
	// successful call is stored in the result variable.
	var result string
	if err := jsonrpcConnB.Call(ctx, "sayHello", nil, &result); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(result)

	// Output: hello world
}

// myHandler is the Handler used by jsonrpcConnA.
type myHandler struct{}

// Handle implements the Handler interface.
func (h *myHandler) Handle(ctx context.Context, c *Conn, r *Request) {
	switch r.Method {
	case "sayHello":
		if err := c.Reply(ctx, r.ID, "hello world"); err != nil {
			log.Println(err)
			return
		}
	default:
		err := &Error{Code: CodeMethodNotFound, Message: "Method not found"}
		if err := c.ReplyWithError(ctx, r.ID, err); err != nil {
			log.Println(err)
			return
		}
	}
}
