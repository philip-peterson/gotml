# Gotml

Gotml is a component-based HTML templating library for Go. It allows you to create and render HTML structures using a component-oriented approach, offering an alternative to Go's built-in template package. With Gotml, you can define reusable components and compose HTML elements in a more modular and expressive way.

## Overview

### Why Use Gotml?

- **Component-Based Approach**: Unlike Go templates, which are primarily text-based and less modular, Gotml enables you to define components as functions. This allows for better reuse and organization of HTML structures.

- **Dynamic Composition**: Easily compose complex HTML structures by combining simple components. Gotml supports variadic parameters for components, making it straightforward to include multiple children or attributes.

- **Improved Readability**: Gotml’s approach helps to maintain readability by separating concerns into manageable components, avoiding the verbosity and complexity often found in traditional template-based HTML generation.

### Basic Usage

Define a component with attributes and children:

```go
package gotml

import "github.com/philip-peterson/gotml"

var BorderedDiv Component = func(attrs map[string]interface{}, children ...gotml.GotmlTree) gotml.GotmlTree {
    return gotml.Gotml("div",
        gotml.Attr("style", "border: 3px double black"),
        children...,
    )
}
```

Render HTML from a component:

```go
package gotml

import (
    "fmt"
    "github.com/philip-peterson/gotml"
)

func main() {
    ctx := map[string]interface{}{}
    
    myTree := gotml.Gotml(BorderedDiv,
        gotml.Gotml("p", "Hello, world!"),
    )

    result := gotml.Render(ctx, myTree)
    fmt.Println(result)
}
```

For more detailed usage and examples, please see the [tests](#main_test.go).