package gohtml

import (
	"fmt"
	"strings"
)

type Component func(inner string) string
type Attribute func() string

func p(attrs ...Attribute) Component {
	attrString := ""
	for _, attr := range attrs {
		attrString += attr()
	}
	return func(inner string) string {
		return fmt.Sprintf("<p %s>%s</p>", attrString, inner)
	}
}

func class_(value ...string) Attribute {
	return func() string {
		return fmt.Sprintf("class=\"%s\"", strings.Join(value, " "))
	}
}
