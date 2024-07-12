package gohtml

import (
	"strings"
)

func Class_(value ...string) Attribute {
	return func() string {
		return formatAttribute("class", strings.Join(value, " "))
	}
}

func Id_(value string) Attribute {
	return func() string {
		return formatAttribute("id", value)
	}
}
