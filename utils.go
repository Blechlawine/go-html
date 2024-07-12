package gohtml

import "fmt"

func formatElement(name string, attributes []Attribute, content string) string {
	attrString := ""
	for _, attr := range attributes {
		attrString += attr()
	}
	return fmt.Sprintf("<%s %s>%s</%s>", name, attrString, content, name)
}

func formatAttribute(name string, value string) string {
	return fmt.Sprintf("%s=\"%s\"", name, value)
}
