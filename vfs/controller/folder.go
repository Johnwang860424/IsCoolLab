package controller

import (
	"fmt"
	"iscool/vfs/controller/validate"
	"sort"
	"strings"
	"time"
)

// CreateFolder creates a new folder for the user
func (fs *FileSystem) CreateFolder(username string, foldername string, description string) error {
	user := fs.getUserByUsername(username)
	if user == nil {
		return fmt.Errorf("Error: %s does not exist.", username)
	}

	if validate.ValidateNoInvalidChars(foldername) {
		return fmt.Errorf("Error: The %s contain invalid chars.", foldername)
	}

	if validate.ValidateLength(foldername, 100) {
		return fmt.Errorf("Error: Foldername be under %d characters.", 100)
	}

	if user.isFolderExists(foldername) {
		return fmt.Errorf("Error: The %s has already existed.", foldername)
	}

	folder := &Folder{
		Name:        foldername,
		Description: description,
		CreatedAt:   time.Now(),
		Files:       make(map[string]*File),
	}

	user.Folders[foldername] = folder

	return nil
}

// ListFolders lists all the folders for the user
func (fs *FileSystem) ListFolders(username string, sortBy string, sortOrder string) (string, error) {
	user := fs.getUserByUsername(username)
	if user == nil {
		return "", fmt.Errorf("Error: %s doesn't exist.", username)
	}

	switch sortBy + " " + sortOrder {
	case "--sort-name asc", "--sort-name desc", "--sort-created asc", "--sort-created desc", " ":
	default:
		// suggest a valid flag to the user
		return "", fmt.Errorf("Error: Unknown flag. Valid flags are '--sort-name asc' '--sort-name asc' '--sort-created asc' '--sort-created desc'")
	}

	folders := user.Folders

	if len(folders) == 0 {
		return "", fmt.Errorf("Warning: The %s doesn't have any folders.", username)
	}

	// Create a slice to store the folder information
	folderInfo := make([]*Folder, 0, len(folders))

	// Append the folder information to the slice
	for _, folder := range folders {
		folderInfo = append(folderInfo, folder)
	}

	// Sort the folderInfo based on the selected sorting option
	switch sortBy {
	case "--sort-name":
		sort.Slice(folderInfo, func(i, j int) bool {
			if sortOrder == "asc" {
				return folderInfo[i].Name < folderInfo[j].Name
			}
			return folderInfo[i].Name > folderInfo[j].Name
		})
	case "--sort-created":
		sort.Slice(folderInfo, func(i, j int) bool {
			if sortOrder == "asc" {
				return folderInfo[i].CreatedAt.Before(folderInfo[j].CreatedAt)
			}
			return folderInfo[i].CreatedAt.After(folderInfo[j].CreatedAt)
		})
	default:
		// Sort by folder name in ascending order by default
		sort.Slice(folderInfo, func(i, j int) bool {
			return folderInfo[i].Name < folderInfo[j].Name
		})
	}

	var output []string
	// Print the folder information in the specified format
	for _, folder := range folderInfo {
		output = append(output, fmt.Sprintf("%s %s %s %s", folder.Name, folder.Description, folder.CreatedAt.Format("2006-01-02 15:04:05"), username))
	}
	result := strings.Join(output, "\n")
	return result, nil
}

// DeleteFolder deletes the specified folder for the user
func (fs *FileSystem) DeleteFolder(username string, foldername string) error {
	user := fs.getUserByUsername(username)
	if user == nil {
		return fmt.Errorf("Error: %s doesn't exist.", username)
	}
	folder := user.Folders[foldername]
	if folder == nil {
		return fmt.Errorf("Error: %s doesn't exist.", foldername)
	}
	delete(user.Folders, foldername)
	return nil
}

// RenameFolder renames the specified folder for the user
func (fs *FileSystem) RenameFolder(username string, foldername string, newFolderName string) error {
	user := fs.getUserByUsername(username)
	if user == nil {
		return fmt.Errorf("Error: %s doesn't exist.", username)
	}
	folder := user.getFolderByName(foldername)
	if folder == nil {
		return fmt.Errorf("Error: %s doesn't exist.", foldername)
	}
	if folder.Name == newFolderName {
		return nil // No need to rename if the new folder name is the same
	}
	if validate.ValidateNoInvalidChars(newFolderName) {
		return fmt.Errorf("Error: The %s contain invalid chars.", newFolderName)
	}
	if user.isFolderExists(newFolderName) {
		return fmt.Errorf("Error: The %s has already existed.", newFolderName)
	}
	user.Folders[newFolderName] = folder
	folder.Name = newFolderName
	delete(user.Folders, foldername)
	return nil
}

// getUserByUsername returns the specified user
func (fs *FileSystem) getUserByUsername(name string) *User {
	user, ok := fs.Users[strings.ToLower(name)]
	if ok {
		return user
	}
	return nil
}

// getFolderByName returns the specified folder for the user
func (u *User) getFolderByName(foldername string) *Folder {
	folder, ok := u.Folders[foldername]
	if ok {
		return folder
	}
	return nil
}

// isFolderExists checks if the specified folder exists for the user
func (u *User) isFolderExists(foldername string) bool {
	_, exists := u.Folders[foldername]
	return exists
}
