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
 * File: state.go
 * Description: Startup stuff. Gets the VM going.
 ******************************************************************/

package europa

import "container/vector"

type State struct {
	lobby IObject
}

type IState interface {
	GetLobby() IObject
	InitializeState()
	EvaluateTree(vector.Vector)
}

func (state *State) GetLobby() IObject {
	return state.lobby
}

func (state *State) InitializeState() {
	state.lobby = NewObject(state, nil, false, false, false)
	object := NewObject(state, state.lobby, false, false, false)
	state.lobby.SetProto(object)
	state.lobby.SetSlot("Lobby", state.lobby)
	state.lobby.SetSlot("Object", object)
}

func (state *State) EvaluateTree(tree vector.Vector) {
	tree.Do(func(elem interface{}) {
		msg := elem.(IMessage)
		println(msg.GetName())
		msg.GetArguments().Do(func(item interface{}) {
			arg := item.(IMessage)
			println(arg.GetName())
		})
	})
}
