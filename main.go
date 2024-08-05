// You can edit this code!
// Click here and start typing.
package gotml

import (
	"fmt"
	html "html"
	"os"
)

type GotmlTree struct {
	terminalTagName *string
	gotmlFunc       *Component

	attrs       map[string]interface{}
	children    []GotmlTree
	textContent string
}

type Bag = map[string]interface{}

type SetAttr struct {
	K string
	V interface{}
}

type Component func(Bag, ...GotmlTree) GotmlTree

func Gotml(tagName any) GotmlTree {
	var realChildren []GotmlTree
	attrs := Bag{}

	var terminalTagName *string
	var gotmlFunc *Component

	switch any(tagName).(type) {
	case string:
		m := any(tagName).(string)
		if m == "#text" {
			fmt.Fprintf(os.Stderr, "Text is not allowed as the top level. Use #fragment instead.")
		}
		terminalTagName = &m
	default:
		casted, ok := any(tagName).(Component)
		if !ok {
			fmt.Fprintf(os.Stderr, "Unsupported tag type. Perhaps you forgot to declare it as a Component?")
		} else {
			gotmlFunc = &casted
		}
	}

	return GotmlTree{
		terminalTagName: terminalTagName,
		gotmlFunc:       gotmlFunc,
		children:        realChildren,
		attrs:           attrs,
	}
}

func (g GotmlTree) Attrs(attrs ...SetAttr) GotmlTree {
	copied := g
	copied.attrs = map[string]interface{}{}

	for _, inst := range attrs {
		copied.attrs[inst.K] = inst.V
	}

	return copied
}

func (g GotmlTree) Children(children ...any) GotmlTree {
	copied := g
	copied.children = []GotmlTree{}

	for _, inst := range children {
		switch any(inst).(type) {
		case GotmlTree:
			// Child instruction
			casted := any(inst).(GotmlTree)
			copied.children = append(copied.children, casted)
		case string:
			// Child instruction (text node)
			casted := any(inst).(string)
			t := "#text"
			v := GotmlTree{
				terminalTagName: &t,
				textContent:     casted,
			}
			copied.children = append(copied.children, v)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported child type")
		}
	}

	return copied
}

func Render(ctx Bag, tree GotmlTree) string {
	if tree.gotmlFunc != nil {
		f := *tree.gotmlFunc
		return Render(
			ctx,
			f(tree.attrs, tree.children...),
		)
	}

	tagName := *tree.terminalTagName

	if tagName == "#text" {
		return html.EscapeString(tree.textContent)
	}

	isFragment := tagName == "#fragment"
	result := ""
	if !isFragment {
		result += "<" + tagName
	}

	if len(tree.attrs) > 0 && isFragment {
		fmt.Fprintf(os.Stderr, "Document fragments cannot have attributes")
	} else {
		for name, val := range tree.attrs {
			result += " "
			result += name
			result += "=\""
			valAsString, ok := any(val).(string)
			if !ok {
				fmt.Fprintf(os.Stderr, "Could not serialize attribute %s", name)
			}
			result += html.EscapeString(valAsString)
			result += "\""
		}
	}

	if len(tree.children) > 0 {
		if !isFragment {
			result += ">"
		}
		for _, child := range tree.children {
			result += Render(ctx, child)
		}
		if !isFragment {
			result += "</"
			result += tagName
			result += ">"
		}
		return result
	}

	return result + " />"
}

func Attr(k string, v interface{}) SetAttr {
	return SetAttr{K: k, V: v}
}
