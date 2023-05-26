package controller

import "time"

type FileSystem struct {
	Users map[string]*User
}

type User struct {
	Name    string
	Folders map[string]*Folder
}

type Folder struct {
	Name        string
	Description string
	CreatedAt   time.Time
	Files       map[string]*File
}

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}
