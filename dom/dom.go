//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"

	"github.com/abdorrahmani/go-wasm/js"
)

// Document represents the DOM document
type Document struct {
	Value *js.Value
}

// Global returns the global document object
func Global() *Document {
	return &Document{
		Value: js.Global().Get("document"),
	}
}

// CreateElement creates a new element with the given tag name
func (d *Document) CreateElement(tagName string) *Element {
	return &Element{
		Value: d.Value.Call("createElement", tagName),
	}
}

// GetElementByID returns an element by its ID
func (d *Document) GetElementByID(id string) *Element {
	return &Element{
		Value: d.Value.Call("getElementById", id),
	}
}

// QuerySelector returns the first element matching the selector
func (d *Document) QuerySelector(selector string) *Element {
	return &Element{
		Value: d.Value.Call("querySelector", selector),
	}
}

// QuerySelectorAll returns all elements matching the selector
func (d *Document) QuerySelectorAll(selector string) []*Element {
	value := d.Value.Call("querySelectorAll", selector)
	length := value.MustLength()
	elements := make([]*Element, length)
	for i := 0; i < length; i++ {
		elements[i] = &Element{
			Value: value.Get(fmt.Sprintf("%d", i)),
		}
	}
	return elements
}

// CreateTextNode creates a new text node
func (d *Document) CreateTextNode(text string) *Node {
	return &Node{
		Value: d.Value.Call("createTextNode", text),
	}
}

// GetBody returns the document body element
func (d *Document) GetBody() *Element {
	return &Element{
		Value: d.Value.Get("body"),
	}
}

// GetHead returns the document head element
func (d *Document) GetHead() *Element {
	return &Element{
		Value: d.Value.Get("head"),
	}
}

// Title returns the document title
func (d *Document) Title() string {
	return d.Value.Get("title").MustString()
}

// SetTitle sets the document title
func (d *Document) SetTitle(title string) {
	d.Value.Set("title", title)
}

// URL returns the current document URL
func (d *Document) URL() string {
	return d.Value.Get("URL").MustString()
}

// ReadyState returns the document's ready state
func (d *Document) ReadyState() string {
	return d.Value.Get("readyState").MustString()
}

// IsReady returns true if the document is fully loaded
func (d *Document) IsReady() bool {
	return d.ReadyState() == "complete"
}
