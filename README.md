# dom

A Go library for DOM manipulation, designed for use in WebAssembly (WASM) applications.

## Basic Usage

Here is a simple example of how to create a new `div` element and add it to the document's body.

```go
package main

import (
    "github.com/your-username/dom" // Adjust the import path
    "syscall/js"
)

func main() {
    // 1. Get a container element from the DOM.
    // The internal `query` function can be exposed or a helper can be used.
    // For this example, let's assume a helper `dom.Query()` exists.
    body, err := dom.Query("body")
    if err != nil {
        // Handle error: the body element was not found.
        return
    }

    // 2. Create an `HtmlElement` struct.
    newElement := dom.HtmlElement{
        Container: body,
        Name:      "div",
        Id:        "my-new-div",
        Class:     "my-class",
        Content:   "<h1>Hello from Go!</h1>",
    }

    // 3. Add the new element to the DOM.
    newElement.Add()
}

```

*Note: The `dom.Query` function is an example of how you might expose the internal `query` function for use in your application.*

## Documentation

For a more detailed guide to the API and its functions, please see the [DOM API Documentation](./docs/dom-api.md).

We also have a document outlining potential [improvements to the library](./docs/improvements.md).