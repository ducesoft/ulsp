/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package cause

import (
	"bytes"
	"flag"
	"fmt"
	"runtime"
	"strings"
)

type Codeable interface {

	// GetCode Get the code.
	GetCode() string

	// GetMessage Get the message.
	GetMessage() string
}

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

func Error(err error) error {
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

type MultiError struct {
	errs []error
}

func (that *MultiError) Append(err error) {
	if nil != err {
		that.errs = append(that.errs, err)
	}
}

func (that *MultiError) AsError() error {
	if len(that.errs) < 1 {
		return nil
	}
	return that
}

func (that *MultiError) Error() string {
	b := &bytes.Buffer{}
	for _, err := range that.errs {
		b.WriteString(err.Error())
		b.WriteString("\n")
	}
	return b.String()
}

type BindError struct {
	Code   string
	Format string
}

func (that *BindError) New(args ...any) error {
	return &Cause{
		code: that.Code,
		at:   Caller(3),
		err:  fmt.Errorf(that.Format, args...),
	}
}

func (that *BindError) Is(err error) bool {
	if b, ok := err.(Codeable); ok {
		return b.GetCode() == that.Code
	}
	return false
}
