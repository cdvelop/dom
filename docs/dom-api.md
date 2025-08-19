# DOM API Documentation

This document provides a detailed overview of the `dom` library's public API.

## Core Concepts

The library is designed to manipulate the Document Object Model (DOM) from Go code, which is particularly useful in WebAssembly (WASM) applications.

### The `HtmlElement` Struct

The core of element creation is the `HtmlElement` struct. It is defined as follows:

```go
type HtmlElement struct {
    Container js.Value // The parent element to which the new element will be appended.
    Name      string   // The tag name of the element (e.g., "div", "li", "p").
    Id        string   // The ID of the element.
    Class     string   // The CSS class (or classes) for the element.
    Content   string   // The inner HTML content of the element.
}
```

### Adding an Element to the DOM

To add an element to the DOM, you first create an instance of `HtmlElement` and then call the `Add()` method on it.

```go
func (e HtmlElement) Add()
```

**Example:**

```go
// Get the body element to serve as a container.
container, err := query("body")
if err != nil {
    // Handle error
}

// Create a new div element.
newElement := HtmlElement{
    Container: container,
    Name:      "div",
    Id:        "my-new-div",
    Class:     "my-class",
    Content:   "<h1>Hello, World!</h1>",
}

// Add the element to the DOM.
newElement.Add()
```

## Querying the DOM

The library provides several functions to select elements from the DOM. These functions are methods of the `Dom` struct.

### `query(selector string) (js.Value, string)`

This is an internal function that serves as the basis for many selection operations. It uses `document.querySelector` to find the first element that matches the specified CSS selector.

- **`selector`**: A string containing a CSS selector.
- **Returns**: A `js.Value` representing the found element and an error string if the element is not found.

### `(d Dom) GetHtmlModule(module_name string) (any, string)`

This function is specialized for finding a module's container element, which is expected to be a `div` with an ID matching the `module_name`.

- **`module_name`**: The name of the module, which is used as the element's ID.
- **Returns**: The `js.Value` of the module's container element and an error string if not found.

### `(d Dom) SelectContent(o model.SelectDomOptions) (any, string)`

This function selects an element and retrieves content from it. The `model.SelectDomOptions` struct offers fine-grained control over the operation.

- **`o`**: An instance of `model.SelectDomOptions`.

### `(d Dom) ElementClicking(querySelector string) (err string)`

Finds an element using the provided selector and simulates a `click` event on it.

- **`querySelector`**: The CSS selector for the target element.
- **Returns**: An error string if the element is not found or the click fails.

### `(d Dom) ElementFocus(querySelector string) (err string)`

Finds an element using the provided selector and gives it focus.

- **`querySelector`**: The CSS selector for the target element.
- **Returns**: An error string if the element is not found or the focus action fails.

## DOM Manipulation

### `(d Dom) ToggleClass(elem js.Value, className string)`

Toggles a CSS class on a given element.

- **`elem`**: The `js.Value` of the element to modify.
- **`className`**: The name of the class to toggle.
