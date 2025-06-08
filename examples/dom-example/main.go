//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"time"

	"github.com/abdorrahmani/go-wasm/dom"
)

func main() {
	// Get the global document
	doc := dom.Global()
	if doc == nil {
		fmt.Println("Failed to get global document")
		return
	}

	// Wait for DOM to be ready
	for i := 0; i < 10; i++ {
		if doc.IsReady() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Get elements
	output := doc.GetElementByID("output")
	controls := doc.GetElementByID("controls")
	eventOutput := doc.GetElementByID("eventOutput")

	if output == nil || controls == nil || eventOutput == nil {
		fmt.Println("Failed to get required elements")
		return
	}

	// Add event listeners to buttons
	addElementBtn := doc.GetElementByID("addElement")
	changeStyleBtn := doc.GetElementByID("changeStyle")
	addEventBtn := doc.GetElementByID("addEvent")

	// Add Element button click handler
	addElementBtn.AddEventListener("click", func(e *dom.Event) {
		// Create new element
		newDiv := doc.CreateElement("div")
		newDiv.SetClassName("box")
		newDiv.SetTextContent(fmt.Sprintf("New element created at %s", time.Now().Format("15:04:05")))

		// Append to output
		output.AppendChild(&dom.Node{Value: newDiv.Value})
	})

	// Change Style button click handler
	changeStyleBtn.AddEventListener("click", func(e *dom.Event) {
		// Toggle highlight class
		if output.GetClassName() == "box highlight" {
			output.SetClassName("box")
		} else {
			output.SetClassName("box highlight")
		}

		// Change some inline styles
		style := output.GetStyle()
		style.SetBorder("2px solid #2196f3")
		style.SetPadding("15px")
		style.SetBorderRadius("5px")
	})

	// Add Event button click handler
	addEventBtn.AddEventListener("click", func(e *dom.Event) {
		// Create a new element with event
		newElement := doc.CreateElement("div")
		newElement.SetClassName("box")
		newElement.SetTextContent("Click me to see event details")

		// Add click event to the new element
		newElement.AddEventListener("click", func(e *dom.Event) {
			// Get event details
			target := e.GetTarget()
			eventType := e.GetType()
			timestamp := e.GetTimeStamp()

			// Create event details element
			details := doc.CreateElement("p")
			details.SetTextContent(fmt.Sprintf(
				"Event: %s\nTarget: %s\nTime: %f",
				eventType,
				target.GetTagName(),
				timestamp,
			))

			// Append to event output
			eventOutput.AppendChild(&dom.Node{Value: details.Value})
		})

		// Append to output
		output.AppendChild(&dom.Node{Value: newElement.Value})
	})

	// Keep the program running
	select {}
}
