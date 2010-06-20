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

	SetName(string)
	SetArguments([]*Message)
	SetNext(*Message)
}

func (msg *Message) Clone() IObject {
	r := new(Message)
	r.proto = msg
	return r
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
