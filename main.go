/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package main

import (
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/log"
	"github.com/ducesoft/ulsp/server"
)

func main() {
	s := &server.Server{}
	log.Catch(s.ListenAndServe("0.0.0.0:8888", &config.Config{}))
}
