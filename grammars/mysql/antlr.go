/*
 * Copyright (c) 2000, 2099, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package mysql

//go:generate curl -OL https://github.com/mysql/mysql-workbench/blob/8.0/library/parsers/grammars/MySQLParser.g4
//go:generate curl -OL https://github.com/mysql/mysql-workbench/blob/8.0/library/parsers/grammars/MySQLLexer.g4
//go:generate mvn antlr4:antlr4 -X -Dlanguage=Go
