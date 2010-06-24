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
	"os"
	"fmt"
	"flag"
	"./europa"
)

const EUROPA_VERSION = "0.1.0"

func displayUsage() {
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(os.Stderr, "\t-%-10s%s\n", f.Name, f.Usage)
	})
}

func main() {
	var eflag string
	var versionflag bool
	var hflag bool
	flag.StringVar(&eflag, "e", "", "Specify code to evaluate at the command line.")
	flag.BoolVar(&versionflag, "version", false, "Displays the interpreter version information.")
	flag.BoolVar(&versionflag, "v", false, "Displays the interpreter version information.")
	flag.BoolVar(&hflag, "h", false, "Displays the usage information to the screen.")
	flag.Parse()

	if hflag {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
		displayUsage()
		return
	}

	if versionflag {
		println(string(EUROPA_VERSION))
		return
	}

	println("Setting up VM State...")
	state := new(europa.State)
	state.InitializeState()

	if eflag == "" {
		europa.Parse("test.io")
	} else {
		europa.ParseString(eflag)
	}
}
