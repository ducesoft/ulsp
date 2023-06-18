/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package server

import (
	"context"
	"github.com/ducesoft/ulsp/internal/config"
	"github.com/ducesoft/ulsp/internal/service"
	"github.com/ducesoft/ulsp/lsp"
	"net/http"
)

type Server struct {
	server *http.Server
	serves *service.Server
}

func (that *Server) Start(address string, conf *config.Config) error {
	that.serves = service.NewServer(conf)
	that.server = &http.Server{
		Addr:    address,
		Handler: that.serves,
	}
	if err := that.serves.Start(); nil != err {
		return err
	}
	return that.server.ListenAndServe()
}

func (that *Server) Stop(ctx context.Context) error {
	if nil != that.serves {
		lsp.Catch(that.serves.Stop())
	}
	if nil != that.server {
		lsp.Catch(that.server.Shutdown(ctx))
	}
	return nil
}
