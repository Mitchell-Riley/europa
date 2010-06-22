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
	input *strings.Reader
	current *strings.Reader
	next *strings.Reader
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
	lparenRune := strings.NewReader("(")
	commaRune := strings.NewReader(",")
	for lex.current != nil {
		if lex.current == lparenRune {
			if len(arg) > 0 {
				args.Push(arg)
			}
		} else if lex.current == commaRune {
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
	commaRune := strings.NewReader(",")
	lparenRune := strings.NewReader("(")
	rparenRune := strings.NewReader(")")
	for lex.current != nil {
		if lex.current == commaRune {
			break
		} else if lex.current == rparenRune {
			break
		} else if lex.current == lparenRune {
			lex.Consume()
			args := lex.ParseArguments()
			if lex.current == rparenRune {
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
			// XXX: Don't know if the cast works as expected
			tree.Push(NewMessage((string)(*lex.current), new(vector.Vector)))
		}
	}
	
	return tree
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
