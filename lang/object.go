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
 * File: object.go
 * Description: Defines the base object type.
 ******************************************************************/

package main

type Object struct {
	/* Our parent. */
	proto IObject

	/* Slot table. Simple string->object mapping. */
	slots map[string] IObject

	/* Are we a locals object? */
	locals bool

	/* Are we activatable? */
	activatable bool
}

type IObject interface {
	Clone() IObject
	SetSlot(string, IObject)
	SetLocals(bool)
	GetActivatable() bool
	SetActivatable(bool)
	GetSlot(string) (IObject, IObject)
	Perform(IObject, IMessage) IObject
	Forward(IObject, IMessage) IObject
	Activate(IObject, IObject, IMessage, IObject) IObject
}

func (obj *Object) Clone() IObject {
	r := new(Object)
	r.proto = obj.proto
	r.locals = false
	return r
}

func (obj *Object) SetSlot(key string, value IObject) {
	obj.slots[key] = value
}

func (obj *Object) SetLocals(val bool) {
	obj.locals = val;
}

func (obj *Object) GetActivatable() bool {
	return obj.activatable
}

func (obj *Object) SetActivatable(val bool) {
	obj.activatable = val
}


func (obj *Object) GetSlot(key string) (v IObject, ctx IObject) {
	var ok bool

	if v, ok = obj.slots[key]; ok {
		ctx = obj
		return v, ctx
	}

	if v == nil {
		return nil, nil
	}

	return obj.proto.GetSlot(key)
}

func (obj *Object) Perform(locals IObject, msg IMessage) IObject {
	v, ctx := obj.GetSlot(msg.GetName())

	if v != nil && ctx != nil {
		return v.Activate(obj, locals, msg, ctx)
	}

	return obj.Forward(locals, msg)
}

func (obj *Object) Forward(locals IObject, msg IMessage) IObject {
	v, ctx := obj.GetSlot("forward")

	if v != nil && ctx != nil {
		if obj.locals {
			if delegate, ok := obj.slots["self"]; ok {
				return delegate.Perform(locals, msg)
			}
			return nil
		} else {
			return v.Activate(obj, locals, msg, ctx)
		}
	}
	return nil
}

func (obj *Object) Activate(target IObject, locals IObject, msg IMessage, ctx IObject) IObject {
	// TODO: Check if the object is activatable, and if so...call it's activate function, otherwise return obj.
	return obj
}
