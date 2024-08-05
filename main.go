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

	attrs       *AttrList
	children    []GotmlTree
	textContent string
}

type Bag = map[string]interface{}

type AttrNode struct {
	Key   string
	Value interface{}
}

type Component func(*AttrList, ...GotmlTree) GotmlTree

type AttrList struct {
	this AttrNode
	next *AttrList
}

func Tree(tagName any) GotmlTree {
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
	}
}

func (g GotmlTree) Attr(key string, value interface{}) GotmlTree {
	copied := g

	copied.attrs = &AttrList{
		this: AttrNode{Key: key, Value: value},
		next: copied.attrs,
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
			fmt.Fprintf(os.Stderr, "Unsupported child type. Perhaps you need AsAny(...).")
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

	if tree.attrs != nil && isFragment {
		fmt.Fprintf(os.Stderr, "Document fragments cannot have attributes")
	} else {
		cur := tree.attrs
		for cur != nil {
			name := cur.this.Key
			val := cur.this.Value

			result += " "
			result += name
			result += "=\""
			valAsString, ok := any(val).(string)
			if !ok {
				fmt.Fprintf(os.Stderr, "Could not serialize attribute %s", name)
			}
			result += html.EscapeString(valAsString)
			result += "\""

			cur = cur.next
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

func Attr(k string, v interface{}) AttrNode {
	return AttrNode{Key: k, Value: v}
}

func (list *AttrList) ToBag() Bag {
	b := Bag{}
	cur := list
	for cur != nil {
		b[cur.this.Key] = cur.this.Value
		cur = cur.next
	}
	return b
}

func AsAny(children []GotmlTree) []any {
	r := make([]any, len(children))
	for i, c := range children {
		r[i] = c
	}
	return r
}
