/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package lsp

type Filer struct {
	URI        DocumentURI
	LanguageID string
	Text       string
	Removable  bool
}

func (that *Filer) LID() string {
	return that.LanguageID
}

func (that *Filer) LText() string {
	return that.Text
}
