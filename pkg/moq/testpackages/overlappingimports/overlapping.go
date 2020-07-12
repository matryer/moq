package template

import (
	"html/template"
	custom "text/template"

	"github.com/matryer/moq/pkg/moq/testpackages/overlappingimports/one"
)

type Example interface {
	HTML() template.HTML
	Text() custom.Template
	Self() Example
	MultiFile() one.Example
}
