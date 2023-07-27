/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package cause

var (
	SystemError     = &BindError{Code: "E0000000520", Format: "%v"}
	ErrNoConnection = &BindError{Code: "E0000000520", Format: "Not found database connection config"}
	FileNotFound    = &BindError{Code: "E0000000520", Format: "File %v not found"}
	DBConfigNotSet  = &BindError{Code: "E0000000520", Format: "DBConfig not set"}
)
