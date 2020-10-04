package main

import "fmt"

// CLI is some light wrappers around Printf, Println, Sprintf
//	History:
//		- 20201003 JC: Created.
type CLI struct {
	lastEchoWasEmpty bool
}

// Echo wraps EchoClean with "[ arg ]"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) Echo(args ...interface{}) {
	outputStr := fmt.Sprint(args...)
	if len(outputStr) > 0 {
		cli.EchoClean("[ ", outputStr, " ]")
	} else {
		cli.EchoClean()
	}
}

// Echof wraps EchofClean with "[ arg ]"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) Echof(argFormat string, vals ...interface{}) {
	outputStr := fmt.Sprintf(argFormat, vals...)
	if len(outputStr) > 0 {
		cli.EchoClean("[ ", outputStr, " ]")
	} else {
		cli.EchoClean()
	}
}

// EchoClean is the basis for the others"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) EchoClean(args ...interface{}) {
	outputStr := fmt.Sprint(args...)
	if len(outputStr) > 0 {
		fmt.Println(outputStr)
		cli.lastEchoWasEmpty = false
	} else {
		if !cli.lastEchoWasEmpty {
			fmt.Println()
		}
		cli.lastEchoWasEmpty = true
	}
}

// EchoCleanf is the basis for the others"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) EchoCleanf(argFormat string, vals ...interface{}) {
	outputStr := fmt.Sprintf(argFormat, vals...)
	if len(outputStr) > 0 {
		fmt.Println(outputStr)
		cli.lastEchoWasEmpty = false
	} else {
		if !cli.lastEchoWasEmpty {
			fmt.Println()
		}
		cli.lastEchoWasEmpty = true
	}
}

// ResetBlankCounter so that next echo can be blank no matter what
//	History:
//		- 20201003 JC: Created.
func (cli CLI) ResetBlankCounter() {
	cli.lastEchoWasEmpty = false
}
