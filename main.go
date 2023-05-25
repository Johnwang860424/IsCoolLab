package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

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

// NewFileSystem creates a new file system
func NewFileSystem() *FileSystem {
	return &FileSystem{
		Users: make(map[string]*User),
	}
}

func (fs *FileSystem) Register(name string) error {

	_, exists := fs.Users[strings.ToLower(name)]
	if exists {
		return fmt.Errorf("Error: The %s has already existed", name)
	}

	fs.Users[strings.ToLower(name)] = &User{
		Name:    name,
		Folders: make(map[string]*Folder),
	}

	fmt.Printf("Add %s successfully.\n", name)
	return nil
}

func (fs *FileSystem) GetUser(name string) (*User, error) {
	user, ok := fs.Users[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("Error: %s does not exist.", name)
	}

	return user, nil
}

func (u *User) CreateFolder(name string, description string) error {
	if validLength := ValidateLength(name, 100); validLength != nil {
		return validLength
	}
	regex := `[\\/:*?"<>|\s]`
	if validChars := ValidateNoInvalidChars(name, regex); validChars != nil {
		return validChars
	}
	if _, exists := u.Folders[name]; exists {
		return fmt.Errorf("Error: The %s has already existed.", name)
	}

	folder := &Folder{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Files:       make(map[string]*File),
	}

	u.Folders[name] = folder
	fmt.Printf("Create %s successfully.\n", name)
	return nil
}

func (u *User) DeleteFolder(name string) error {
	if _, exists := u.Folders[name]; !exists {
		return fmt.Errorf("Error: The %s doesn't exist.", name)
	}
	delete(u.Folders, name)
	fmt.Printf("Delete %s successfully.\n", name)
	return nil
}

func (u *User) ListFolders(sortBy string, sortOrder string) {
	if len(u.Folders) == 0 {
		fmt.Fprintf(os.Stderr, "Warning: The %s doesn't have any folders.\n", u.Name)
	}

	// Create a slice to store the folder information
	var folders []Folder

	// Append the folder information to the slice
	for _, folder := range u.Folders {
		folders = append(folders, *folder)
	}

	// Sort the folders based on the selected sorting option
	switch sortBy {
	case "--sort-name":
		sort.Slice(folders, func(i, j int) bool {
			if sortOrder == "asc" {
				return folders[i].Name < folders[j].Name
			}
			return folders[i].Name > folders[j].Name
		})
	case "--sort-created":
		sort.Slice(folders, func(i, j int) bool {
			if sortOrder == "asc" {
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			}
			return folders[i].CreatedAt.After(folders[j].CreatedAt)
		})
	default:
		// Sort by folder name in ascending order by default
		sort.Slice(folders, func(i, j int) bool {
			return folders[i].Name < folders[j].Name
		})
	}

	// Print the folder information in the specified format
	for _, folder := range folders {
		fmt.Printf("%s %s %s %s\n", folder.Name, folder.Description, folder.CreatedAt.Format("2006-01-02 15:04:05"), u.Name)
	}
}

func (u *User) RenameFolder(oldName string, newName string) error {
	if _, exists := u.Folders[oldName]; !exists {
		return fmt.Errorf("Error: The %s doesn't exist.", oldName)
	}
	if _, exists := u.Folders[newName]; exists {
		return fmt.Errorf("Error: The %s has already existed.", newName)
	}
	u.Folders[newName] = u.Folders[oldName]
	delete(u.Folders, oldName)
	fmt.Printf("Rename %s to %s successfully.\n", oldName, newName)
	return nil
}

func ValidateNoInvalidChars(name string, regex string) error {
	if match, _ := regexp.MatchString(regex, name); match {
		return fmt.Errorf("Error: The %s contain invalid chars.", name)
	}
	return nil
}

func ValidateLength(name string, length int) error {
	if len(name) > length {
		return fmt.Errorf("Error: Must be under %d characters.", length)
	}
	return nil
}

func main() {
	fs := NewFileSystem()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		command := args[0]
		commandArgs := args[1:]
		switch command {
		case "register":
			if len(commandArgs) < 1 {
				fmt.Fprintln(os.Stderr, "Error: Missing argument.")
				continue
			}
			username := strings.TrimSpace(strings.Join(commandArgs[0:], " "))
			regex := `[\\/:*?"<>|\s]`
			if err := ValidateLength(username, 50); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if err := ValidateNoInvalidChars(username, regex); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if err := fs.Register(strings.TrimSpace(username)); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
		case "create-folder":
			if len(commandArgs) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Missing arguments.")
				continue
			}
			username := commandArgs[0]
			foldername := commandArgs[1]
			description := ""
			regex := `[\\/:*?"<>|\s]`
			user, err := fs.GetUser(username)
			if len(commandArgs) > 2 {
				description = strings.Join(commandArgs[2:], "")
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if err := ValidateLength(username, 100); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if err := ValidateNoInvalidChars(username, regex); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if err := user.CreateFolder(foldername, description); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

		case "delete-folder":
			if len(commandArgs) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Missing arguments.")
				continue
			}
			username := commandArgs[0]
			foldername := commandArgs[1]
			user, err := fs.GetUser(username)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if err := user.DeleteFolder(foldername); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "list-folders":
			if len(commandArgs) < 1 {
				fmt.Fprintln(os.Stderr, "Error: Missing arguments.")
				continue
			}
			username := commandArgs[0]
			user, err := fs.GetUser(username)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			var sortBy, sortOrder string
			// Parse optional flags
			if len(commandArgs) == 3 || len(commandArgs) == 2 {
				sortBy = commandArgs[1]
				switch sortBy {
				case "--sort-name", "--sort-created":
				default:
					// suggest a valid flag to the user
					fmt.Fprintf(os.Stderr, "Error: Unknown flag '%s'. Valid flags are '--sort-name' and '--sort-created'.\n", sortBy)
					continue
				}
				if len(commandArgs) > 2 {
					sortOrder = commandArgs[2]
					switch sortOrder {
					case "asc", "desc":
					default:
						fmt.Fprintf(os.Stderr, "Error: Unknown flag '%s'. Valid flags are 'asc' and 'desc'.\n", sortOrder)
						continue
					}
				} else {
					fmt.Fprintf(os.Stderr, "Error: Missing sort order flag. Valid flags are 'asc' and 'desc'.\n")
					continue
				}
			}
			user.ListFolders(sortBy, sortOrder)
		case "rename-folder":
			if len(commandArgs) < 3 {
				fmt.Fprintln(os.Stderr, "Error: Missing arguments.")
				continue
			}
			username := commandArgs[0]
			foldername := commandArgs[1]
			newFoldername := commandArgs[2]
			user, err := fs.GetUser(username)
			regex := `[\\/:*?"<>|\s]`
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if validLength := ValidateLength(newFoldername, 100); validLength != nil {
				fmt.Fprintln(os.Stderr, validLength)
				continue
			}
			if validChars := ValidateNoInvalidChars(newFoldername, regex); validChars != nil {
				fmt.Fprintln(os.Stderr, validChars)
				continue
			}
			if err := user.RenameFolder(foldername, newFoldername); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "Error: Unknown command.")
		}
	}
}
