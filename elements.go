package gohtml

func P(attrs ...Attribute) Element {
	return func(inner string) string {
		return FormatElement("p", attrs, inner)
	}
}

func Div(attrs ...Attribute) Element {
	return func(inner string) string {
		return FormatElement("div", attrs, inner)
	}
}
