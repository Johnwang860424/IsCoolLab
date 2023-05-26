package validate

import (
	"regexp"
)

func ValidateNoInvalidChars(name string) bool {
	regex := `[\\/:*?"<>|\s]`

	match, _ := regexp.MatchString(regex, name)
	return match
}

func ValidateLength(name string, length int) bool {
	return len(name) > length
}
