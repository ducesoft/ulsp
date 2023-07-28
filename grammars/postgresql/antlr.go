/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package mysql

//go:generate curl -OL https://github.com/antlr/grammars-v4/raw/master/sql/postgresql/PostgreSQLLexer.g4
//go:generate curl -OL https://github.com/antlr/grammars-v4/raw/master/sql/postgresql/PostgreSQLParser.g4
//go:generate mvn antlr4:antlr4 -X -Dlanguage=Go
