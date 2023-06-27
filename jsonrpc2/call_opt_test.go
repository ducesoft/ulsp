package jsonrpc2

import (
	"context"
	"fmt"
	"testing"
)

func TestPickID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, b := inMemoryPeerConns()
	defer a.Close()
	defer b.Close()

	handler := handlerFunc(func(ctx context.Context, conn *Conn, req *Request) {
		if err := conn.Reply(ctx, req.ID, fmt.Sprintf("hello, #%s: %s", req.ID, *req.Params)); err != nil {
			t.Error(err)
		}
	})
	connA := NewConn(ctx, NewBufferedStream(a, VSCodeObjectCodec{}), handler)
	connB := NewConn(ctx, NewBufferedStream(b, VSCodeObjectCodec{}), noopHandler{})
	defer connA.Close()
	defer connB.Close()

	const n = 100
	for i := 0; i < n; i++ {
		var opts []CallOption
		id := ID{Num: uint64(i)}

		// This is the actual test, every 3rd request we specify the
		// ID and ensure we get a response with the correct ID echoed
		// back
		if i%3 == 0 {
			id = ID{
				Str:      fmt.Sprintf("helloworld-%d", i/3),
				IsString: true,
			}
			opts = append(opts, PickID(id))
		}

		var got string
		if err := connB.Call(ctx, "f", []int32{1, 2, 3}, &got, opts...); err != nil {
			t.Fatal(err)
		}
		if want := fmt.Sprintf("hello, #%s: [1,2,3]", id); got != want {
			t.Errorf("got result %q, want %q", got, want)
		}
	}
}

func TestStringID(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, b := inMemoryPeerConns()
	defer a.Close()
	defer b.Close()

	handler := handlerFunc(func(ctx context.Context, conn *Conn, req *Request) {
		replyWithError := func(msg string) {
			respErr := &Error{Code: CodeInvalidRequest, Message: msg}
			if err := conn.ReplyWithError(ctx, req.ID, respErr); err != nil {
				t.Error(err)
			}
		}
		if !req.ID.IsString {
			replyWithError("ID.IsString should be true")
			return
		}
		if len(req.ID.Str) == 0 {
			replyWithError("ID.Str should be populated but is empty")
			return
		}
		if err := conn.Reply(ctx, req.ID, "ok"); err != nil {
			t.Error(err)
		}
	})
	connA := NewConn(ctx, NewBufferedStream(a, VSCodeObjectCodec{}), handler)
	connB := NewConn(ctx, NewBufferedStream(b, VSCodeObjectCodec{}), noopHandler{})
	defer connA.Close()
	defer connB.Close()

	var res string
	if err := connB.Call(ctx, "f", nil, &res, StringID()); err != nil {
		t.Fatal(err)
	}
}

func TestExtraField(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, b := inMemoryPeerConns()
	defer a.Close()
	defer b.Close()

	handler := handlerFunc(func(ctx context.Context, conn *Conn, req *Request) {
		replyWithError := func(msg string) {
			respErr := &Error{Code: CodeInvalidRequest, Message: msg}
			if err := conn.ReplyWithError(ctx, req.ID, respErr); err != nil {
				t.Error(err)
			}
		}
		var sessionID string
		for _, field := range req.ExtraFields {
			if field.Name != "sessionId" {
				continue
			}
			var ok bool
			sessionID, ok = field.Value.(string)
			if !ok {
				t.Errorf("\"sessionId\" is not a string: %v", field.Value)
			}
		}
		if sessionID == "" {
			replyWithError("sessionId must be set")
			return
		}
		if sessionID != "session" {
			replyWithError("sessionId has the wrong value")
			return
		}
		if err := conn.Reply(ctx, req.ID, "ok"); err != nil {
			t.Error(err)
		}
	})
	connA := NewConn(ctx, NewBufferedStream(a, VSCodeObjectCodec{}), handler)
	connB := NewConn(ctx, NewBufferedStream(b, VSCodeObjectCodec{}), noopHandler{})
	defer connA.Close()
	defer connB.Close()

	var res string
	if err := connB.Call(ctx, "f", nil, &res, ExtraField("sessionId", "session")); err != nil {
		t.Fatal(err)
	}
}
