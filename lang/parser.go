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
	s string
	current string
	positions vector.Vector
	tokens vector.Vector
}

type Token struct {
	name string
	tokenType TokenType
	charNum int
	lineNum int
	next *Token
	error string
}

type IToken interface {
	SetName(string)
	GetName() string
	
	SetType(TokenType)
	GetType() TokenType
	
	SetNext(*Token)
	GetNext() *Token
	
	SetError(string)
	GetError() string
}

func NewToken(name string) *Token {
	return &Token{name, TK_NONE, -1, -1, nil, ""}
}

func (tok *Token) SetName(name string) {
	tok.name = name
}

func (tok *Token) GetName() string {
	return tok.name
}

func (tok *Token) SetType(tt TokenType) {
	tok.tokenType = tt
}

func (tok *Token) GetType() TokenType {
	return tok.tokenType
}

func (tok *Token) SetNext(token *Token) {
	tok.next = token
}

func (tok *Token) GetNext() *Token {
	return tok.next
}

func (tok *Token) SetError(error string) {
	tok.error = error
}

func (tok *Token) GetError() string {
	return tok.error
}
