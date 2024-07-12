package gohtml

type Props map[string]string
type Component func(props Props) string
type Element func(inner string) string
type Attribute func() string
