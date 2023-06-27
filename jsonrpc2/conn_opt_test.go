package jsonrpc2

import (
	"bufio"
	"context"
	"io"
	"net"
	"testing"
)

func TestSetLogger(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rd, wr := io.Pipe()
	defer rd.Close()
	defer wr.Close()

	buf := bufio.NewReader(rd)

	a, b := net.Pipe()
	connA := NewConn(
		ctx,
		NewBufferedStream(a, VSCodeObjectCodec{}),
		noopHandler{},
	)
	connB := NewConn(
		ctx,
		NewBufferedStream(b, VSCodeObjectCodec{}),
		noopHandler{},
	)
	defer connA.Close()
	defer connB.Close()

	// Write a response with no corresponding request.
	if err := connB.Reply(ctx, ID{Num: 0}, nil); err != nil {
		t.Fatal(err)
	}

	want := "jsonrpc2: ignoring response #0 with no corresponding request\n"
	got, err := buf.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
