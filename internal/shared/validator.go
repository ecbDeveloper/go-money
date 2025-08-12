package shared

import (
	"context"
	"net/mail"
	"strings"
	"unicode/utf8"
)

type Validator interface {
	Valid(context.Context) ErrorsValidator
}

type ErrorsValidator map[string]string

func (e *ErrorsValidator) AddFieldError(key, message string) {
	if *e == nil {
		*e = make(map[string]string)
	}

	if _, exists := (*e)[key]; !exists {
		(*e)[key] = message
	}
}

func (e *ErrorsValidator) CheckField(ok bool, key, message string) {
	if !ok {
		e.AddFieldError(key, message)
	}
}

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
