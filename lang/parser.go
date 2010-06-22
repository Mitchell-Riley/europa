/*******************************************************************
 * Europa Programming Language
 * Copyright (C) 2010, Jeremy Tregunna, All Rights Reserved.
 *
 * This software project, which includes this module, is protected
 * under Canadian copyright legislation, as well as international
 * treaty. It may be used only as directed by the copyright holder.
 * The copyright holder hereby consents to usage and distribution
 * based on the terms and conditions of the MIT license, which may
 * be found in the LICENSE.MIT file included in this distribution.
 *******************************************************************
 * Project: Europa Programming Language
 * File: parser.go
 * Description: Parsing subsystem. Derived in large part from the
 *              go-parsec library.
 ******************************************************************/

package parser

import (
	"container/vector"
	//"./europa"
)

type TokenType int

const (
	TK_NONE = iota
	TK_LPAREN
	TK_COMMA
	TK_RPAREN
	TK_DQUOTE
	TK_TQUOTE
	TK_IDENT
	TK_TERM
	TK_COMMENT
	TK_NUMBER
	TK_HEX
)

type Lexer struct {
	input string
	current *Token
	next *Token
}

type ILexer interface {
}

type Token struct {
	name string
	arguments vector.Vector
}

type IToken interface {
	SetName(string)
	GetName() string
	GetArguments() vector.Vector
	Equal(IToken) bool
}

func NewToken(name string) *Token {
	return &Token{name: name}
}
func (tok *Token) SetName(name string) {
	tok.name = name
}
func (tok *Token) GetName() string {
	return tok.name
}
func (tok *Token) GetArguments() vector.Vector {
	return tok.arguments
}
func (tok *Token) Equal(token IToken) bool {
	return tok.name == token.GetName()
}

func NewLexer(str string) *Lexer {
	r := new(Lexer)
	r.input = str
	r.current = nil
	r.next = nil
	return r
}
