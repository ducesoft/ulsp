/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package lsp

import (
	"context"
	"errors"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/internal/database"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/log"
	"io"
	"time"
)

func withContext(ctx context.Context, cfg *config.Config, wc *jsonrpc2.Conn) Context {
	if c, ok := ctx.(Context); ok {
		return c
	}
	return &sct{
		id:        "",
		db:        "",
		ctx:       ctx,
		cfg:       cfg,
		conn:      nil,
		repo:      nil,
		dfg:       nil,
		messenger: NewLspMessenger(wc),
		worker:    database.NewWorker(),
		files:     make(map[DocumentURI]*Filer),
	}
}

func ContextWith(ctx context.Context) Context {
	if c, ok := ctx.(Context); ok {
		return c
	}
	panic("unexpected runtime environment.")
}

type Context interface {
	context.Context
	io.Closer
	ID() string
	Messanger() Messenger
	Init(dfg *config.DBConfig) error
	DB() *database.DBCache
	Config() *config.Config
	Sync(f *Filer)
	Open(uri DocumentURI) (*Filer, error)
	DBConfig() (*config.DBConfig, error)
	Repository() (database.DBRepository, error)
}

type sct struct {
	id        string
	db        string
	ctx       context.Context
	cfg       *config.Config
	messenger Messenger
	worker    *database.Worker
	files     map[DocumentURI]*Filer
	dfg       *config.DBConfig       // optional
	conn      *database.DBConnection // optional
	repo      database.DBRepository  // optional
}

func (that *sct) Deadline() (deadline time.Time, ok bool) {
	return that.ctx.Deadline()
}

func (that *sct) Done() <-chan struct{} {
	return that.ctx.Done()
}

func (that *sct) Err() error {
	return that.ctx.Err()
}

func (that *sct) Value(key any) any {
	return that.ctx.Value(key)
}

func (that *sct) Config() *config.Config {
	return that.cfg
}

func (that *sct) Open(uri DocumentURI) (*Filer, error) {
	if f, ok := that.files[uri]; ok {
		return f, nil
	}
	return nil, cause.Errorf("%s not found", string(uri))
}

func (that *sct) DB() *database.DBCache {
	return that.worker.Cache()
}

func (that *sct) DBConfig() (*config.DBConfig, error) {
	return that.dfg, nil
}

func (that *sct) Repository() (database.DBRepository, error) {
	return that.repo, nil
}

func (that *sct) Sync(f *Filer) {
	if x, ok := that.files[f.URI]; ok {
		x.LanguageID = f.LanguageID
		x.Text = f.Text
	} else {
		that.files[f.URI] = f
	}
}

func (that *sct) Messanger() Messenger {
	return that.messenger
}

func (that *sct) Init(dfg *config.DBConfig) error {
	that.dfg = dfg
	// Initialize database database connection
	// NOTE: If no connection is found at this point,
	// it is possible that the connection settings are sent to workspace config
	// so don't make an error
	if err := that.Reconnection(that); err != nil {
		if !errors.Is(cause.ErrNoConnection, err) {
			if err = that.Messanger().ShowInfo(that, err.Error()); err != nil {
				log.Error(that, "send info, %s", err.Error())
				return err
			}
		} else {
			log.Error(that, "send err, %s", err.Error())
			if err = that.Messanger().ShowError(that, err.Error()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (that *sct) ID() string {
	return that.id
}

func (that *sct) Close() error {
	if nil != that.worker {
		that.worker.Stop()
	}
	if nil != that.conn {
		return that.conn.Close()
	}
	return nil
}

func (that *sct) Reconnection(ctx context.Context) (err error) {
	if nil == that.dfg {
		return cause.Errorf("not found database connection config")
	}
	that.conn, err = database.Open(that.dfg)
	if nil != err {
		return err
	}
	repo, err := database.CreateRepository(that.dfg.Driver, that.conn.Conn)
	if nil != err {
		return err
	}
	if err = that.worker.ReCache(ctx, repo); nil != err {
		return err
	}
	return nil
}
