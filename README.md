# gotml

<p align="center">
	<img src="https://raw.githubusercontent.com/philip-peterson/gotml/main/gotml_logo.png" />
</p>

`gotml` is a component-based HTML templating library for Go. It allows you to create and render HTML structures using a component-oriented approach, offering an alternative to Go's built-in template package. With `gotml`, you can define reusable components and compose HTML elements in a more modular and expressive way.

## Overview

### Disclaimer

This library was the result of a "weekend project" as it were. It should be considered a toy and little more, especially when there is [Gomponents](https://github.com/maragudk/gomponents). Before writing this project, I did not know Gomponents existed, and since discovering it, I think I can safely say you should use it instead of `gotml`. So, that said...

### Why use `gotml` over `html/template`?

- **Component-Based Approach**: Unlike Go templates, which are primarily text-based and less modular, `gotml` enables you to define components as functions. This allows for better reuse and organization of HTML structures.

- **Reduce Nesting**: In Go templates, refactors can easily create merge conflicts due to deep nesting of HTML. In contrast, `gotml` allows discrete components to be defined allowing blocks to be moved to the top level, reducing the average amount of nesting.

- **Dynamic Composition**: Easily compose complex HTML structures by combining simple components. `gotml` supports variadic parameters for components, making it straightforward to include multiple children or attributes.

- **Improved Readability**: `gotml`’s approach helps to maintain readability by separating concerns into manageable components, avoiding the verbosity and complexity often found in traditional template-based HTML generation.

### Basic Usage: Define a tree structure

You can define a simple tree structure just like HTML or other frameworks.

```go
package main

import (
	"fmt"

	g "github.com/philip-peterson/gotml"
)

func main() {
	ctx := g.Bag{}

	var App g.Component = func(attrs *g.AttrList, children ...g.GotmlTree) g.GotmlTree {
		return g.Tree("html").Children(
			g.Tree("head"),
			g.Tree("body").Children(
				g.Tree("div").
					Attr("style", "color: red").
					Children("Lorem ipsum"),
				g.Tree("hr"),
				g.Tree("div").Children("Hello world"),
			),
		)
	}

	myTree := g.Tree(App)
	result := g.Render(ctx, myTree)

	fmt.Println(result)
}
```

The output*:

```
<html>
   <head />
   <body>
      <div style="color: red">
         Lorem ipsum
      </div>
      <hr />
      <div>
         Hello world
      </div>
   </body>
</html>
```

### Advanced Usage: Passing through Children

Define a component with attributes and children:

```go
package main

import (
	"fmt"

	g "github.com/philip-peterson/gotml"
)

func main() {
	var BorderedDiv g.Component = func(attrs *g.AttrList, children ...g.GotmlTree) g.GotmlTree {
		return g.Tree("div").
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
myTree := g.Tree(BorderedDiv).Children(
    g.Tree("p").Children("Hello, world!"),
)

result := g.Render(ctx, myTree)
fmt.Println(result)
```

Creating the output*:

```
<div style="border: 3px double black">
   <p>
      Hello, world!
   </p>
</div>
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
		return g.Tree("div").
			Attr("style", "border: 3px double black").
			Children(
				g.AsAny(children)...,
			)
	}

	ctx := map[string]interface{}{}

	myTree := g.Tree(BorderedDiv).Children(
		g.Tree("p").Children("Hello, world!"),
	)

	result := g.Render(ctx, myTree)
	fmt.Println(result)
}
```

For more detailed usage and examples, please see the [tests](https://github.com/philip-peterson/gotml/blob/main/main_test.go).

### Final Words

After creating this package, I found out about [Gomponents](https://github.com/maragudk/gomponents). If you're doing serious work, you should probably just use that instead.

If you can think of a reason this library is preferable to Gomponents for some reason, please reach out to let me know by opening an issue. It would be interesting to hear if this different approach poses some benefit and warrants builting out the library further.

\* Output has been prettified for readability.
