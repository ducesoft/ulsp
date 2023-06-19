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
	"github.com/ducesoft/ulsp/internal/config"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/lsp"
	ws "github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/sourcegraph/jsonrpc2/websocket"
	"net/http"
	"os"
	"time"
)

var (
	ErrNoConnection = errors.New("no database connection")
	grader          = ws.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

type Server struct {
	*lsp.AbsServer
	SpecificFileCfg    *config.Config
	DefaultFileCfg     *config.Config
	WSCfg              *config.Config
	dbConn             *database.DBConnection
	curDBCfg           *database.DBConfig
	curDBName          string
	curConnectionIndex int

	// The initOptionDBConfig is an optional param
	// sent by the client as part of the LSP InitializationOptions
	// payload. If non-nil, the server will ignore all
	// other configuration sources (workspace and user).
	initOptionDBConfig *database.DBConfig

	worker  *database.Worker
	files   map[lsp.DocumentURI]*File
	options []jsonrpc2.ConnOpt
}

func NewServer(conf *config.Config, options ...jsonrpc2.ConnOpt) *Server {
	worker := database.NewWorker()
	server := &Server{
		DefaultFileCfg: conf,
		files:          make(map[lsp.DocumentURI]*File),
		worker:         worker,
		options:        options,
	}
	server.options = append(server.options, jsonrpc2.LogMessages(server))
	return server
}

func (that *Server) Start() error {
	if nil == that.DefaultFileCfg {
		cfg, err := config.GetDefaultConfig()
		if err != nil && !errors.Is(config.ErrNotFoundConfig, err) {
			return fmt.Errorf("cannot read default config, %w", err)
		}
		that.DefaultFileCfg = cfg
	}
	that.worker.Start()
	return nil
}

func (that *Server) Stop() error {
	if err := that.dbConn.Close(); err != nil {
		return err
	}
	that.worker.Stop()
	return nil
}

func (that *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if err := lsp.PanicEfx(func() error {
		c, err := grader.Upgrade(writer, request, nil)
		if nil != err {
			return err
		}
		defer func() {
			lsp.Catch(c.Close())
		}()
		wc := jsonrpc2.NewConn(
			request.Context(),
			websocket.NewObjectStream(c),
			jsonrpc2.HandlerWithError(that.ServeRPC),
			that.options...)
		<-wc.DisconnectNotify()
		return c.WriteControl(ws.CloseMessage, []byte{}, time.Now().Add(time.Second))
	}); nil != err {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (that *Server) ServeSTD(ctx context.Context) {
	h := jsonrpc2.HandlerWithError(that.ServeRPC)
	wc := jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(&stdPip{}, jsonrpc2.VSCodeObjectCodec{}),
		h,
		that.options...,
	)
	<-wc.DisconnectNotify()
}
func (that *Server) ServeRPC(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) (any, error) {
	return lsp.Handle(ctx, that, conn, r)
}

func (that *Server) Printf(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
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
