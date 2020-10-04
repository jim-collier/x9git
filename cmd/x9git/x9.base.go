package main

import (
	"os"
	"path/filepath"
	"strings"
)

//"github.com/alexflint/go-arg"

// Version is for x9go_build to set
var Version string

// GitCommitHash is for x9go_build to set
var GitCommitHash string

// BuildDateTime is for x9go_build to set
var BuildDateTime string

// GoVersion is for x9go_build to set
var GoVersion string

// ShowVersion outputs the above version-related strings.
//	History:
//		- 20201003 JC: Created.
func ShowVersion() string {
	// Show variables injected at compile-time
	slice := []string{
		//filepath.Base(os.Args[0]),
		meName,
		"Version .................: " + Version,
		"Built with go Version ...: " + GoVersion,
		"Build date/time .........: " + BuildDateTime,
		"Git commit hash .........: " + GitCommitHash,
	}

	//fmt.Println("Version .................: ", Version)
	//fmt.Println("Built with go Version ...: ", GoVersion)
	//fmt.Println("Git commit hash .........: ", GitCommitHash)
	//fmt.Println("Build date/time .........: ", BuildDateTime)

	return (strings.Join(slice, "\n"))
}

var meCommand string
var meName string
var args []string
var argsStr string

// Init sets variables
func Init() {

	// Get the basics
	meCommand = os.Args[0]
	meName = filepath.Base(meCommand)
	args = os.Args[1:]

	// Loop through slice and build single-string argStr.
	// If any string has spaces in it, surround it with quotes.
	argsStr = ""
	for _, arg := range args {
		if len(arg) > 0 {
			if len(argsStr) > 0 {
				argsStr = argsStr + ` `
			}
			if strings.ContainsAny(arg, " ") {
				argsStr = argsStr + `"` + arg + `"`
			} else {
				argsStr = argsStr + arg
			}
		}
	}

}
