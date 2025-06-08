//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"

	"github.com/abdorrahmani/go-wasm/js"
)

// Element represents a DOM element
type Element struct {
	Value *js.Value
}

// DOMRect represents a rectangle with position and dimensions
type DOMRect struct {
	Value *js.Value
}

// GetID returns the element's ID
func (e *Element) GetID() string {
	return e.Value.Get("id").MustString()
}

// SetID sets the element's ID
func (e *Element) SetID(id string) {
	e.Value.Set("id", id)
}

// GetClassName returns the element's class name
func (e *Element) GetClassName() string {
	return e.Value.Get("className").MustString()
}

// SetClassName sets the element's class name
func (e *Element) SetClassName(className string) {
	e.Value.Set("className", className)
}

// GetInnerHTML returns the element's inner HTML
func (e *Element) GetInnerHTML() string {
	return e.Value.Get("innerHTML").MustString()
}

// SetInnerHTML sets the element's inner HTML
func (e *Element) SetInnerHTML(html string) {
	e.Value.Set("innerHTML", html)
}

// GetTextContent returns the element's text content
func (e *Element) GetTextContent() string {
	return e.Value.Get("textContent").MustString()
}

// SetTextContent sets the element's text content
func (e *Element) SetTextContent(text string) {
	e.Value.Set("textContent", text)
}

// GetAttribute returns the value of an attribute
func (e *Element) GetAttribute(name string) string {
	return e.Value.Call("getAttribute", name).MustString()
}

// SetAttribute sets an attribute value
func (e *Element) SetAttribute(name, value string) {
	e.Value.Call("setAttribute", name, value)
}

// RemoveAttribute removes an attribute
func (e *Element) RemoveAttribute(name string) {
	e.Value.Call("removeAttribute", name)
}

// HasAttribute checks if an attribute exists
func (e *Element) HasAttribute(name string) bool {
	return e.Value.Call("hasAttribute", name).MustBool()
}

// GetStyle returns the element's style object
func (e *Element) GetStyle() *Style {
	return &Style{
		Value: e.Value.Get("style"),
	}
}

// AppendChild appends a child node
func (e *Element) AppendChild(child *Node) error {
	if e.Value == nil || e.Value.Raw().IsNull() || e.Value.Raw().IsUndefined() {
		return fmt.Errorf("element is nil or undefined/null")
	}
	if child == nil || child.Value == nil || child.Value.Raw().IsNull() || child.Value.Raw().IsUndefined() {
		return fmt.Errorf("child node is nil or undefined/null")
	}
	e.Value.Call("appendChild", child.Value.Raw())
	return nil
}

// RemoveChild removes a child node
func (e *Element) RemoveChild(child *Node) error {
	if e.Value == nil || e.Value.Raw().IsNull() || e.Value.Raw().IsUndefined() {
		return fmt.Errorf("element is nil or undefined/null")
	}
	if child == nil || child.Value == nil || child.Value.Raw().IsNull() || child.Value.Raw().IsUndefined() {
		return fmt.Errorf("child node is nil or undefined/null")
	}
	e.Value.Call("removeChild", child.Value)
	return nil
}

// GetChildNodes returns all child nodes
func (e *Element) GetChildNodes() []*Node {
	value := e.Value.Get("childNodes")
	length := value.MustLength()
	nodes := make([]*Node, length)
	for i := 0; i < length; i++ {
		nodes[i] = &Node{
			Value: value.Get(fmt.Sprintf("%d", i)),
		}
	}
	return nodes
}

// GetFirstChild returns the first child node
func (e *Element) GetFirstChild() *Node {
	return &Node{
		Value: e.Value.Get("firstChild"),
	}
}

// GetLastChild returns the last child node
func (e *Element) GetLastChild() *Node {
	return &Node{
		Value: e.Value.Get("lastChild"),
	}
}

// GetParentNode returns the parent node
func (e *Element) GetParentNode() *Node {
	return &Node{
		Value: e.Value.Get("parentNode"),
	}
}

// GetNextSibling returns the next sibling node
func (e *Element) GetNextSibling() *Node {
	return &Node{
		Value: e.Value.Get("nextSibling"),
	}
}

// GetPreviousSibling returns the previous sibling node
func (e *Element) GetPreviousSibling() *Node {
	return &Node{
		Value: e.Value.Get("previousSibling"),
	}
}

// AddEventListener adds an event listener
func (e *Element) AddEventListener(eventType string, handler func(*Event)) {
	callback := js.NewCallback(func(args []*js.Value) {
		handler(&Event{
			Value: args[0],
		})
	})
	e.Value.Call("addEventListener", eventType, callback)
}

// RemoveEventListener removes an event listener
func (e *Element) RemoveEventListener(eventType string, handler func(*Event)) {
	callback := js.NewCallback(func(args []*js.Value) {
		handler(&Event{
			Value: args[0],
		})
	})
	e.Value.Call("removeEventListener", eventType, callback)
}

// DispatchEvent dispatches an event
func (e *Element) DispatchEvent(event *Event) bool {
	return e.Value.Call("dispatchEvent", event.Value).MustBool()
}

// GetBoundingClientRect returns the element's bounding rectangle
func (e *Element) GetBoundingClientRect() *DOMRect {
	return &DOMRect{
		Value: e.Value.Call("getBoundingClientRect"),
	}
}

// ScrollIntoView scrolls the element into view
func (e *Element) ScrollIntoView() {
	e.Value.Call("scrollIntoView")
}

// Focus sets focus to the element
func (e *Element) Focus() {
	e.Value.Call("focus")
}

// Blur removes focus from the element
func (e *Element) Blur() {
	e.Value.Call("blur")
}

