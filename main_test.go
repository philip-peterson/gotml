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

func TestNestedElements(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("div").Children(
			Tree("p").Children("Paragraph 1"),
			Tree("div").Children(
				Tree("span").Children("Nested span"),
			),
		)
	}

	myTree := Tree(App)
	expected := "<div><p>Paragraph 1</p><div><span>Nested span</span></div></div>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

func TestAttributesHandling(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("button").Attr("id", "submit-btn").Children("Submit")
	}

	myTree := Tree(App)
	expected := "<button id=\"submit-btn\">Submit</button>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

func TestEmptyComponent(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("div") // No children, no attributes
	}

	myTree := Tree(App)
	expected := "<div />"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

func TestMultipleChildren(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("section").Children(
			Tree("header").Children("Header content"),
			Tree("article").Children("Article content"),
			Tree("footer").Children("Footer content"),
		)
	}

	myTree := Tree(App)
	expected := "<section><header>Header content</header><article>Article content</article><footer>Footer content</footer></section>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

func TestTextContentOnly(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("div").Children("Just text")
	}

	myTree := Tree(App)
	expected := "<div>Just text</div>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

func TestEmptyChildren(t *testing.T) {
	ctx := Bag{}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("ul").Children(
			Tree("li").Children("Item 1"),
			Tree("li").Children("Item 2"),
		)
	}

	myTree := Tree(App)
	expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}

func TestNestedComponents(t *testing.T) {
	ctx := Bag{}

	var MainContent Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree("div").Attr("class", "main-content").Children(
			Tree("p").Children(AsAny(children)...),
		)
	}

	var App Component = func(attrs *AttrList, children ...GotmlTree) GotmlTree {
		return Tree(MainContent).Children(
			Tree("section").Children("This is the main content."),
		)
	}

	myTree := Tree(App)
	expected := "<div class=\"main-content\"><p><section>This is the main content.</section></p></div>"
	result := Render(ctx, myTree)

	assert.Equal(t, expected, result)
}
