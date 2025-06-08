//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"

	"github.com/abdorrahmani/go-wasm/js"
)

// Node represents a DOM node
type Node struct {
	Value *js.Value
}

// NodeType constants
const (
	ElementNode               = 1
	AttributeNode             = 2
	TextNode                  = 3
	CDATASectionNode          = 4
	EntityReferenceNode       = 5
	EntityNode                = 6
	ProcessingInstructionNode = 7
	CommentNode               = 8
	DocumentNode              = 9
	DocumentTypeNode          = 10
	DocumentFragmentNode      = 11
	NotationNode              = 12
)

// GetNodeType returns the node type
func (n *Node) GetNodeType() int {
	return n.Value.Get("nodeType").MustInt()
}

// GetNodeName returns the node name
func (n *Node) GetNodeName() string {
	return n.Value.Get("nodeName").MustString()
}

// GetNodeValue returns the node value
func (n *Node) GetNodeValue() string {
	return n.Value.Get("nodeValue").MustString()
}

// SetNodeValue sets the node value
func (n *Node) SetNodeValue(value string) {
	n.Value.Set("nodeValue", value)
}

// GetParentNode returns the parent node
func (n *Node) GetParentNode() *Node {
	return &Node{
		Value: n.Value.Get("parentNode"),
	}
}

// GetChildNodes returns all child nodes
func (n *Node) GetChildNodes() []*Node {
	value := n.Value.Get("childNodes")
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
func (n *Node) GetFirstChild() *Node {
	return &Node{
		Value: n.Value.Get("firstChild"),
	}
}

// GetLastChild returns the last child node
func (n *Node) GetLastChild() *Node {
	return &Node{
		Value: n.Value.Get("lastChild"),
	}
}

// GetPreviousSibling returns the previous sibling node
func (n *Node) GetPreviousSibling() *Node {
	return &Node{
		Value: n.Value.Get("previousSibling"),
	}
}

// GetNextSibling returns the next sibling node
func (n *Node) GetNextSibling() *Node {
	return &Node{
		Value: n.Value.Get("nextSibling"),
	}
}

// GetOwnerDocument returns the owner document
func (n *Node) GetOwnerDocument() *Document {
	return &Document{
		Value: n.Value.Get("ownerDocument"),
	}
}

// HasChildNodes checks if the node has child nodes
func (n *Node) HasChildNodes() bool {
	return n.Value.Call("hasChildNodes").MustBool()
}

// CloneNode creates a copy of the node
func (n *Node) CloneNode(deep bool) *Node {
	return &Node{
		Value: n.Value.Call("cloneNode", deep),
	}
}

// CompareDocumentPosition compares the position of two nodes
func (n *Node) CompareDocumentPosition(other *Node) int {
	return n.Value.Call("compareDocumentPosition", other.Value).MustInt()
}

// Contains checks if the node contains another node
func (n *Node) Contains(other *Node) bool {
	return n.Value.Call("contains", other.Value).MustBool()
}

// InsertBefore inserts a node before a reference node
func (n *Node) InsertBefore(newNode, referenceNode *Node) error {
	if n.Value == nil || n.Value.Raw().IsNull() || n.Value.Raw().IsUndefined() {
		return fmt.Errorf("node is nil or undefined/null")
	}
	if newNode == nil || newNode.Value == nil || newNode.Value.Raw().IsNull() || newNode.Value.Raw().IsUndefined() {
		return fmt.Errorf("new node is nil or undefined/null")
	}
	if referenceNode == nil || referenceNode.Value == nil || referenceNode.Value.Raw().IsNull() || referenceNode.Value.Raw().IsUndefined() {
		return fmt.Errorf("reference node is nil or undefined/null")
	}
	n.Value.Call("insertBefore", newNode.Value, referenceNode.Value)
	return nil
}

// ReplaceChild replaces a child node
func (n *Node) ReplaceChild(newNode, oldNode *Node) error {
	if n.Value == nil || n.Value.Raw().IsNull() || n.Value.Raw().IsUndefined() {
		return fmt.Errorf("node is nil or undefined/null")
	}
	if newNode == nil || newNode.Value == nil || newNode.Value.Raw().IsNull() || newNode.Value.Raw().IsUndefined() {
		return fmt.Errorf("new node is nil or undefined/null")
	}
	if oldNode == nil || oldNode.Value == nil || oldNode.Value.Raw().IsNull() || oldNode.Value.Raw().IsUndefined() {
		return fmt.Errorf("old node is nil or undefined/null")
	}
	n.Value.Call("replaceChild", newNode.Value, oldNode.Value)
	return nil
}

// RemoveChild removes a child node
func (n *Node) RemoveChild(child *Node) error {
	if n.Value == nil || n.Value.Raw().IsNull() || n.Value.Raw().IsUndefined() {
		return fmt.Errorf("node is nil or undefined/null")
	}
	if child == nil || child.Value == nil || child.Value.Raw().IsNull() || child.Value.Raw().IsUndefined() {
		return fmt.Errorf("child node is nil or undefined/null")
	}
	n.Value.Call("removeChild", child.Value)
	return nil
}

// AppendChild appends a child node
func (n *Node) AppendChild(child *Node) error {
	if n.Value == nil || n.Value.Raw().IsNull() || n.Value.Raw().IsUndefined() {
		return fmt.Errorf("node is nil or undefined/null")
	}
	if child == nil || child.Value == nil || child.Value.Raw().IsNull() || child.Value.Raw().IsUndefined() {
		return fmt.Errorf("child node is nil or undefined/null")
	}
	n.Value.Call("appendChild", child.Value)
	return nil
}

// Normalize normalizes the node's text nodes
func (n *Node) Normalize() {
	n.Value.Call("normalize")
}

// IsDefaultNamespace checks if the node is in the default namespace
func (n *Node) IsDefaultNamespace(namespaceURI string) bool {
	return n.Value.Call("isDefaultNamespace", namespaceURI).MustBool()
}

// LookupNamespaceURI looks up the namespace URI for a prefix
func (n *Node) LookupNamespaceURI(prefix string) string {
	return n.Value.Call("lookupNamespaceURI", prefix).MustString()
}

// LookupPrefix looks up the prefix for a namespace URI
func (n *Node) LookupPrefix(namespaceURI string) string {
	return n.Value.Call("lookupPrefix", namespaceURI).MustString()
}

// IsEqualNode checks if two nodes are equal
func (n *Node) IsEqualNode(other *Node) bool {
	return n.Value.Call("isEqualNode", other.Value).MustBool()
}

// IsSameNode checks if two nodes are the same
func (n *Node) IsSameNode(other *Node) bool {
	return n.Value.Call("isSameNode", other.Value).MustBool()
}

// GetBaseURI returns the node's base URI
func (n *Node) GetBaseURI() string {
	return n.Value.Get("baseURI").MustString()
}

// GetTextContent returns the node's text content
func (n *Node) GetTextContent() string {
	return n.Value.Get("textContent").MustString()
}

// SetTextContent sets the node's text content
func (n *Node) SetTextContent(text string) {
	n.Value.Set("textContent", text)
}
