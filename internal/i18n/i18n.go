/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package i18n

import (
	"context"
	"embed"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
	"strings"
)

func init() {
	if err := importLocales(); nil != err {
		panic(err)
	}
}

const i18nContext = "i18n"

//go:embed locale.*.yml
var locales embed.FS
var printers = map[string]*message.Printer{}

func importLocales() error {
	fs, err := locales.ReadDir("")
	if nil != err {
		return err
	}
	for _, f := range fs {
		if err = importLocale(f.Name()); nil != err {
			return err
		}
	}
	return nil
}

func importLocale(name string) error {
	b, err := locales.ReadFile(name)
	if nil != err {
		return err
	}
	lag := language.Make(strings.Split(name, ".")[1])
	var kv map[string]any
	if err = yaml.Unmarshal(b, &kv); nil != err {
		return err
	}
	for k, v := range kv {
		if err = message.SetString(lag, k, fmt.Sprintf("%v", v)); nil != err {
			return err
		}
	}
	printers[lag.String()] = message.NewPrinter(lag)
	return nil
}

func Context(ctx context.Context, tag language.Tag) context.Context {
	return context.WithValue(ctx, i18nContext, tag.String())
}

func Sprintf(ctx context.Context, format string, args ...any) string {
	t, ok := ctx.Value(i18nContext).(string)
	if !ok {
		return fmt.Sprintf(format, args...)
	}
	if p, ojbk := printers[t]; ojbk {
		return p.Sprintf(format, args...)
	} else {
		return fmt.Sprintf(format, args...)
	}
}
