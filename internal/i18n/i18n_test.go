/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package i18n

import (
	"context"
	"golang.org/x/text/language"
	"testing"
)

func TestI18N(t *testing.T) {
	ctx := Context(context.Background(), language.SimplifiedChinese)
	t.Log(Sprintf(ctx, "Execute Query"))
}
