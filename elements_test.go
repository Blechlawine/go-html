package gohtml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP(testing *testing.T) {
	element := p(Class_("hi", "hey"))("Hello")
	assert.Equal(testing, element, `<p class="hi hey">Hello</p>`)
}
