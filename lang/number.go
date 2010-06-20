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
	Add(INumber) INumber
	Sub(INumber) INumber
}

func NewNumber(num float64) INumber {
	r := new(Number)
	r.value = num
	return r
}

func (num *Number) GetValue() float64 {
	return num.value
}

func (num *Number) SetValue(val float64) {
	num.value = val
}

func (num *Number) Add(other INumber) INumber {
	return NewNumber(num.value + other.GetValue())
}

func (num *Number) Sub(other INumber) INumber {
	return NewNumber(num.value - other.GetValue())
}
