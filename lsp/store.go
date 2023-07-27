/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package lsp

import (
	"context"
	"github.com/ducesoft/ulsp/cause"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"sync"
)

func NewStore() *Store {
	return &Store{sessions: map[string]Context{}}
}

type Store struct {
	sessions map[string]Context
	lock     sync.RWMutex
}

func (that *Store) Open(ctx context.Context, cfg *config.Config, wc *jsonrpc2.Conn) Context {
	stx := withContext(ctx, cfg, wc)
	that.lock.Lock()
	defer that.lock.Unlock()
	that.sessions[stx.ID()] = stx
	return stx
}

func (that *Store) Release(ctx Context) {
	that.lock.Lock()
	defer that.lock.Unlock()
	delete(that.sessions, ctx.ID())
}

func (that *Store) Close() error {
	that.lock.RLock()
	defer that.lock.RUnlock()
	e := &cause.MultiError{}
	for _, s := range that.sessions {
		if err := s.Close(); nil != err {
			e.Append(err)
		}
	}
	return e.AsError()
}
