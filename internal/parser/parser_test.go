/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package parser

import (
	"context"
	"testing"
)

func TestParseKeywords(t *testing.T) {
	t.Log(ParseKeywords(context.TODO(), `
	select count(*);
	`))
	t.Log(ParseKeywords(context.TODO(), `
	select count(*) from x where id in(select id from c where 1=1);
	`))
}
