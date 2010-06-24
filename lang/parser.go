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
	"os"
	"bytes"
	"strconv"
	"container/vector"
)

type Lexer struct {
	input string
	current string
	next string
	line int
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

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') ||
	   (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func NewLexer(str string) *Lexer {
	r := new(Lexer)
	r.input = str
	r.current = ""
	r.next = ""
	r.line = 1
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
	inlen := len(lex.input)
	var i int
	for i = 0; i < inlen && isLetter(lex.input[i]); i++ {}
	lex.next = lex.input[0:i]
	lex.input = lex.input[i:]
}
func (lex *Lexer) ParseNumber() {
	inlen := len(lex.input)
	var i int
	for i = 0; i < inlen && isDigit(lex.input[i]); i++ {}
	lex.next = lex.input[0:i]
	lex.input = lex.input[i:]
}
func (lex *Lexer) ParseString() {
	inlen := len(lex.input)
	var i int
	for i = 1; i < inlen && lex.input[i] != '"'; i++ {}
	lex.next = lex.input[0:i + 1]
	lex.input = lex.input[i + 1:]
}
func (lex *Lexer) ParseArguments() *vector.Vector {
	args := new(vector.Vector)
	arg := new(vector.Vector)
	for lex.current != "" {
		if lex.current == ")" {
			if arg.Len() > 0 {
				args.Push(arg)
			}
			lex.Consume()
		} else if lex.current == "," {
			args.Push(arg)
			arg = new(vector.Vector)
			lex.Consume()
		} else {
			arg = lex.ParseExpression()
		}
	}
	
	return args
}
func (lex *Lexer) ParseExpression() *vector.Vector {
	tree := new(vector.Vector)
	for lex.current != "" {
		if lex.current == "," {
			break
		} else if lex.current == ")" {
			break
		} else if lex.current == "(" {
			lex.Consume()
			args := lex.ParseArguments()
			if lex.current == ")" {
				if tree.Len() == 0 {
					tree.Push(NewMessage("", new(vector.Vector)))
				}
				
				if tree.Last().(IMessage).GetArguments().Len() > 0 {
					tree.Push(NewMessage("", new(vector.Vector)))
				}
				
				tree.Last().(IMessage).SetArguments(args)
				lex.Consume()
			} else {
				println("Syntax Error: ')' expected")
			}
		} else {
			println("*** (ParseExpression) / fallback (line: " + strconv.Itoa(lex.line) + ") -- lex.current = " + lex.current + "; lex.next = " + lex.next)
			tree.Push(NewMessage(lex.current, new(vector.Vector)))
			lex.Consume()
		}
	}
	
	return tree
}
func (lex *Lexer) Lex() {
	if lex.input == "" {
		lex.next = ""
	} else if lex.CurrentChar() == '\n' {
		lex.line += 1
		lex.next = ";"
		lex.NextChar()
	} else if lex.CurrentChar() == ' ' {
		lex.NextChar()
		lex.Lex()
	} else if isLetter(lex.CurrentChar()) {
		lex.ParseIdent()
	} else if isDigit(lex.CurrentChar()) {
		lex.ParseNumber()
	} else if lex.CurrentChar() == '"' {
		lex.ParseString()
	} else {
		lex.next = string(lex.CurrentChar())
		lex.NextChar()
	}
}

func Parse(filename string) os.Error {
	f, err := os.Open(filename, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	
	var result []byte
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf[0:])
		result = bytes.Add(result, buf[0:n])
		if err != os.EOF {
			break
		}
	}
	
	lex := NewLexer(string(result))
	expr := lex.ParseExpression()
	println(strconv.Itoa(expr.Len()) + " expressions")
	
	return nil
}

func ParseString(code string) {
	lex := NewLexer(code)
	lex.ParseExpression()
}
