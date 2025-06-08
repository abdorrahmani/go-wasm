//go:build js && wasm
// +build js,wasm

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/abdorrahmani/go-wasm/dom"
	"github.com/abdorrahmani/go-wasm/js"
)

func waitForDOM() error {
	// Wait for document to be ready
	doc := js.Global().Get("document")
	if doc.IsNull() || doc.IsUndefined() {
		return fmt.Errorf("document is not available")
	}

	// Wait for document.body to be available
	body := doc.Get("body")
	if body.IsNull() || body.IsUndefined() {
		return fmt.Errorf("document.body is not available")
	}

	return nil
}

func TestDOM(t *testing.T) {
	fmt.Println("Starting DOM WebAssembly tests...")

	// Wait for DOM to be ready with retries
	var err error
	for i := 0; i < 10; i++ {
		err = waitForDOM()
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("Failed to initialize DOM: %v", err)
	}

	// Initialize test container
	doc := dom.Global()
	if doc == nil {
		t.Fatal("Failed to get global document")
	}

	// Create and append test container
	testContainer := doc.CreateElement("div")
	if testContainer == nil {
		t.Fatal("Failed to create test container")
	}
	testContainer.SetID("test-container")
	if err := doc.GetBody().AppendChild(&dom.Node{Value: testContainer.Value}); err != nil {
		t.Fatalf("Failed to append test container: %v", err)
	}

	// Verify test container was added
	if doc.GetElementByID("test-container") == nil {
		t.Fatal("Failed to add test container to document")
	}

	tests := []struct {
		name     string
		setup    func() error
		validate func() error
	}{
		{
			name: "Element Creation and Attributes",
			setup: func() error {
				div := doc.CreateElement("div")
				if div == nil {
					return fmt.Errorf("failed to create div element")
				}
				div.SetID("test-element")
				div.SetClassName("test-class")
				div.SetAttribute("data-test", "test-value")
				node := &dom.Node{Value: div.Value}
				if node.Value == nil {
					return fmt.Errorf("failed to create node from element")
				}
				if err := testContainer.AppendChild(node); err != nil {
					return fmt.Errorf("failed to append div: %v", err)
				}
				return nil
			},
			validate: func() error {
				foundDiv := doc.GetElementByID("test-element")
				if foundDiv == nil {
					return fmt.Errorf("failed to find div by ID")
				}
				if foundDiv.GetClassName() != "test-class" {
					return fmt.Errorf("failed to get class attribute")
				}
				if foundDiv.GetAttribute("data-test") != "test-value" {
					return fmt.Errorf("failed to get data-test attribute")
				}
				return nil
			},
		},
		{
			name: "Text Content",
			setup: func() error {
				div := doc.CreateElement("div")
				if div == nil {
					return fmt.Errorf("failed to create div element")
				}
				div.SetID("test-text-content-element")
				div.SetTextContent("Hello, World!")
				// Force a small delay to ensure DOM updates
				time.Sleep(50 * time.Millisecond)
				node := &dom.Node{Value: div.Value}
				if node.Value == nil {
					return fmt.Errorf("failed to create node from element")
				}
				if err := testContainer.AppendChild(node); err != nil {
					return fmt.Errorf("failed to append div: %v", err)
				}
				// Force another small delay to ensure DOM updates
				time.Sleep(50 * time.Millisecond)
				return nil
			},
			validate: func() error {
				div := doc.GetElementByID("test-text-content-element")
				if div == nil {
					return fmt.Errorf("failed to find div")
				}

				fmt.Printf("Text Content Test - Found div: %+v\n", div)
				fmt.Printf("Text Content Test - div.Value: %+v\n", div.Value)
				fmt.Printf("Text Content Test - div.Value.Raw(): %+v\n", div.Value.Raw())

				content := div.GetTextContent()
				if content != "Hello, World!" {
					return fmt.Errorf("failed to get text content, got: %s", content)
				}
				return nil
			},
		},
		{
			name: "Child Nodes",
			setup: func() error {
				parent := doc.CreateElement("div")
				if parent == nil {
					return fmt.Errorf("failed to create parent div")
				}
				parent.SetID("parent-element")
				child1 := doc.CreateElement("span")
				if child1 == nil {
					return fmt.Errorf("failed to create child span")
				}
				child1.SetTextContent("Child 1")
				child2 := doc.CreateElement("span")
				if child2 == nil {
					return fmt.Errorf("failed to create child span")
				}
				child2.SetTextContent("Child 2")
				if err := parent.AppendChild(&dom.Node{Value: child1.Value}); err != nil {
					return fmt.Errorf("failed to append child1: %v", err)
				}
				if err := parent.AppendChild(&dom.Node{Value: child2.Value}); err != nil {
					return fmt.Errorf("failed to append child2: %v", err)
				}
				if err := testContainer.AppendChild(&dom.Node{Value: parent.Value}); err != nil {
					return fmt.Errorf("failed to append parent: %v", err)
				}
				return nil
			},
			validate: func() error {
				parent := doc.GetElementByID("parent-element")
				if parent == nil {
					return fmt.Errorf("failed to find parent node")
				}
				children := parent.GetChildNodes()
				if len(children) != 2 {
					return fmt.Errorf("expected 2 children, got %d", len(children))
				}
				if children[0].GetTextContent() != "Child 1" {
					return fmt.Errorf("failed to get child1 text content")
				}
				if children[1].GetTextContent() != "Child 2" {
					return fmt.Errorf("failed to get child2 text content")
				}
				return nil
			},
		},
		{
			name: "Event Handling",
			setup: func() error {
				div := doc.CreateElement("div")
				if div == nil {
					return fmt.Errorf("failed to create div element")
				}
				div.SetID("event-test")
				div.AddEventListener("click", func(e *dom.Event) {
					fmt.Println("Element clicked!")
				})
				if err := testContainer.AppendChild(&dom.Node{Value: div.Value}); err != nil {
					return fmt.Errorf("failed to append div: %v", err)
				}
				return nil
			},
			validate: func() error {
				div := doc.GetElementByID("event-test")
				if div == nil {
					return fmt.Errorf("failed to find event test div")
				}
				return nil
			},
		},
		{
			name: "Style Manipulation",
			setup: func() error {
				div := doc.CreateElement("div")
				if div == nil {
					return fmt.Errorf("failed to create div element")
				}
				div.SetID("styled-element")
				style := div.GetStyle()
				style.SetWidth("100px")
				style.SetHeight("100px")
				style.SetBackgroundColor("red")
				style.SetColor("white")
				style.SetPadding("10px")
				style.SetBorder("1px solid black")
				if err := testContainer.AppendChild(&dom.Node{Value: div.Value}); err != nil {
					return fmt.Errorf("failed to append div: %v", err)
				}
				return nil
			},
			validate: func() error {
				div := doc.GetElementByID("styled-element")
				if div == nil {
					return fmt.Errorf("failed to find styled element")
				}
				style := div.GetStyle()
				if style.GetWidth() != "100px" {
					return fmt.Errorf("failed to get width")
				}
				if style.GetHeight() != "100px" {
					return fmt.Errorf("failed to get height")
				}
				if style.GetBackgroundColor() != "red" {
					return fmt.Errorf("failed to get background color")
				}
				if style.GetColor() != "white" {
					return fmt.Errorf("failed to get text color")
				}
				if style.GetPadding() != "10px" {
					return fmt.Errorf("failed to get padding")
				}
				if style.GetBorder() != "1px solid black" {
					return fmt.Errorf("failed to get border")
				}
				return nil
			},
		},
		{
			name: "Element Removal",
			setup: func() error {
				parent := doc.CreateElement("div")
				if parent == nil {
					return fmt.Errorf("failed to create parent div")
				}
				parent.SetID("parent-to-remove")
				child := doc.CreateElement("div")
				if child == nil {
					return fmt.Errorf("failed to create child div")
				}
				child.SetID("child-to-remove")
				child.SetTextContent("child-to-remove")
				childNode := &dom.Node{Value: child.Value}
				if childNode.Value == nil {
					return fmt.Errorf("failed to create child node")
				}
				if err := parent.AppendChild(childNode); err != nil {
					return fmt.Errorf("failed to append child: %v", err)
				}
				parentNode := &dom.Node{Value: parent.Value}
				if parentNode.Value == nil {
					return fmt.Errorf("failed to create parent node")
				}
				if err := testContainer.AppendChild(parentNode); err != nil {
					return fmt.Errorf("failed to append parent: %v", err)
				}
				return nil
			},
			validate: func() error {
				parent := doc.GetElementByID("parent-to-remove")
				if parent == nil {
					return fmt.Errorf("failed to find parent node")
				}
				children := parent.GetChildNodes()
				if len(children) != 1 {
					return fmt.Errorf("expected 1 child, got %d", len(children))
				}
				content := children[0].GetTextContent()
				if content != "child-to-remove" {
					return fmt.Errorf("failed to find child node, got content: %s", content)
				}
				return nil
			},
		},
		{
			name: "Element Replacement",
			setup: func() error {
				parent := doc.CreateElement("div")
				if parent == nil {
					return fmt.Errorf("failed to create parent div")
				}
				parent.SetID("parent-to-replace")

				// Create and append old child
				oldChild := doc.CreateElement("div")
				if oldChild == nil {
					return fmt.Errorf("failed to create old child div")
				}
				oldChild.SetID("old-child")
				oldChild.SetTextContent("Old Child")
				oldChildNode := &dom.Node{Value: oldChild.Value}
				if oldChildNode.Value == nil {
					return fmt.Errorf("failed to create old child node")
				}
				if err := parent.AppendChild(oldChildNode); err != nil {
					return fmt.Errorf("failed to append old child: %v", err)
				}

				// Append parent to test container
				parentNode := &dom.Node{Value: parent.Value}
				if parentNode.Value == nil {
					return fmt.Errorf("failed to create parent node")
				}
				if err := testContainer.AppendChild(parentNode); err != nil {
					return fmt.Errorf("failed to append parent: %v", err)
				}

				// Force a small delay to ensure DOM updates
				time.Sleep(50 * time.Millisecond)

				// Create new child
				newChild := doc.CreateElement("div")
				if newChild == nil {
					return fmt.Errorf("failed to create new child div")
				}
				newChild.SetID("new-child")
				newChild.SetTextContent("New Child")

				// Get fresh references to parent and old child
				parent = doc.GetElementByID("parent-to-replace")
				if parent == nil {
					return fmt.Errorf("failed to find parent after append")
				}
				oldChild = doc.GetElementByID("old-child")
				if oldChild == nil {
					return fmt.Errorf("failed to find old child after append")
				}

				// Ensure newChildNode and oldChildNode are based on fresh values
				freshNewChildNode := &dom.Node{Value: newChild.Value}
				if freshNewChildNode.Value == nil {
					return fmt.Errorf("failed to create fresh new child node")
				}
				freshOldChildNode := &dom.Node{Value: oldChild.Value}
				if freshOldChildNode.Value == nil {
					return fmt.Errorf("failed to create fresh old child node")
				}

				fmt.Printf("Before ReplaceChild - parent.Value: %+v\n", parent.Value)
				fmt.Printf("Before ReplaceChild - freshNewChildNode.Value: %+v\n", freshNewChildNode.Value)
				fmt.Printf("Before ReplaceChild - freshOldChildNode.Value: %+v\n", freshOldChildNode.Value)

				// Perform replacement
				if err := parent.ReplaceChild(freshNewChildNode, freshOldChildNode); err != nil {
					return fmt.Errorf("failed to replace child: %v", err)
				}

				// Force a small delay to ensure DOM updates
				time.Sleep(50 * time.Millisecond)
				return nil
			},
			validate: func() error {
				parent := doc.GetElementByID("parent-to-replace")
				if parent == nil {
					return fmt.Errorf("failed to find parent node")
				}
				children := parent.GetChildNodes()
				if len(children) != 1 {
					return fmt.Errorf("expected 1 child, got %d", len(children))
				}
				content := children[0].GetTextContent()
				if content != "New Child" {
					return fmt.Errorf("failed to get new child text content, got: %s", content)
				}
				return nil
			},
		},
		{
			name: "Element Cloning",
			setup: func() error {
				original := doc.CreateElement("div")
				if original == nil {
					return fmt.Errorf("failed to create original div")
				}
				original.SetID("original")
				original.SetClassName("original-class")
				original.SetTextContent("Original Content")
				if err := testContainer.AppendChild(&dom.Node{Value: original.Value}); err != nil {
					return fmt.Errorf("failed to append original: %v", err)
				}
				return nil
			},
			validate: func() error {
				original := doc.GetElementByID("original")
				if original == nil {
					return fmt.Errorf("failed to find original div")
				}
				clone := original.CloneNode(true)
				if clone == nil {
					return fmt.Errorf("failed to clone node")
				}
				if clone.GetID() != "original" {
					return fmt.Errorf("failed to get cloned ID")
				}
				if clone.GetClassName() != "original-class" {
					return fmt.Errorf("failed to get cloned class")
				}
				if clone.GetTextContent() != "Original Content" {
					return fmt.Errorf("failed to get cloned text content")
				}
				return nil
			},
		},
		{
			name: "Element Querying",
			setup: func() error {
				container := doc.CreateElement("div")
				if container == nil {
					return fmt.Errorf("failed to create container div")
				}
				container.SetID("query-container")

				// Append container to testContainer
				if err := testContainer.AppendChild(&dom.Node{Value: container.Value}); err != nil {
					return fmt.Errorf("failed to append query container: %v", err)
				}

				element1 := doc.CreateElement("div")
				if element1 == nil {
					return fmt.Errorf("failed to create element1 div")
				}
				element1.SetClassName("test-class")
				element1.SetID("element1")
				if err := container.AppendChild(&dom.Node{Value: element1.Value}); err != nil {
					return fmt.Errorf("failed to append element1: %v", err)
				}
				element2 := doc.CreateElement("div")
				if element2 == nil {
					return fmt.Errorf("failed to create element2 div")
				}
				element2.SetClassName("test-class")
				element2.SetID("element2")
				if err := container.AppendChild(&dom.Node{Value: element2.Value}); err != nil {
					return fmt.Errorf("failed to append element2: %v", err)
				}
				return nil
			},
			validate: func() error {
				container := doc.GetElementByID("query-container")
				if container == nil {
					return fmt.Errorf("failed to find query container by ID")
				}
				children := container.GetChildNodes()
				if len(children) != 2 {
					return fmt.Errorf("expected 2 children, got %d", len(children))
				}

				// Convert children to elements for ID/ClassName check (if they are elements)
				elem1 := &dom.Element{Value: children[0].Value}
				if elem1.Value == nil || elem1.Value.Raw().IsNull() || elem1.Value.Raw().IsUndefined() {
					return fmt.Errorf("child 0 value is nil or undefined/null")
				}
				elem2 := &dom.Element{Value: children[1].Value}
				if elem2.Value == nil || elem2.Value.Raw().IsNull() || elem2.Value.Raw().IsUndefined() {
					return fmt.Errorf("child 1 value is nil or undefined/null")
				}

				if elem1.GetID() != "element1" {
					return fmt.Errorf("failed to find element1, got ID: %s", elem1.GetID())
				}
				if elem2.GetID() != "element2" {
					return fmt.Errorf("failed to find element2, got ID: %s", elem2.GetID())
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.name)
			if err := tt.setup(); err != nil {
				t.Errorf("setup failed: %v", err)
				fmt.Printf("❌ Test setup failed: %s - %v\n", tt.name, err)
				return
			}
			if err := tt.validate(); err != nil {
				t.Errorf("validation failed: %v", err)
				fmt.Printf("❌ Test failed: %s - %v\n", tt.name, err)
			} else {
				fmt.Printf("✅ Test passed: %s\n", tt.name)
			}
		})
	}

	fmt.Println("\nTest summary:")
	fmt.Printf("Total tests: %d\n", len(tests))
	fmt.Println("Tests completed!")
}
