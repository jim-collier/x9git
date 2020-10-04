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
	"fmt"
	"strings"

	"github.com/alexflint/go-arg"
	//"github.com/alexflint/go-arg"
)

// Version is for x9go_build to set
var Version string

// GitCommitHash is for x9go_build to set
var GitCommitHash string

// BuildDateTime is for x9go_build to set
var BuildDateTime string

// GoVersion is for x9go_build to set
var GoVersion string

func main() {

	type cmdP2f struct {
		help bool `arg:"-h" help:"Show help"`
	}

	var appArgs struct {
		p2f *cmdP2f `arg:"subcommand:p2f" placeholder:"COMMAND" help:"Command to execute."`
	}

	arg.MustParse(&appArgs) // arg
	fmt.Println(appArgs.p2f)

}

func getVersion() string {
	// Show variables injected at compile-time
	slice := []string{
		"Version .................: " + Version,
		"Built with go Version ...: " + GoVersion,
		"Build date/time .........: " + BuildDateTime,
		"Git commit hash .........: " + GitCommitHash}

	//fmt.Println("Version .................: ", Version)
	//fmt.Println("Built with go Version ...: ", GoVersion)
	//fmt.Println("Git commit hash .........: ", GitCommitHash)
	//fmt.Println("Build date/time .........: ", BuildDateTime)

	return (strings.Join(slice, "\n"))
}
