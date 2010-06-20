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
 * File: block.go
 * Description: Implements closures
 ******************************************************************/

package europa

type Block struct {
	*Object
	cfunc (func(IObject, IObject, IMessage) IObject)
	message Message
	argNames []string
	scope IObject
}

type IBlock interface {
	IObject
}

func (blk *Block) Clone() IObject {
	r := new(Block)
	r.proto = blk
	r.message = blk.message
	r.argNames = blk.argNames
	r.scope = blk.scope
	r.activatable = blk.activatable
	return r
}

func (blk *Block) Activate(target IObject, locals IObject, m IMessage, ctx IObject) IObject {
	if blk.cfunc != nil {
		return blk.cfunc(target, locals, m)
	}

	scope := blk.scope
	blockLocals := ctx.Clone()

	blockLocals.SetLocals(true)
	if scope == nil {
		scope = target
	}
	callObject := NewCall(locals, target, m, ctx, blk)

	blockLocals.SetSlot("call", callObject)
	blockLocals.SetSlot("self", scope)
	/*blockLocals.SetSlot("updateSlot", someIBlock)*/

	for i, name := range blk.argNames {
		arg := m.ArgAt(locals, i)
		blockLocals.SetSlot(name, arg)
	}

	return blk.message.PerformOn(blockLocals, blockLocals)
}
