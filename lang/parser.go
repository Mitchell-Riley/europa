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
	"strings"
	"unicode"
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
	input *strings.Reader
	current *strings.Reader
	next *strings.Reader
}
type ILexer interface {
	Consume()
	CurrentRune() int
	NextRune()
	Lex()
	ParseArguments() vector.Vector
	ParseExpression() vector.Vector
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

func NewLexer(str *strings.Reader) *Lexer {
	r := new(Lexer)
	r.input = str
	r.current = strings.NewReader("")
	r.next = strings.NewReader("")
	r.Consume()
	r.Consume()
	return r
}
func (lex *Lexer) Consume() {
	lex.current = lex.next
	lex.Lex()
}
func (lex *Lexer) CurrentRune() int {
	r, _, _ := lex.input.ReadRune()
	return r
}
func (lex *Lexer) NextRune() {
	tmp := (*lex.input)[1:]
	lex.input = &tmp
}
func (lex *Lexer) ParseIdent() {
	r, size, _ := lex.input.ReadRune()
	var i int
	for i = 0; i < len(*lex.input) && unicode.IsLetter(r); i += size {}
	tmp := (*lex.input)[0:i]
	lex.next = &tmp
	tmp = (*lex.input)[i:]
	lex.input = &tmp
}
func (lex *Lexer) ParseNumber() {
	for i, r := range *lex.input {
		if !unicode.IsDigit(r) { break }
		tmp := (*lex.input)[0:i]
		lex.next = &tmp
		tmp = (*lex.input)[i:]
		lex.input = &tmp
	}
}
func (lex *Lexer) ParseArguments() vector.Vector {
	var (
		args vector.Vector
		arg vector.Vector
	)
	arg = nil
	lparenRune := strings.NewReader("(")
	commaRune := strings.NewReader(",")
	for lex.current != nil {
		if lex.current == lparenRune {
			if arg != nil {
				args.Push(arg)
			}
		} else if lex.current == commaRune {
			args.Push(arg)
			arg = nil
			lex.Consume()
		} else {
			arg = lex.ParseExpression()
		}
	}
	
	return args
}
func (lex *Lexer) ParseExpression() vector.Vector {
	var tree vector.Vector
}
func (lex *Lexer) Lex() {
	newlineRune, _, _ := strings.NewReader("\n").ReadRune()
	spaceRune, _, _ := strings.NewReader(" ").ReadRune()
	
	if lex.input == strings.NewReader("") {
		lex.next = nil
	} else if lex.CurrentRune() == newlineRune {
		lex.next = strings.NewReader(";")
		lex.NextRune()
	} else if lex.CurrentRune() == spaceRune {
		lex.NextRune()
		lex.Lex()
	} else if unicode.IsLetter(lex.CurrentRune()) {
		lex.ParseIdent()
	} else if unicode.IsDigit(lex.CurrentRune()) {
		lex.ParseNumber()
	} else {
		tmp, _ := lex.input.ReadByte()
		lex.next = strings.NewReader(string(tmp))
		lex.NextRune()
	}
}