// Click simulates a click on the element
func (e *Element) Click() {
	e.Value.Call("click")
}

// GetTagName returns the element's tag name
func (e *Element) GetTagName() string {
	return e.Value.Get("tagName").MustString()
}

// GetNamespaceURI returns the element's namespace URI
func (e *Element) GetNamespaceURI() string {
	return e.Value.Get("namespaceURI").MustString()
}

// GetPrefix returns the element's prefix
func (e *Element) GetPrefix() string {
	return e.Value.Get("prefix").MustString()
}

// GetLocalName returns the element's local name
func (e *Element) GetLocalName() string {
	return e.Value.Get("localName").MustString()
}

// GetBaseURI returns the element's base URI
func (e *Element) GetBaseURI() string {
	return e.Value.Get("baseURI").MustString()
}

// GetOwnerDocument returns the element's owner document
func (e *Element) GetOwnerDocument() *Document {
	return &Document{
		Value: e.Value.Get("ownerDocument"),
	}
}

// GetNodeType returns the element's node type
func (e *Element) GetNodeType() int {
	return e.Value.Get("nodeType").MustInt()
}

// GetNodeName returns the element's node name
func (e *Element) GetNodeName() string {
	return e.Value.Get("nodeName").MustString()
}

// GetNodeValue returns the element's node value
func (e *Element) GetNodeValue() string {
	return e.Value.Get("nodeValue").MustString()
}

// SetNodeValue sets the element's node value
func (e *Element) SetNodeValue(value string) {
	e.Value.Set("nodeValue", value)
}

// CloneNode creates a copy of the element
func (e *Element) CloneNode(deep bool) *Element {
	return &Element{
		Value: e.Value.Call("cloneNode", deep),
	}
}

// CompareDocumentPosition compares the position of two nodes
func (e *Element) CompareDocumentPosition(other *Element) int {
	return e.Value.Call("compareDocumentPosition", other.Value).MustInt()
}

// Contains checks if the element contains another element
func (e *Element) Contains(other *Element) bool {
	return e.Value.Call("contains", other.Value).MustBool()
}

// HasChildNodes checks if the element has child nodes
func (e *Element) HasChildNodes() bool {
	return e.Value.Call("hasChildNodes").MustBool()
}

// InsertBefore inserts a node before a reference node
func (e *Element) InsertBefore(newNode, referenceNode *Node) error {
	if e.Value == nil {
		return fmt.Errorf("element is nil")
	}
	if newNode == nil || newNode.Value == nil {
		return fmt.Errorf("new node is nil")
	}
	if referenceNode == nil || referenceNode.Value == nil {
		return fmt.Errorf("reference node is nil")
	}
	e.Value.Call("insertBefore", newNode.Value, referenceNode.Value)
	return nil
}

// ReplaceChild replaces a child node
func (e *Element) ReplaceChild(newNode, oldNode *Node) error {
	if e.Value == nil || e.Value.Raw().IsNull() || e.Value.Raw().IsUndefined() {
		return fmt.Errorf("element is nil or undefined/null")
	}
	if newNode == nil || newNode.Value == nil || newNode.Value.Raw().IsNull() || newNode.Value.Raw().IsUndefined() {
		return fmt.Errorf("new node is nil or undefined/null")
	}
	if oldNode == nil || oldNode.Value == nil || oldNode.Value.Raw().IsNull() || oldNode.Value.Raw().IsUndefined() {
		return fmt.Errorf("old node is nil or undefined/null")
	}
	e.Value.Call("replaceChild", newNode.Value.Raw(), oldNode.Value.Raw())
	return nil
}

// Normalize normalizes the element's text nodes
func (e *Element) Normalize() {
	e.Value.Call("normalize")
}

// IsDefaultNamespace checks if the element is in the default namespace
func (e *Element) IsDefaultNamespace(namespaceURI string) bool {
	return e.Value.Call("isDefaultNamespace", namespaceURI).MustBool()
}

// LookupNamespaceURI looks up the namespace URI for a prefix
func (e *Element) LookupNamespaceURI(prefix string) string {
	return e.Value.Call("lookupNamespaceURI", prefix).MustString()
}

// LookupPrefix looks up the prefix for a namespace URI
func (e *Element) LookupPrefix(namespaceURI string) string {
	return e.Value.Call("lookupPrefix", namespaceURI).MustString()
}

// IsEqualNode checks if two nodes are equal
func (e *Element) IsEqualNode(other *Element) bool {
	return e.Value.Call("isEqualNode", other.Value).MustBool()
}

// IsSameNode checks if two nodes are the same
func (e *Element) IsSameNode(other *Element) bool {
	return e.Value.Call("isSameNode", other.Value).MustBool()
}

// GetTop returns the top position
func (r *DOMRect) GetTop() float64 {
	return r.Value.Get("top").MustFloat()
}

// GetRight returns the right position
func (r *DOMRect) GetRight() float64 {
	return r.Value.Get("right").MustFloat()
}

// GetBottom returns the bottom position
func (r *DOMRect) GetBottom() float64 {
	return r.Value.Get("bottom").MustFloat()
}

// GetLeft returns the left position
func (r *DOMRect) GetLeft() float64 {
	return r.Value.Get("left").MustFloat()
}

// GetWidth returns the width
func (r *DOMRect) GetWidth() float64 {
	return r.Value.Get("width").MustFloat()
}

// GetHeight returns the height
func (r *DOMRect) GetHeight() float64 {
	return r.Value.Get("height").MustFloat()
}

// GetX returns the x position
func (r *DOMRect) GetX() float64 {
	return r.Value.Get("x").MustFloat()
}

// GetY returns the y position
func (r *DOMRect) GetY() float64 {
	return r.Value.Get("y").MustFloat()
}
