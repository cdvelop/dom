# Proposal for Library Improvements

After reviewing the `dom` library, here are some suggestions for improvements that could enhance its functionality, usability, and robustness.

## 1. Richer Element Creation API

The current `HtmlElement` struct and its `Add()` method are a good start, but the API could be more expressive and less error-prone.

### Suggestion: A Fluent API for Building Elements

A fluent API would allow for more readable and chainable element construction.

**Example of the proposed API:**

```go
// Proposed API
div := dom.NewElement("div").
    ID("my-div").
    Class("container").
    Content("<h2>New Section</h2>").
    AppendTo(body)

p := dom.NewElement("p").
    Text("This is a paragraph.").
    AppendTo(div)
```

This would involve:
- A `NewElement(tagName string)` constructor function.
- Methods like `ID()`, `Class()`, `Content()`, `Text()`, `Attr()` for setting attributes.
- Methods like `AppendTo(parent js.Value)` or `PrependTo(parent js.Value)` to attach the element to the DOM.

## 2. Improved Event Handling

The current event handling seems to be spread across different files and is not very generic.

### Suggestion: A Generic `On()` Method

A generic `On(eventName string, handler func(event js.Value))` method would be more flexible.

**Example of the proposed API:**

```go
button := dom.Query("#my-button")

button.On("click", func(event js.Value) {
    fmt.Println("Button clicked!")
    // Prevent default behavior
    event.Call("preventDefault")
})
```

This would provide a unified way to handle any DOM event.

## 3. Optimized Error and String Handling

To maintain a small binary size, which is crucial for WebAssembly, the library should avoid parts of the standard library that are heavy, such as `fmt`, `errors`, and `strconv`.

### Suggestion: Use `cdvelop/tinystring`

The `cdvelop/tinystring` library is optimized for creating small binaries and provides functionality for string manipulation, number conversion, and error handling.

Instead of returning `(..., string)` for errors, or using the standard `error` interface, functions should adopt the error handling mechanism provided by `cdvelop/tinystring`. This will ensure consistency and keep the final WASM binary as small as possible.

Similarly, any internal string or number conversions should also use `cdvelop/tinystring`.

## 4. Comprehensive Documentation

The inline code comments are minimal.

### Suggestion: Add GoDoc Comments

Add GoDoc-style comments to all public functions and structs to enable automated documentation generation.

## 5. Unit and Integration Testing

There are no tests in the repository.

### Suggestion: Implement a Test Suite

- **Unit Tests**: For individual functions (e.g., testing the logic of the fluent API builders).
- **Integration Tests**: Using a headless browser environment (like `playwright-go` or others) to test the DOM manipulation in a real browser context.
