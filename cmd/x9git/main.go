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
)

//import "github.com/alexflint/go-arg"

var cli CLI

//meCommand := os.Args[0]
//meName := filepath.Base(os.Args[0])

func showHelp() string {
	return "Help"
}

func showLicense() string {
	return "GPL v3"
}

func showAbout() string {
	return ""
}

type tParsedArgs struct {
	command     string
	functionPtr func() (err error)
}

var parsedArgs tParsedArgs

func main() {
	Init()
	cli.Echo()

	// Parse Args
	doParseArgs()
	if parsedArgs.functionPtr != nil {
		cli.Echo("Invoking function.")
		if err := parsedArgs.functionPtr(); err == nil {
			cli.Echo("Ran.")
		}
	}

	cli.Echo()
	cli.Echo("Done.")
	cli.Echo()
}

func cmdShowVersion() (err error) {
	cli.Echo("Running ...")
	ShowVersion()
	return nil
}

func doParseArgs() {

	// Look at args as a whole
	lcargsStr := strings.ToLower(argsStr) + " "
	switch true {

	case getFirstVal(regexp.MatchString(`^(-v|-{0,2}version) .*`, lcargsStr)):
		parsedArgs.command = "version"
		parsedArgs.functionPtr = cmdShowVersion

	}

	cli.Echof("parsedArgs.command = '%s'", parsedArgs.command)

	// Parse arguments
	for _, arg := range args {
		lcarg := strings.ToLower(arg)
		cli.Echof("lcarg = '%s'", lcarg)
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
