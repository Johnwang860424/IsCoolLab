package controller

import (
	"testing"
)

func TestRegister(t *testing.T) {
	fs := NewFileSystem()

	err := fs.Register("test_user")
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err.Error())
	}

	err = fs.Register("test_user")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	err = fs.Register("test_user_with/invalid_chars")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}

	err = fs.Register("test_user_with_a_very_long_name_that_is_over_fifty_characters")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}
