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

package europa

const DEFAULT_SLOTS_SIZE = 8

type Object struct {
	/* VM State */
	state *IState

	/* Our parent. */
	proto IObject

	/* Slot table. Simple string->object mapping. */
	slots map[string] IObject

	/* Are we a locals object? */
	locals bool

	/* Are we activatable? */
	activatable bool

	/* What about if we've scanned this object already in lookup? */
	scanned bool
}

type IObject interface {
	Clone() IObject
	GetState() IState
	SetSlot(string, IObject)
	SetProto(IObject)
	SetLocals(bool)
	GetActivatable() bool
	SetActivatable(bool)
	GetSlot(string) (IObject, IObject)
	Perform(IObject, IMessage) IObject
	Forward(IObject, IMessage) IObject
	Activate(IObject, IObject, IMessage, IObject) IObject
}

func NewObject(state IState, proto IObject, locals bool, activatable bool, scanned bool) IObject {
	r := new(Object)
	r.state = &state
	r.slots = make(map[string]IObject, DEFAULT_SLOTS_SIZE)
	r.proto = proto
	r.locals = locals
	r.activatable = activatable
	r.scanned = scanned
	return r
}

func (obj *Object) Clone() IObject {
	r := new(Object)
	r.proto = obj.proto
	r.slots = make(map[string]IObject, DEFAULT_SLOTS_SIZE)
	r.locals = false
	return r
}

func (obj *Object) GetState() IState {
	return *obj.state
}

func (obj *Object) SetSlot(key string, value IObject) {
	obj.slots[key] = value
}

func (obj *Object) SetProto(proto IObject) {
	obj.proto = proto
}

func (obj *Object) SetLocals(val bool) {
	obj.locals = val
}

func (obj *Object) GetActivatable() bool {
	return obj.activatable
}

func (obj *Object) SetActivatable(val bool) {
	obj.activatable = val
}


func (obj *Object) GetSlot(key string) (v IObject, ctx IObject) {
	var ok bool

	obj.scanned = true

	if v, ok = obj.slots[key]; ok {
		ctx = obj
		return v, ctx
	}

	if v == nil {
		return nil, nil
	}

	if obj.proto != nil {
		v, ctx = obj.proto.GetSlot(key)
	} else {
		v, ctx = nil, nil
	}

	obj.scanned = false

	return v, ctx
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
	if obj.activatable {
		v, context := obj.GetSlot("activate")
		if v != nil && context != nil {
			return v.Activate(target, locals, msg, context)
		}
	}
	return obj
}
