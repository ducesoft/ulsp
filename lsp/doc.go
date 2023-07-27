// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lsp contains the structs that map directly to the
// request and response messages of the Language Server Protocol.
//
// It is a literal transcription, with unmodified comments, and only the changes
// required to make it go code.
// Names are uppercased to export them.
// All fields have JSON tags added to correct the names.
// Fields marked with a ? are also marked as "omitempty"
// Fields that are "|| null" are made pointers
// Fields that are string or number are left as string
// Fields that are type "number" are made float64
package lsp

import (
	"context"
	"encoding/json"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/log"
	"io"
	"runtime/debug"
)

var (
	ErrRequestCancelled = &jsonrpc2.Error{Code: -32800, Message: "JSON RPC cancelled"}
	ErrMethodNotFound   = &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "JSON RPC method not found"}
	ErrCodeParseFn      = func(err error) error { return &jsonrpc2.Error{Code: jsonrpc2.CodeParseError, Message: err.Error()} }
)

// Request is the shared interface to jsonrpc2 messages that request
// a method be invoked.
// The request types are a closed set of *Call and *Notification.
type Request interface {
	// Method is a string containing the method name to invoke.
	Method() string
	// Params is either a struct or an array with the parameters of the method.
	Params() json.RawMessage
}

type jsonRPC2Request struct {
	r *jsonrpc2.Request
}

func (that *jsonRPC2Request) Method() string {
	return that.r.Method
}

func (that *jsonRPC2Request) Params() json.RawMessage {
	return *that.r.Params
}

type connSender interface {
	io.Closer

	Notify(ctx context.Context, method string, params interface{}) error
	Call(ctx context.Context, method string, params, result interface{}) error
}

type clientDispatcher struct {
	sender connSender
}

type serverDispatcher struct {
	sender connSender
}

func Handle(ctx Context, server Server, conn *jsonrpc2.Conn, r *jsonrpc2.Request) (an any, err error) {
	return PanicEf(ctx, func() (any, error) {
		log.Warn(ctx, "%s,%v", r.Method, r.ID)
		if ctx.Err() != nil {
			return nil, ErrRequestCancelled
		}
		return serverDispatch(ctx, server, conn, &jsonRPC2Request{r: r})
	})
}

func PanicEf[T any](ctx context.Context, fn func() (T, error)) (r T, err error) {
	defer func() {
		if c := recover(); nil != c {
			log.Error(ctx, string(debug.Stack()))
			err = cause.Errorf("%v", c)
		}
	}()
	return fn()
}

func PanicEfx(ctx context.Context, fn func() error) error {
	_, err := PanicEf[any](ctx, func() (any, error) {
		return nil, fn()
	})
	return err
}
