package controller

import (
	"testing"
	"time"
)

func TestCreateFile(t *testing.T) {
	fs := NewFileSystem()

	// Create a user and a folder
	err := fs.Register("test_user")
	if err != nil {
		t.Fatalf("Failed to register user: %s", err)
	}
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err != nil {
		t.Fatalf("Failed to create folder: %s", err)
	}

	// Create a file in the folder
	err = fs.CreateFile("test_user", "test_folder", "test_file.txt", "This is a test file.")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}

	// Try to create the same file again
	err = fs.CreateFile("test_user", "test_folder", "test_file.txt", "This is a test file.")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to create a file with an invalid character in the name
	err = fs.CreateFile("test_user", "test_folder", "test_file/.txt", "This is a test file.")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to create a file with a name that is too long
	longName := "test_file_with_a_very_long_name_that_is_over_fifty_characterstest_file_with_a_very_long_name_that_is_over_fifty_characterstest_file_with_a_very_long_name_that_is_over_fifty_characterstest_file_with_a_very_long_name_that_is_over_fifty_characterstest_file_with_a_very_long_name_that_is_over_fifty_characterstest_file_with_a_very_long_name_that_is_over_fifty_characterstest_file_with_a_very_long_name_that_is_over_fifty_characters.txt"
	err = fs.CreateFile("test_user", "test_folder", longName, "This is a test file.")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to create a file in a non-existent folder
	err = fs.CreateFile("test_user", "non_existent_folder", "test_file.txt", "This is a test file.")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to create a file for a non-existent user
	err = fs.CreateFile("non_existent_user", "test_folder", "test_file.txt", "This is a test file.")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestDeleteFile(t *testing.T) {
	fs := NewFileSystem()

	// Create a user, a folder, and a file
	err := fs.Register("test_user")
	if err != nil {
		t.Fatalf("Failed to register user: %s", err)
	}
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err != nil {
		t.Fatalf("Failed to create folder: %s", err)
	}
	err = fs.CreateFile("test_user", "test_folder", "test_file.txt", "This is a test file.")
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}

	// Delete the file
	err = fs.DeleteFile("test_user", "test_folder", "test_file.txt")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}

	// Try to delete the file again
	err = fs.DeleteFile("test_user", "test_folder", "test_file.txt")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to delete a non-existent file
	err = fs.DeleteFile("test_user", "test_folder", "non_existent_file.txt")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to delete a file from a non-existent folder
	err = fs.DeleteFile("test_user", "non_existent_folder", "test_file.txt")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to delete a file for a non-existent user
	err = fs.DeleteFile("non_existent_user", "test_folder", "test_file.txt")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestListFiles(t *testing.T) {
	fs := NewFileSystem()

	// Create a user, a folder, and some files
	err := fs.Register("test_user")
	if err != nil {
		t.Fatalf("Failed to register user: %s", err)
	}
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err != nil {
		t.Fatalf("Failed to create folder: %s", err)
	}
	err = fs.CreateFile("test_user", "test_folder", "file1.txt", "This is file 1.")
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	err = fs.CreateFile("test_user", "test_folder", "file2.txt", "This is file 2.")
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	err = fs.CreateFile("test_user", "test_folder", "file3.txt", "This is file 3.")
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}

	// List the files in the folder
	result, err := fs.ListFiles("test_user", "test_folder", "--sort-created", "asc")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	expected := "file1.txt This is file 1. " + time.Now().Format("2006-01-02 15:04:05") + " test_folder test_user\n" +
		"file2.txt This is file 2. " + time.Now().Format("2006-01-02 15:04:05") + " test_folder test_user\n" +
		"file3.txt This is file 3. " + time.Now().Format("2006-01-02 15:04:05") + " test_folder test_user"
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}

	// Try to list the files in a non-existent folder
	_, err = fs.ListFiles("test_user", "non_existent_folder", "--sort-created", "asc")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to list the files for a non-existent user
	_, err = fs.ListFiles("non_existent_user", "test_folder", "--sort-created", "asc")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to list the files with an invalid sort order
	_, err = fs.ListFiles("test_user", "test_folder", "--sort-created", "invalid_sort_order")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Try to list the files with an invalid sort option
	_, err = fs.ListFiles("test_user", "test_folder", "--invalid-sort-option", "asc")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestFolder_isFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		f    *Folder
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.isFileExists(tt.args.filename); got != tt.want {
				t.Errorf("Folder.isFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
