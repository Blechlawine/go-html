package gohtml

import "fmt"

func FormatElement(name string, attributes []Attribute, content string) string {
	attrString := ""
	for _, attr := range attributes {
		attrString += attr()
	}
	return fmt.Sprintf("<%s %s>%s</%s>", name, attrString, content, name)
}

func FormatAttribute(name string, value string) string {
	return fmt.Sprintf("%s=\"%s\"", name, value)
}
