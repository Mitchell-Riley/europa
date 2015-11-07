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
 * File: number.go
 * Description: A simple numeric base type
 ******************************************************************/

package europa

type Number struct {
	*Object
	value float64
}

type INumber interface {
	IObject

	GetValue() float64
	SetValue(float64)
	Add(INumber, IObject, IMessage) INumber
	Sub(INumber, IObject, IMessage) INumber
}

func NewNumber(num float64) INumber {
	return &Number{
		Object: &Object{slots: make(map[string]IObject, DefaultSlotsSize)},
		value:  num,
	}
}

func (num *Number) Clone() IObject {
	r := NewNumber(num.value)
	r.SetProto(num)
	return r
}

func (num *Number) GetValue() float64 {
	return num.value
}

func (num *Number) SetValue(val float64) {
	num.value = val
}

func (num *Number) Add(self INumber, locals IObject, msg IMessage) INumber {
	other := msg.NumberArgAt(locals, 0)
	return NewNumber(self.GetValue() + other.GetValue())
}

func (num *Number) Sub(self INumber, locals IObject, msg IMessage) INumber {
	other := msg.NumberArgAt(locals, 0)
	return NewNumber(self.GetValue() - other.GetValue())
}
