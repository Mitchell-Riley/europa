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
 * File: string.go
 * Description: UTF-8 Strings
 ******************************************************************/

package europa

type String struct {
	*Object
	value string
}

type IString interface {
	IObject
	
	SetValue(string)
	GetValue() string
}

func NewString(str string) IString {
	r := new(String)
	r.slots = make(map[string]IObject, DEFAULT_SLOTS_SIZE)
	r.value = str
	return r
}

func (str *String) Clone() IObject {
	r := NewString(str.value)
	r.SetProto(str)
	return r
}

func (str *String) SetValue(val string) {
	str.value = val
}

func (str *String) GetValue() string {
	return str.value
}
