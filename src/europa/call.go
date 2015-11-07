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
 * File: call.go
 * Description: Contains information about method calls
 ******************************************************************/

package europa

type Call struct {
	*Object
	sender    IObject
	message   IMessage
	target    IObject
	context   IObject
	activated IBlock
}

type ICall interface {
	IObject
}

func NewCall(sender IObject, target IObject, msg IMessage, context IObject, activated IBlock) ICall {
	return &Call{
		sender:    sender,
		target:    target,
		message:   msg,
		context:   context,
		activated: activated,
	}
}
