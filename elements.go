package gohtml

func P(attrs ...Attribute) Element {
	return func(inner string) string {
		return formatElement("p", attrs, inner)
	}
}

func Div(attrs ...Attribute) Element {
	return func(inner string) string {
		return formatElement("div", attrs, inner)
	}
}
