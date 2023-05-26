package controller

import (
	"testing"
)

func TestNewFileSystem(t *testing.T) {
	fs := NewFileSystem()

	if fs == nil {
		t.Errorf("Expected a non-nil FileSystem object but got nil")
	}

	if fs.Users == nil {
		t.Errorf("Expected a non-nil Users map but got nil")
	}
}
