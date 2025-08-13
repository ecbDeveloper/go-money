package shared

import (
	"net/mail"
	"strings"
	"unicode/utf8"
)

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}
