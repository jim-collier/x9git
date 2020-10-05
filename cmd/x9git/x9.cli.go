package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//
// Data
//

// CLI is some light wrappers around Printf, Println, Sprintf
//	History:
//		- 20201003 JC: Created.
type CLI struct {
	lastEchoWasEmpty bool
}

//
// Behavior
//

// Echo wraps EchoClean with "[ arg ]"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) Echo(args ...interface{}) {
	outputStr := fmt.Sprint(args...)
	if len(outputStr) == 0 {
		cli.EchoClean()
	} else {
		cli.EchoClean("[ ", outputStr, " ]")
	}
}

// EchoClean is the basis for the others"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) EchoClean(args ...interface{}) {
	outputStr := fmt.Sprint(args...)
	if len(outputStr) == 0 {
		if !cli.lastEchoWasEmpty {
			fmt.Println()
		}
		cli.lastEchoWasEmpty = true
	} else {
		fmt.Println(outputStr)
		cli.lastEchoWasEmpty = false
	}
}

// Echof wraps EchofClean with "[ arg ]"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) Echof(argFormat string, vals ...interface{}) {
	outputStr := fmt.Sprintf(argFormat, vals...)
	cli.Echo(outputStr)
}

// EchoCleanf is the basis for the others"
//	History:
//		- 20201003 JC: Created.
func (cli CLI) EchoCleanf(argFormat string, vals ...interface{}) {
	outputStr := fmt.Sprintf(argFormat, vals...)
	cli.EchoClean(outputStr)
}

// ResetBlankCounter so that next echo can be blank no matter what
//	History:
//		- 20201003 JC: Created.
func (cli CLI) ResetBlankCounter() {
	cli.lastEchoWasEmpty = false
}

//	Shell execution
//	This is an excellent reference for shelling out to shell for sure shelling shored shelly shell, shell:
//		https://medium.com/rungo/executing-shell-commands-script-files-and-executables-in-go-894814f1c0f7

// Exec runs a function and outputs to stdout and stderr on the CLI.
//		The idea is to closely mimic just running a command in Bash or CMD.
// Returns:
//		command's exit code
//		Go-specific error (not command error - that's in stdout).
// History:
//		- 20201004 JC: Created.
// Notes:
//		- Probably want to use os.exec.Command().Run()
//		- Fun fact, Go's .Run() command is a simple wrapper around .Start() and .Wait().
func (cli CLI) Exec(command string, args []string) (returnCode int, err error) { //TODO: , stdErr string) {

	returnCode = -1 // Default
	err = *new(error)
	var commandExists bool

	// Validate existence of `command`. exec.Command() does almost exactly this, but we can't access the error it generates [private to exec.go], unless we do it ourself.
	if filepath.Base(command) == command {
		// Look in path
		_, err = exec.LookPath(command)
		commandExists = (err == nil)
	} else {
		// Validate existence directly
		commandExists, err = IsFileVisible(command)
	}

	if commandExists {

		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				returnCode = exitError.ExitCode()
				// Debug
				fmt.Printf("Exit Status: %d", returnCode)
			}
		}
	}

	return
}

// GetOutput runs a shell command but outputs nothing, only capturing and returning both stdout and stderr.
//	History:
//		- 20201004 JC: Created.
//	Notes:
//		- Probably want to use os.exec.Command().Output()
func (cli CLI) GetOutput(command string, args []string) (stdOut string, stdErr string, returnCode int) {
	return
}
