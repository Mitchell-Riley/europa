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
 * File: msg.go
 * Description: Messages
 ******************************************************************/

package main

type Message struct {
	*Object

	/* Name of the message */
	name string

	/* List of arguments */
	args []*Message

	/* Next message in the chain */
	next *Message
}

type IMessage interface {
	IObject

	GetName() string
	SetName(string)
	SetArguments([]*Message)
	SetNext(*Message)

	PerformOn(IObject, IObject) IObject
}

func (msg *Message) Clone() IObject {
	r := new(Message)
	r.proto = msg
	return r
}

func (msg *Message) GetName() string {
	return msg.name
}

func (msg *Message) SetName(str string) {
	msg.name = str
}

func (msg *Message) SetArguments(args []*Message) {
	msg.args = args
}

func (msg *Message) SetNext(next *Message) {
	msg.next = next
}

func (msg *Message) PerformOn(locals IObject, target IObject) IObject {
	var cached = target
	var m = msg
	var result IObject

	for ; m.next != nil; m = m.next {
		if m.name == ";" {
			target = cached;
		} else {
			result = target.Perform(locals, m)
		}
	}

	return result;
}

func (msg *Message) Activate(target IObject, locals IObject, m IMessage, ctx IObject) IObject {
	return msg.PerformOn(locals, locals)
}
