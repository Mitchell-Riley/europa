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

package europa

import (
	"strings"
	"unicode"
	"container/vector"
)

type Lexer struct {
	input string
	current string
	next string
}
type ILexer interface {
	Consume()
	CurrentRune() int
	NextRune()
	Lex()
	ParseIdent()
	ParseNumber()
	ParseArguments() vector.Vector
	ParseExpression() vector.Vector
}

func isAlnum(text string) bool {
	for i, c := range text {
		if i % 1 == 0 && !unicode.IsLetter(c) {
			return false
		} else {
			if !(unicode.IsLetter(c) || unicode.IsDigit(c)) {
				return false
			}
		}
	}
	return true
}

func isDigit(text string) bool {
	for _, c := range text {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func NewLexer(str string) *Lexer {
	r := new(Lexer)
	r.input = str
	r.current = ""
	r.next = ""
	r.Consume()
	r.Consume()
	return r
}
func (lex *Lexer) Consume() {
	lex.current = lex.next
	lex.Lex()
}
func (lex *Lexer) CurrentChar() byte {
	return lex.input[0]
}
func (lex *Lexer) NextChar() {
	lex.input = lex.input[1:]
}
func (lex *Lexer) ParseIdent() {
	s := strings.Split(lex.input, " ", 1)[0]
	if isAlnum(s) {
		lex.next = lex.input[0:len(s)]
		lex.input = lex.input[len(s):]
	}
}
func (lex *Lexer) ParseNumber() {
	s := strings.Split(lex.input, " ", 1)[0]
	if isDigit(s) {
		lex.next = lex.input[0:len(s)]
		lex.input = lex.input[len(s):]
	}
}
func (lex *Lexer) ParseArguments() vector.Vector {
	var (
		args vector.Vector
		arg vector.Vector
	)
	for lex.current != "" {
		if lex.current == "(" {
			if len(arg) > 0 {
				args.Push(arg)
			}
		} else if lex.current == "," {
			args.Push(arg)
			arg = make(vector.Vector, 0)
			lex.Consume()
		} else {
			arg = lex.ParseExpression()
		}
	}
	
	return args
}
func (lex *Lexer) ParseExpression() vector.Vector {
	var tree vector.Vector
	for lex.current != "" {
		if lex.current == "," {
			break
		} else if lex.current == ")" {
			break
		} else if lex.current == "(" {
			lex.Consume()
			args := lex.ParseArguments()
			if lex.current == ")" {
				if len(tree) == 0 {
					tree.Push(NewMessage("", new(vector.Vector)))
				}
				
				if tree.Last().(IMessage).GetArguments().Len() > 0 {
					tree.Push(NewMessage("", new(vector.Vector)))
				}
				
				tree.Last().(IMessage).SetArguments(&args)
				lex.Consume()
			} else {
				println("Syntax Error: ')' expected")
			}
		} else {
			tree.Push(NewMessage(lex.current, new(vector.Vector)))
		}
	}
	
	return tree
}
func (lex *Lexer) Lex() {
	if lex.input == "" {
		lex.next = ""
	} else if lex.CurrentChar() == '\n' {
		lex.next = ";"
		lex.NextChar()
	} else if lex.CurrentChar() == ' ' {
		lex.NextChar()
		lex.Lex()
	} else if isAlnum(string(lex.CurrentChar())) {
		lex.ParseIdent()
	} else if isDigit(string(lex.CurrentChar())) {
		lex.ParseNumber()
	} else {
		lex.next = string(lex.CurrentChar())
		lex.NextChar()
	}
}

func Parse(str string) {
	lex := NewLexer(str)
	for _, e := range lex.ParseExpression() {
		println(e)
	}
}
