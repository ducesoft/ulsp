/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package log

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
)

type Cause struct {
	code string
	at   string
	err  error
}

func (that *Cause) Error() string {
	return fmt.Sprintf("%s:%s", that.GetMessage(), that.at)
}

func (that *Cause) GetCode() string {
	return that.code
}

func (that *Cause) GetMessage() string {
	if nil == that.err {
		return "Unknown"
	}
	return that.err.Error()
}

func Errors(err error) error {
	if nil == err {
		return err
	}
	if cause, ok := err.(*Cause); ok {
		return cause
	}
	return &Cause{
		code: "E0000000520",
		at:   Caller(2),
		err:  err,
	}
}

func Errorf(format string, args ...interface{}) error {
	return &Cause{
		code: "E0000000520",
		at:   Caller(2),
		err:  fmt.Errorf(format, args...),
	}
}

func Caller(skip int) string {
	_, name, line, _ := runtime.Caller(skip)
	if isTest() {
		return fmt.Sprintf("%s:%d", name, line)
	}
	return fmt.Sprintf("%s:%d", name[strings.LastIndex(name, "/")+1:], line)
}

// IsTest is golang testing.
func isTest() bool {
	return nil != flag.Lookup("test.v")
}
