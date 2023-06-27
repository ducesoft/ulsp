/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package server

import (
	"context"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/serves"
	"github.com/ducesoft/ulsp/log"
	"net/http"
)

var _ http.Handler = new(Server)

type Server struct {
	server *http.Server
	serves *serves.Server
}

func (that *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if nil == that.serves {
		http.NotFound(writer, request)
		return
	}
	that.serves.ServeHTTP(writer, request)
}

func (that *Server) Start(conf *config.Config) error {
	that.serves = serves.NewServer(conf)
	return that.serves.Start()
}

func (that *Server) ListenAndServe(address string, conf *config.Config) (err error) {
	if err = that.Start(conf); nil != err {
		return err
	}
	that.server = &http.Server{
		Addr:    address,
		Handler: that.serves,
	}
	return that.server.ListenAndServe()
}

func (that *Server) Stop(ctx context.Context) error {
	if nil != that.serves {
		log.Catch(that.serves.Stop())
	}
	if nil != that.server {
		log.Catch(that.server.Shutdown(ctx))
	}
	return nil
}
