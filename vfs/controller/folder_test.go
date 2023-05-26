package controller

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateFolder(t *testing.T) {
	fs := NewFileSystem()

	// Test creating a folder for a user that doesn't exist
	err := fs.CreateFolder("test_user", "test_folder", "test_description")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test creating a folder with an invalid name
	err = fs.Register("test_user")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	err = fs.CreateFolder("test_user", "test_folder/", "test_description")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test creating a folder with a name that is too long
	longFolderName := "test_folder_with_a_very_long_name_that_is_over_one_hundred_characters_long_and_therefore_is_too_long_to_be_a_valid_folder_name"
	err = fs.CreateFolder("test_user", longFolderName, "test_description")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test creating a folder with a valid name
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}

	// Test creating a folder with the same name as an existing folder
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestListFolders(t *testing.T) {
	fs := NewFileSystem()

	// Test listing folders for a user that doesn't exist
	_, err := fs.ListFolders("test_user", "--sort-name", "asc")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test listing folders with an invalid sort order
	err = fs.Register("test_user")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	_, err = fs.ListFolders("test_user", "--sort-name", "invalid_sort_order")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test listing folders with no folders
	result, err := fs.ListFolders("test_user", "--sort-name", "asc")
	if err == nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	if result != "" {
		t.Errorf("Expected an empty string but got '%s'", result)
	}

	// Test listing folders with folders
	err = fs.CreateFolder("test_user", "folder1", "description1")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	err = fs.CreateFolder("test_user", "folder2", "description2")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	result, err = fs.ListFolders("test_user", "--sort-name", "asc")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	expectedResult := "folder1 description1 " + time.Now().Format("2006-01-02 15:04:05") + " test_user\nfolder2 description2 " + time.Now().Format("2006-01-02 15:04:05") + " test_user"
	if result != expectedResult {
		t.Errorf("Expected '%s' but got '%s'", expectedResult, result)
	}
}

func TestDeleteFolder(t *testing.T) {
	fs := NewFileSystem()

	// Test deleting a folder for a user that doesn't exist
	err := fs.DeleteFolder("test_user", "test_folder")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test deleting a folder that doesn't exist
	err = fs.Register("test_user")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	err = fs.DeleteFolder("test_user", "test_folder")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test deleting a folder that exists
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	err = fs.DeleteFolder("test_user", "test_folder")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
}

func TestRenameFolder(t *testing.T) {
	fs := NewFileSystem()

	// Test renaming a folder for a user that doesn't exist
	err := fs.RenameFolder("test_user", "test_folder", "new_test_folder")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test renaming a folder that doesn't exist
	err = fs.Register("test_user")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	err = fs.RenameFolder("test_user", "test_folder", "new_test_folder")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test renaming a folder with an invalid name
	err = fs.CreateFolder("test_user", "test_folder", "test_description")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
	err = fs.RenameFolder("test_user", "test_folder", "new_test_fold/er")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test renaming a folder with a name that is too long
	longFolderName := "new_test_folder_with_a_very_long_name_that_is_over_one_hundred_characters_new_test_folder_with_a_very_long_name_that_is_over_one_hundred_charactersnew_test_folder_with_a_very_long_name_that_is_over_one_hundred_characters_long_name_that_is_over_one_hundred_characters_long_name_that_is_over_one_hundred_characters_long_name_that_is_over_one_hundred_characters"
	err = fs.RenameFolder("test_user", "test_folder", longFolderName)
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	// Test renaming a folder with a valid name
	err = fs.RenameFolder("test_user", "test_folder", "new_test_folder")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}
}

func TestFileSystem_getUserByUsername(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		fs   *FileSystem
		args args
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fs.getUserByUsername(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileSystem.getUserByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_getFolderByName(t *testing.T) {
	type args struct {
		foldername string
	}
	tests := []struct {
		name string
		u    *User
		args args
		want *Folder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.getFolderByName(tt.args.foldername); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.getFolderByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_isFolderExists(t *testing.T) {
	type args struct {
		foldername string
	}
	tests := []struct {
		name string
		u    *User
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.isFolderExists(tt.args.foldername); got != tt.want {
				t.Errorf("User.isFolderExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
