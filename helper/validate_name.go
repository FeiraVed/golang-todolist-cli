package helper

import "strings"

func Validate(data string) bool {

	if strings.TrimSpace(data) == "" {
		return false
	}

	return true
}
