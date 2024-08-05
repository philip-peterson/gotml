# gotml

`gotml` is a component-based HTML templating library for Go. It allows you to create and render HTML structures using a component-oriented approach, offering an alternative to Go's built-in template package. With `gotml`, you can define reusable components and compose HTML elements in a more modular and expressive way.

## Overview

### Why Use `gotml`?

- **Component-Based Approach**: Unlike Go templates, which are primarily text-based and less modular, `gotml` enables you to define components as functions. This allows for better reuse and organization of HTML structures.

- **Reduce Nesting**: In Go templates, refactors can easily create merge conflicts due to deep nesting of HTML. In contrast, `gotml` 

- **Dynamic Composition**: Easily compose complex HTML structures by combining simple components. `gotml` supports variadic parameters for components, making it straightforward to include multiple children or attributes.

- **Improved Readability**: `gotml`’s approach helps to maintain readability by separating concerns into manageable components, avoiding the verbosity and complexity often found in traditional template-based HTML generation.

### Basic Usage

Define a component with attributes and children:

```go
package main

import (
	"fmt"

	g "github.com/philip-peterson/gotml"
)

func main() {
	var BorderedDiv g.Component = func(attrs *g.AttrList, children ...g.GotmlTree) g.GotmlTree {
		return g.Gotml("div").
			Attr("style", "border: 3px double black").
			Children(
				g.AsAny(children)...,
			)
	}
}
```

Note that due to Go's quirk where it cannot convert a slice of a type, `[]T` to a slice of any, `[]any` without help, we use `g.AsAny(...)` to perform this conversion.

Now, we can render the component to HTML:

```go
myTree := g.Gotml(BorderedDiv).Children(
    g.Gotml("p").Children("Hello, world!"),
)

result := g.Render(ctx, myTree)
fmt.Println(result)
```

Creating the output:

```
<div style="border: 3px double black"><p>Hello, world!</p></div>
```

The full program is therefore as follows:

```go
package main

import (
	"fmt"

	g "github.com/philip-peterson/gotml"
)

func main() {
	var BorderedDiv g.Component = func(attrs *g.AttrList, children ...g.GotmlTree) g.GotmlTree {
		return g.Gotml("div").
			Attr("style", "border: 3px double black").
			Children(
				g.AsAny(children)...,
			)
	}

	ctx := map[string]interface{}{}

	myTree := g.Gotml(BorderedDiv).Children(
		g.Gotml("p").Children("Hello, world!"),
	)

	result := g.Render(ctx, myTree)
	fmt.Println(result)
}
```

For more detailed usage and examples, please see the [tests](https://github.com/philip-peterson/gotml/blob/main/main_test.go).
