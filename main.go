package main

import (
	"bufio"
	"fmt"
	"iscool/vfs/controller"
	"os"
	"strings"
)

func main() {
	fs := controller.NewFileSystem()

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
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := strings.TrimSpace(strings.Join(commandArgs[0:], " "))
			err := fs.Register(username)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			} else {
				fmt.Printf("Add %s successfully.\n", username)
			}

		case "create-folder":
			if len(commandArgs) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			foldername := commandArgs[1]
			description := ""
			if len(commandArgs) > 2 {
				description = strings.Join(commandArgs[2:], "")
			}

			err := fs.CreateFolder(username, foldername, description)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Printf("Create %s successfully.\n", foldername)
			}

		case "delete-folder":
			if len(commandArgs) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			foldername := commandArgs[1]
			err := fs.DeleteFolder(username, foldername)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Delete %s successfully.\n", foldername)
			}

		case "list-folders":
			if len(commandArgs) < 1 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			var sortBy, sortOrder string

			// Parse optional flags
			switch len(commandArgs) {
			case 2:
				sortBy = commandArgs[1]
			case 3:
				sortBy = commandArgs[1]
				sortOrder = commandArgs[2]
			}

			output, err := fs.ListFolders(username, sortBy, sortOrder)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}

		case "rename-folder":
			if len(commandArgs) < 3 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			foldername := commandArgs[1]
			newFolderName := commandArgs[2]
			err := fs.RenameFolder(username, foldername, newFolderName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Rename %s to %s successfully.\n", foldername, newFolderName)
			}

		case "create-file":
			if len(commandArgs) < 3 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			foldername := commandArgs[1]
			filename := commandArgs[2]
			description := ""
			if len(commandArgs) > 3 {
				description = strings.Join(commandArgs[3:], "")
			}
			err := fs.CreateFile(username, foldername, filename, description)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Create %s in %s/%s successfully.\n", filename, username, foldername)
			}

		case "delete-file":
			if len(commandArgs) < 3 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			foldername := commandArgs[1]
			filename := commandArgs[2]
			err := fs.DeleteFile(username, foldername, filename)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Delete %s in %s/%s successfully.\n", filename, username, foldername)
			}

		case "list-files":
			if len(commandArgs) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
				continue
			}

			username := commandArgs[0]
			foldername := commandArgs[1]
			var sortBy, sortOrder string

			switch len(commandArgs) {
			case 3:
				sortBy = commandArgs[2]
			case 4:
				sortBy = commandArgs[2]
				sortOrder = commandArgs[3]
			}

			output, err := fs.ListFiles(username, foldername, sortBy, sortOrder)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "Error: Unrecognized command.")
		}
	}
}
