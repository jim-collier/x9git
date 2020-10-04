//	Purpose: An opinionated wrapper around Git.
//	Notes:
//		- Project go-git (https://github.com/go-git/go-git), a native go implimentation of git,
//		  isn't ready and probably never will be. (Git is very complex!) Therefore, this uses
//		  shell commands.
//		- Basically, this could just as easily have been written in Bash. But it's in Go, for
//		  a little cleaner structure, type checking, etc.
//		- Initially written and tested only on Linux. May be expanded to include Windows later.
//	Syntax design:
//		x9git new repo <project> <URI root> [--withdevbranch]   new repo, rename master to main, [create develop]
//		x9git clone|get <url>
//		x9git new-feature <name>
//		x9git del-branch  [--local|--remote]
//		x9git ch|changeto <[branch|repo [branch]>   cd, [branch], pull
//		x9git quicksave                             stash
//		x9git freshen[-from] [main|develop]         pull from main, then branch
//		x9git commit                                pull, commit
//		x9git p2f|p2feat|push2feat|push-to-feature
//		x9git m2t|m2top|merge-to-top
//		x9git rename <old branch> <new branch>      https://stackoverflow.com/a/29650705/9190681
//	Config file: .config/x9git.yaml
//		v1
//			gitroots:
//				- ~/git
//			overrides:
//				## Normally x9git figures this out through "convention over configuration", or pseudo-reflection and assumptions
//				projects:
//					x9sync:
//						dir: ~/anotherdir/x9sync-cs
//						ssh: git@github.com:x9-testlab/x9sync-csharp.git
//						can commit to non-feature branches: yes
//						highest directly committable branch: develop

package main

import (
	"regexp"
	"strings"
	//"github.com/alexflint/go-arg"
)

const copyrightYears = "2020"
const copyrightAuthors = "Jim Collier"

const regexFlagShort = `(?:-|\/)`
const regexFlagLong = `(?:-{1,2}|\/)`
const regexFlagLongOrNoFlag = `(?:-{1,2}|\/)?`

var cli CLI

//meCommand := os.Args[0]
//meName := filepath.Base(os.Args[0])

type tParsedArgs struct {
	showOnly    bool
	functionPtr func() (err error)
}

var parsedArgs tParsedArgs

func cmdShowAbout() (err error) {
	cli.EchoClean()
	cli.EchoClean("A safe and goal-oriented, workflow-opinionated wrapper around the Git CLI.")
	cli.EchoClean("Written and tested with 'git' version 2.28.0.")
	cli.EchoClean()
	return nil
}

func cmdShowHelp() (err error) {
	cli.EchoClean()
	cmdShowAbout()
	cli.EchoClean()
	cli.EchoCleanf("Syntax: %s <command>", meName)
	cli.EchoClean()
	return nil
}

func cmdShowVersion() (err error) {
	ShowVersion()
	return nil
}

func cmdShowLicense() (err error) {
	cli.EchoCleanf("Copyright (c) %s, %s.", copyrightYears, copyrightAuthors)
	cli.EchoClean("License GPLv3+: GNU GPL version 3 or later, full text at:")
	cli.EchoClean("    https://www.gnu.org/licenses/gpl-3.0.en.html")
	cli.EchoClean("There is no warranty, to the extent permitted by law.")
	return nil
}

func main() {
	Init()
	cli.Echo()
	cli.Echo()

	// Parse Args
	doParseArgs()

	// Execute function
	if parsedArgs.functionPtr != nil {
		parsedArgs.functionPtr()
	}

	cli.Echo()
	cli.Echo("Done.")
	cli.Echo()
}

func doParseArgs() {

	// Look at args as a whole
	lcargsStr := strings.ToLower(argsStr) + " "
	switch true {

	case getFirstVal(regexp.MatchString(`^((?:-{1,2}|\/)?about) .*`, lcargsStr)):
		parsedArgs.functionPtr = cmdShowVersion

	case getFirstVal(regexp.MatchString(`^(?:(?:-|\/)h|(?:-{1,2}|\/)?help) .*`, lcargsStr)):
		parsedArgs.functionPtr = cmdShowHelp

	case getFirstVal(regexp.MatchString(`^(?:(?:-|\/)v|(?:-{1,2}|\/)?ver(?:sion)?) .*`, lcargsStr)):
		parsedArgs.functionPtr = cmdShowVersion

	case getFirstVal(regexp.MatchString(`^((?:-{1,2}|\/)?license) .*`, lcargsStr)):
		parsedArgs.functionPtr = cmdShowVersion

	}

	// Parse arguments
	for _, arg := range args {
		lcarg := strings.ToLower(arg)
		//cli.Echof("lcarg = '%s'", lcarg)
		switch true {

		case getFirstVal(regexp.MatchString(`^a.*b$`, lcarg)):
			cli.Echo("yes")

		}
	}
}

//func main() {
//
//	type cmdP2f struct {
//		help bool `arg:"-h" help:"Show help"`
//	}
//
//	var appArgs struct {
//		p2f *cmdP2f `arg:"subcommand:p2f" placeholder:"COMMAND" help:"Command to execute."`
//
//	}
//	func (appArgs) Version() string {
//
//	arg.MustParse(&appArgs) // arg belongs to alexflint/go-arg namespace, which presumably parses os.Args[1:]
//
//	switch {
//	case args.p2f != nil:
//		showHelp()
//	}
//
//}
