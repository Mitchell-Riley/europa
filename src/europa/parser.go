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
	"fmt"
	"io"
	"os"
	"strconv"
)

type Lexer struct {
	input   string
	current string
	next    string
	line    int
}

type ILexer interface {
	Consume()
	CurrentChar() byte
	NextChar()
	Lex()
	ParseIdent()
	ParseNumber()
	ParseArguments() []interface{}
	ParseExpression() []interface{}
}

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
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
	r := &Lexer{
		input:   str,
		current: "",
		next:    "",
		line:    1,
	}
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
	for i = 0; i < inlen && isLetter(lex.input[i]); i++ {
	}
	lex.next = lex.input[0:i]
	lex.input = lex.input[i:]
}
func (lex *Lexer) ParseNumber() {
	inlen := len(lex.input)
	var i int
	for i = 0; i < inlen && isDigit(lex.input[i]); i++ {
	}
	lex.next = lex.input[0:i]
	lex.input = lex.input[i:]
}
func (lex *Lexer) ParseString() {
	inlen := len(lex.input)
	var i int
	for i = 1; i < inlen && lex.input[i] != '"'; i++ {
	}
	lex.next = lex.input[0 : i+1]
	lex.input = lex.input[i+1:]
}
func (lex *Lexer) ParseArguments() *[]interface{} {
	args := new([]interface{})
	arg := new([]interface{})
	for lex.current != "" {
		if lex.current == ")" {
			if len(*arg) > 0 {
				*args = append(*args, arg)
			}
			lex.Consume()
		} else if lex.current == "," {
			*args = append(*args, new([]interface{}))
			arg = new([]interface{})
			lex.Consume()
		} else {
			arg = lex.ParseExpression()
		}
	}

	return args
}
func (lex *Lexer) ParseExpression() *[]interface{} {
	tree := new([]interface{})
	for lex.current != "" {
		if lex.current == "," {
			break
		} else if lex.current == "(" {
			lex.Consume()
			args := lex.ParseArguments()
			if lex.current == ")" {
				if len(*tree) == 0 {
					*tree = append(*tree, NewMessage("", new([]interface{})))
				}

				if len(*((*tree)[len(*tree)-1].(IMessage).GetArguments())) > 0 {
					fmt.Println("*** Arguments are empty")
					*tree = append(*tree, NewMessage("", new([]interface{})))
				}

				(*tree)[len(*tree)-1].(IMessage).SetArguments(args)
			} else {
				fmt.Println("Syntax Error: ')' expected")
			}
			fmt.Println(len(*args))
		} else if lex.current == ")" {
			break
		} else {
			fmt.Println("*** (ParseExpression) / fallback (line: " + strconv.Itoa(lex.line) + ") -- lex.current = " + lex.current + "; lex.next = " + lex.next)

			*tree = append(*tree, NewMessage(lex.current, new([]interface{})))
			lex.Consume()
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
		lex.line++
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

func Parse(state IState, filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	var result []byte
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != io.EOF {
			break
		}
	}

	lex := NewLexer(string(result))
	expr := lex.ParseExpression()
	state.EvaluateTree(*expr)

	return nil
}

func ParseString(state IState, code string) {
	lex := NewLexer(code)
	expr := lex.ParseExpression()
	state.EvaluateTree(*expr)
}
