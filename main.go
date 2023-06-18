/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package main

import (
	"github.com/ducesoft/ulsp/internal/config"
	"github.com/ducesoft/ulsp/server"
	"github.com/rs/zerolog/log"
)

func main() {
	s := &server.Server{}
	if err := s.Start("0.0.0.0:8888", &config.Config{}); nil != err {
		log.Error().Err(err)
	}
}
