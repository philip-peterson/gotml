package gotml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicHTMLStructure(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("html").Children(
			Tree("head"),
			Tree("body").Children(
				Tree("div").
					Attr("style", "color: red").
					Children("Lorem ipsum"),
				Tree("hr"),
				Tree("div").Children("Hello world"),
			),
		)
	}

	myTree := Tree(App)
	expected := "<html><head /><body><div style=\"color: red\">Lorem ipsum</div><hr /><div>Hello world</div></body></html>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

// func TestNestedElements(t *testing.T) {
// 	ctx := Bag{}

// 	var App Component = func(attrs Bag, children ...GotmlTree) GotmlTree {
// 		return Gotml("div",
// 			Gotml("p", "Paragraph 1"),
// 			Gotml("div",
// 				Gotml("span", "Nested span"),
// 			),
// 		)
// 	}

// 	myTree := Gotml(App)
// 	expected := "<div><p>Paragraph 1</p><div><span>Nested span</span></div></div>"
// 	result := Render(ctx, myTree)

// 	assert.Equal(t, expected, result)
// }

// func TestAttributesHandling(t *testing.T) {
// 	ctx := Bag{}

// 	var App Component = func(attrs Bag, children ...GotmlTree) GotmlTree {
// 		return Gotml("button", Attr("id", "submit-btn"), "Submit")
// 	}

// 	myTree := Gotml(App)
// 	expected := "<button id=\"submit-btn\">Submit</button>"
// 	result := Render(ctx, myTree)

// 	assert.Equal(t, expected, result)
// }

// func TestEmptyComponent(t *testing.T) {
// 	ctx := Bag{}

// 	var App Component = func(attrs Bag, children ...GotmlTree) GotmlTree {
// 		return Gotml("div") // No children, no attributes
// 	}

// 	myTree := Gotml(App)
// 	expected := "<div />"
// 	result := Render(ctx, myTree)

// 	assert.Equal(t, expected, result)
// }

// func TestMultipleChildren(t *testing.T) {
// 	ctx := Bag{}

// 	var App Component = func(attrs Bag, children ...GotmlTree) GotmlTree {
// 		return Gotml("section",
// 			Gotml("header", "Header content"),
// 			Gotml("article", "Article content"),
// 			Gotml("footer", "Footer content"),
// 		)
// 	}

// 	myTree := Gotml(App)
// 	expected := "<section><header>Header content</header><article>Article content</article><footer>Footer content</footer></section>"
// 	result := Render(ctx, myTree)

// 	assert.Equal(t, expected, result)
// }

// func TestTextContentOnly(t *testing.T) {
// 	ctx := Bag{}

// 	var App Component = func(attrs Bag, children ...GotmlTree) GotmlTree {
// 		return Gotml("div", "Just text")
// 	}

// 	myTree := Gotml(App)
// 	expected := "<div>Just text</div>"
// 	result := Render(ctx, myTree)

// 	assert.Equal(t, expected, result)
// }

// func TestEmptyChildren(t *testing.T) {
// 	ctx := Bag{}

// 	var App Component = func(attrs Bag, children ...GotmlTree) GotmlTree {
// 		return Gotml("ul",
// 			Gotml("li", "Item 1"),
// 			Gotml("li", "Item 2"),
// 		)
// 	}

// 	myTree := Gotml(App)
// 	expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"
// 	result := Render(ctx, myTree)

// 	assert.Equal(t, expected, result)
// }
