package controller

import (
	"fmt"
	"iscool/vfs/controller/validate"
	"sort"
	"strings"
	"time"
)

// CreateFile creates a new file in the specified folder for the user
func (fs *FileSystem) CreateFile(username, foldername, filename, description string) error {
	user := fs.getUserByUsername(username)
	if user == nil {
		return fmt.Errorf("Error: The %s doesn't exist.", username)
	}

	folder := user.getFolderByName(foldername)
	if folder == nil {
		return fmt.Errorf("Error: The %s doesn't exist.", foldername)
	}

	if validate.ValidateNoInvalidChars(filename) {
		return fmt.Errorf("Error: The %s contains invalid chars.", filename)
	}

	if validate.ValidateLength(filename, 255) {
		return fmt.Errorf("Error: Filename be under %d characters.", 255)
	}

	if folder.isFileExists(filename) {
		return fmt.Errorf("Error: The %s has already existed.", filename)
	}

	now := time.Now()
	file := &File{
		Name:        filename,
		Description: description,
		CreatedAt:   now,
	}

	folder.Files[filename] = file

	return nil
}

// DeleteFile deletes the specified file from the folder for the user
func (fs *FileSystem) DeleteFile(username, foldername, filename string) error {
	user := fs.getUserByUsername(username)
	if user == nil {
		return fmt.Errorf("Error: The %s doesn't exist.", username)
	}

	folder := user.getFolderByName(foldername)
	if folder == nil {
		return fmt.Errorf("Error: The %s doesn't exist.", foldername)
	}

	if !folder.isFileExists(filename) {
		return fmt.Errorf("Error: The %s doesn't exist.", filename)
	}
	delete(folder.Files, filename)
	return nil
}

// ListFiles lists all the files in the specified folder for the user
func (fs *FileSystem) ListFiles(username, foldername, sortBy, sortOrder string) (string, error) {
	user := fs.getUserByUsername(username)
	if user == nil {
		return "", fmt.Errorf("Error: The %s doesn't exist.", username)
	}

	folder := user.getFolderByName(foldername)
	if folder == nil {
		return "", fmt.Errorf("Error: The %s doesn't exist.", foldername)
	}

	files := folder.Files
	if len(files) == 0 {
		return "", fmt.Errorf("Warning: The folder is empty")
	}

	switch sortBy + " " + sortOrder {
	case "--sort-name asc", "--sort-name desc", "--sort-created asc", "--sort-created desc", " ":
	default:
		// suggest a valid flag to the user
		return "", fmt.Errorf("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
	}

	// Create a slice to store the file information
	fileInfo := make([]*File, 0, len(files))

	// Append the file information to the slice
	for _, file := range files {
		fileInfo = append(fileInfo, file)
	}

	// Sort the fileInfo based on the selected sorting option
	switch sortBy {
	case "--sort-name":
		sort.Slice(fileInfo, func(i, j int) bool {
			if sortOrder == "asc" {
				return fileInfo[i].Name < fileInfo[j].Name
			}
			return fileInfo[i].Name > fileInfo[j].Name
		})
	case "--sort-created":
		sort.Slice(fileInfo, func(i, j int) bool {
			if sortOrder == "asc" {
				return fileInfo[i].CreatedAt.Before(fileInfo[j].CreatedAt)
			}
			return fileInfo[i].CreatedAt.After(fileInfo[j].CreatedAt)
		})
	default:
		// sort by name asc by default
		sort.Slice(fileInfo, func(i, j int) bool {
			return fileInfo[i].Name < fileInfo[j].Name
		})
	}

	var output []string
	// Print the file information in the specified format
	for _, file := range fileInfo {
		output = append(output, fmt.Sprintf("%s %s %s %s %s", file.Name, file.Description, file.CreatedAt.Format("2006-01-02 15:04:05"), foldername, username))
	}
	result := strings.Join(output, "\n")
	return result, nil
}

// isFileExists checks if the file exists in the folder
func (f *Folder) isFileExists(filename string) bool {
	_, exists := f.Files[filename]
	return exists
}
