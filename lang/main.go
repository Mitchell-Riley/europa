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
 * File: 
 * Description: 
 ******************************************************************/

package main

import (
	"./europa"
)

func main() {
	println("Setting up VM State...")
	state := new(europa.State)
	state.InitializeState()
	europa.Parse("Object clone do(1 2)")
}
