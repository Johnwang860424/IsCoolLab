package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type User struct {
	Name string
}

type UserManager struct {
	Users map[string]*User
}

// * represents a pointer to a type
// & represents the address of a variable
func NewUserManager() *UserManager {
	return &UserManager{
		Users: make(map[string]*User),
	}
}

func (um *UserManager) Register(name string) error {
	regex := `[\\/:*?"<>|\s]`
	if match, _ := regexp.MatchString(regex, name); match {
		return fmt.Errorf("Error: Username cannot contain invalid chars.")
	}

	_, exists := um.Users[strings.ToLower(name)]
	if exists {
		return fmt.Errorf("Error: The %s has already existed.", name)
	}

	um.Users[strings.ToLower(name)] = &User{
		Name: name,
	}

	fmt.Printf("Add %s successfully.\n", name)
	return nil
}

func main() {
	um := NewUserManager()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 1 {
			continue
		}
		switch cmd := parts[0]; cmd {
		case "register":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "Error: Missing argument.")
				continue
			}
			args := strings.Join(parts[1:], " ")
			if err := um.Register(strings.TrimSpace(args)); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "Error: Unknown command.")
		}
	}
}
