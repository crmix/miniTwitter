package utils

import "regexp"

var uzbPhoneRegex = regexp.MustCompile(`[+]{1}99{1}[0-9]{10}$`)

func IsPhoneValid(p string) bool {
	return uzbPhoneRegex.MatchString(p)
}