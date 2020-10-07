package main

import (
	"os"
)

// FsObjType exists because a type and const, declared together, emulate an 'enum'
type FsObjType int

// Enums of type FsObjType
// Dir: regular dir. File: REGULAR file (not symlink). Link: symlink, junction, or mountpoint. Socket, NamedPipe: special *nix
const (
	Dir FsObjType = iota
	File
	Link
	Socket
	NamedPipe
)

// FsObj wraps and abstracts the complicated messyness of cross-platform filesystem objects (which no programming language yet invented by man, seems to have really cracked in a directly useful way)
type FsObj struct {
	path                    string
	fsObjType               FsObjType
	linkProps               *FsObjLink // Pointer so that we can leave it null most of the time. This is guaranteed to be non-null, if fsObjType = Link.
	inode                   int64
	canWrite                bool
	isVisible               bool
	bytes                   int64
	crtimeBirthedUTC        int64
	mtimeContentUpdatedUTC  int64
	ctimeMetadataUpdatedUTC int64
	xattrStr                string
	propsUpdatedUTC         int64 // When was this struct updated. Can use to base decisions on if needs to be refreshed.
}

// FsObjLink contains various properties describing
type FsObjLink struct {
	IsGood      bool
	linkedFsObj *FsObj // Pointer so that we can leave it null most of the time
}

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
