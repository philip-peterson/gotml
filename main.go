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

func Gotml(tagName any, instructions ...any) GotmlTree {
	var realChildren []GotmlTree
	attrs := Bag{}

	attrsFinished := false
	for _, inst := range instructions {
		switch any(inst).(type) {
		case SetAttr:
			// Set attribute instruction
			if attrsFinished {
				fmt.Fprintf(os.Stderr, "Attribute found in non-attribute position")
			}
			casted := any(inst).(SetAttr)
			attrs[casted.K] = casted.V
		case GotmlTree:
			// Child instruction
			attrsFinished = true
			casted := any(inst).(GotmlTree)
			realChildren = append(realChildren, casted)
		case string:
			// Child instruction (text node)
			attrsFinished = true
			casted := any(inst).(string)
			t := "#text"
			v := GotmlTree{
				terminalTagName: &t,
				textContent:     casted,
			}
			realChildren = append(realChildren, v)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported instruction type")
		}
	}

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
