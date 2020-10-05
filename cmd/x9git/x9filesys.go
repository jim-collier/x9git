package main

import (
	"os"
)

// IsPathspecVisible returns true if pathspec exists, AND is accessible.
// References:
//		https://stackoverflow.com/a/44815664/9190681
// History:
//		-20201004 JC: Created.
func IsPathspecVisible(pathspec string) (doesExist bool, err error) {
	_, err = os.Stat(pathspec)
	if err == nil {
		doesExist = true // Definitively exists
	} else if os.IsNotExist(err) {
		err = nil
		doesExist = false // Definitively doesn't exist
	} else {
		doesExist = false // Can't tell, err has more info
	}
	return
}

// IsDirVisible returns true IF pathspec is a dir, AND exists, AND is accessible.
//	History:
//		-20201004 JC: Created.
func IsDirVisible(pathspec string) (isDir bool, err error) {
	isDir = false // Default
	var exists bool
	exists, err = IsPathspecVisible(pathspec)
	if exists {
		var fileInfo os.FileInfo
		fileInfo, err = os.Stat(pathspec) // Repeats code in IsPathspecVisible()
		if err != nil {
			isDir = fileInfo.IsDir()
		}
	}
	return
}

// IsFileVisible returns true IF pathspec is a file, AND exists, AND is accessible.
//	History:
//		-20201004 JC: Created.
func IsFileVisible(pathspec string) (isFile bool, err error) {
	isFile = false // Default
	var exists bool
	if exists, err = IsPathspecVisible(pathspec); exists {
		var fileInfo os.FileInfo
		fileInfo, err = os.Stat(pathspec) // Repeats code in IsPathspecVisible()
		if err != nil {
			isFile = fileInfo.Mode().IsRegular()
		}
	}
	return
}
