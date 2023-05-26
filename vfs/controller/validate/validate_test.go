package validate

import (
	"testing"
)

func TestValidateNoInvalidChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid name",
			input:    "my_project",
			expected: false,
		},
		{
			name:     "invalid name with space",
			input:    "my project",
			expected: true,
		},
		{
			name:     "invalid name with forward slash",
			input:    "my/project",
			expected: true,
		},
		{
			name:     "invalid name with backslash",
			input:    "my\\project",
			expected: true,
		},
		{
			name:     "invalid name with colon",
			input:    "my:project",
			expected: true,
		},
		{
			name:     "invalid name with asterisk",
			input:    "my*project",
			expected: true,
		},
		{
			name:     "invalid name with question mark",
			input:    "my?project",
			expected: true,
		},
		{
			name:     "invalid name with double quote",
			input:    "my\"project",
			expected: true,
		},
		{
			name:     "invalid name with less than symbol",
			input:    "my<project",
			expected: true,
		},
		{
			name:     "invalid name with greater than symbol",
			input:    "my>project",
			expected: true,
		},
		{
			name:     "invalid name with vertical bar",
			input:    "my|project",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ValidateNoInvalidChars(test.input)
			if result != test.expected {
				t.Errorf("Expected %v but got %v for input %s", test.expected, result, test.input)
			}
		})
	}
}

func TestValidateLength(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected bool
	}{
		{
			name:     "valid name",
			input:    "my_project",
			length:   5,
			expected: true,
		},
		{
			name:     "invalid name",
			input:    "my_project",
			length:   20,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ValidateLength(test.input, test.length)
			if result != test.expected {
				t.Errorf("Expected %v but got %v for input %s and length %d", test.expected, result, test.input, test.length)
			}
		})
	}
}
