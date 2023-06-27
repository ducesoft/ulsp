/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package log

import (
	"context"
	zog "github.com/rs/zerolog/log"
)

func SetLogger(g Logger) {
	uog = g
}

var uog Logger = new(zero)

type Logger interface {
	Info(ctx context.Context, format string, args ...interface{})

	Warn(ctx context.Context, format string, args ...interface{})

	Error(ctx context.Context, format string, args ...interface{})

	Debug(ctx context.Context, format string, args ...interface{})

	Fatal(ctx context.Context, format string, args ...interface{})

	Catch(err error)
}

func Info(ctx context.Context, format string, args ...interface{}) {
	uog.Info(ctx, format, args...)
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	uog.Warn(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	uog.Error(ctx, format, args...)
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	uog.Debug(ctx, format, args...)
}

func Fatal(ctx context.Context, format string, args ...interface{}) {
	uog.Fatal(ctx, format, args...)
}

func Catch(err error) {
	uog.Catch(err)
}

type zero struct {
}

func (that *zero) Info(ctx context.Context, format string, args ...interface{}) {
	zog.Info().Msgf(format, args...)
}

func (that *zero) Warn(ctx context.Context, format string, args ...interface{}) {
	zog.Warn().Msgf(format, args...)
}

func (that *zero) Error(ctx context.Context, format string, args ...interface{}) {
	zog.Error().Msgf(format, args...)
}

func (that *zero) Debug(ctx context.Context, format string, args ...interface{}) {
	zog.Debug().Msgf(format, args...)
}

func (that *zero) Fatal(ctx context.Context, format string, args ...interface{}) {
	zog.Fatal().Msgf(format, args...)
}

func (that *zero) Catch(err error) {
	if nil != err {
		zog.Error().Msgf(err.Error())
	}
}
