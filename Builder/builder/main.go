package main

import (
	"fmt"
	"strings"
)

// It ends up being better to build more complex structures that build upon the basic ones, that can still utilize the builder paradigm
type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

// String implementation abstraction for string access publicly
func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat("", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*indent+1))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

// Builder constructor
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{
			rootName,
			"",
			[]HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

const (
	indentSize = 2
)

// Now, as a fluent interface, with this you can chain calls together
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}
func main() {
	// some basic data structures, like strings, already come with a builder built in the basic Go sdk

	hello := "Hello"
	sb := strings.Builder{}

	// This API adds strings to the builder, which means that it pushes strings to the buffer that will end up becoming our intended resulting string
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	words := []string{"Hello", "world"}
	sb.Reset()
	// build a UL list <ul><li>...</li></ul>
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("<li>")
	}
	sb.WriteString("<ul>")
	fmt.Println(sb.String())

	// The builder only cares about utility calls, not anything else. The
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "Hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	b_fluent := NewHtmlBuilder("ul")
	b_fluent.AddChildFluent("li", "Hello").AddChildFluent("li", "world").AddChildFluent("li", "fluent")
	fmt.Println(b_fluent.String())
	// a Fluent interface is an interface that allows you to chain calls together
}
