/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package serves

import (
	"context"
	"errors"
	"fmt"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/log"
	"github.com/ducesoft/ulsp/lsp"
	ws "github.com/gorilla/websocket"
	"net/http"
	"os"
	"time"
)

var (
	grader = ws.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

type Server struct {
	*lsp.AbsServer
	cfg      *config.Config
	options  []jsonrpc2.ConnOpt
	sessions *lsp.Store
}

func NewServer(cfg *config.Config, options ...jsonrpc2.ConnOpt) *Server {
	return &Server{
		cfg:      cfg,
		options:  options,
		sessions: lsp.NewStore(),
	}
}

func (that *Server) Start() error {
	if nil == that.cfg {
		cfg, err := config.GetDefaultConfig()
		if err != nil && !errors.Is(config.ErrNotFoundConfig, err) {
			return fmt.Errorf("cannot read default config, %w", err)
		}
		that.cfg = cfg
	}
	return nil
}

func (that *Server) Stop() error {
	return that.sessions.Close()
}

func (that *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := lsp.PanicEfx(r.Context(), func() error {
		c, err := grader.Upgrade(w, r, nil)
		if nil != err {
			return err
		}
		defer func() {
			log.Catch(c.Close())
		}()
		wc := jsonrpc2.NewConn(
			jsonrpc2.NewObjectStream(c),
			jsonrpc2.HandlerWithError(that.ServeRPC),
			that.options...)
		stx := that.sessions.Open(r.Context(), that.cfg, wc)
		defer func() { that.sessions.Release(stx) }()
		wc.Serve(stx)
		<-wc.DisconnectNotify()
		return c.WriteControl(ws.CloseMessage, []byte{}, time.Now().Add(time.Second))
	}); nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (that *Server) ServeSTD(ctx context.Context) {
	h := jsonrpc2.HandlerWithError(that.ServeRPC)
	wc := jsonrpc2.NewConn(
		jsonrpc2.NewBufferedStream(&stdPip{}, jsonrpc2.VSCodeObjectCodec{}),
		h,
		that.options...,
	)
	stx := that.sessions.Open(ctx, that.cfg, wc)
	defer func() { that.sessions.Release(stx) }()
	wc.Serve(stx)
	<-wc.DisconnectNotify()
}

func (that *Server) ServeRPC(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) (any, error) {
	rs, err := lsp.Handle(lsp.ContextWith(ctx), that, conn, r)
	if nil == err {
		log.Info(ctx, "%s,%v,E0000000000", r.Method, r.ID)
		return rs, err
	}
	switch x := err.(type) {
	case *jsonrpc2.Error:
		log.Info(ctx, "%s,%v,%v", r.Method, r.ID, x.Code)
	case *cause.Cause:
		log.Info(ctx, "%s,%v,%v", r.Method, r.ID, x.GetCode())
	default:
		log.Info(ctx, "%s,%v,E0000000520", r.Method, r.ID)
	}
	return rs, err
}

type stdPip struct{}

func (stdPip) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (stdPip) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (stdPip) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	return os.Stdout.Close()
}
