package controller

func NewFileSystem() *FileSystem {
	return &FileSystem{
		Users: make(map[string]*User),
	}
}
