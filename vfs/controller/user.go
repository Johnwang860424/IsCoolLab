package controller

import (
	"fmt"
	"iscool/vfs/controller/validate"
	"strings"
)

// Register register a new user
func (fs *FileSystem) Register(name string) error {
	if validate.ValidateNoInvalidChars(name) {
		return fmt.Errorf("Error: The %s contain invalid chars.", name)
	}

	if validate.ValidateLength(name, 50) {
		return fmt.Errorf("Error: Username is must be under %d characters.", 50)
	}

	if fs.isUserExists(name) {
		return fmt.Errorf("Error: The %s has already existed.", name)
	}

	fs.Users[strings.ToLower(name)] = &User{
		Name:    name,
		Folders: make(map[string]*Folder),
	}
	return nil
}

// check if the user exists
func (fs *FileSystem) isUserExists(username string) bool {
	_, ok := fs.Users[strings.ToLower(username)]
	return ok
}
