//go:build js && wasm
// +build js,wasm

package dom

import (
	"github.com/abdorrahmani/go-wasm/js"
)

// Event represents a DOM event
type Event struct {
	Value *js.Value
}

// EventPhase constants
const (
	EventPhaseNone      = 0
	EventPhaseCapturing = 1
	EventPhaseAtTarget  = 2
	EventPhaseBubbling  = 3
)

// GetType returns the type of the event
func (e *Event) GetType() string {
	return e.Value.Get("type").MustString()
}

// GetTarget returns the target element of the event
func (e *Event) GetTarget() *Element {
	return &Element{
		Value: e.Value.Get("target"),
	}
}

// GetCurrentTarget returns the current target element of the event
func (e *Event) GetCurrentTarget() *Element {
	return &Element{
		Value: e.Value.Get("currentTarget"),
	}
}

// GetEventPhase returns the current phase of the event
func (e *Event) GetEventPhase() int {
	return e.Value.Get("eventPhase").MustInt()
}

// GetBubbles returns whether the event bubbles
func (e *Event) GetBubbles() bool {
	return e.Value.Get("bubbles").MustBool()
}

// GetCancelable returns whether the event is cancelable
func (e *Event) GetCancelable() bool {
	return e.Value.Get("cancelable").MustBool()
}

// GetTimeStamp returns the time when the event was created
func (e *Event) GetTimeStamp() float64 {
	return e.Value.Get("timeStamp").MustFloat()
}

// StopPropagation stops the event from propagating
func (e *Event) StopPropagation() {
	e.Value.Call("stopPropagation")
}

// StopImmediatePropagation stops the event from propagating and prevents other handlers from being called
func (e *Event) StopImmediatePropagation() {
	e.Value.Call("stopImmediatePropagation")
}

// PreventDefault prevents the default action of the event
func (e *Event) PreventDefault() {
	e.Value.Call("preventDefault")
}

// GetDefaultPrevented returns true if preventDefault was called
func (e *Event) GetDefaultPrevented() bool {
	return e.Value.Get("defaultPrevented").MustBool()
}

// GetIsTrusted returns whether the event is trusted
func (e *Event) GetIsTrusted() bool {
	return e.Value.Get("isTrusted").MustBool()
}

// MouseEvent represents a mouse event
type MouseEvent struct {
	Event
}

// GetButton returns the button that was pressed
func (e *MouseEvent) GetButton() int {
	return e.Value.Get("button").MustInt()
}

// GetButtons returns the buttons that are pressed
func (e *MouseEvent) GetButtons() int {
	return e.Value.Get("buttons").MustInt()
}

// GetClientX returns the X coordinate relative to the viewport
func (e *MouseEvent) GetClientX() float64 {
	return e.Value.Get("clientX").MustFloat()
}

// GetClientY returns the Y coordinate relative to the viewport
func (e *MouseEvent) GetClientY() float64 {
	return e.Value.Get("clientY").MustFloat()
}

// GetScreenX returns the X coordinate relative to the screen
func (e *MouseEvent) GetScreenX() float64 {
	return e.Value.Get("screenX").MustFloat()
}

// GetScreenY returns the Y coordinate relative to the screen
func (e *MouseEvent) GetScreenY() float64 {
	return e.Value.Get("screenY").MustFloat()
}

// GetMovementX returns the X coordinate of the mouse movement
func (e *MouseEvent) GetMovementX() float64 {
	return e.Value.Get("movementX").MustFloat()
}

// GetMovementY returns the Y coordinate of the mouse movement
func (e *MouseEvent) GetMovementY() float64 {
	return e.Value.Get("movementY").MustFloat()
}

// GetOffsetX returns the offset X coordinate
func (e *MouseEvent) GetOffsetX() float64 {
	return e.Value.Get("offsetX").MustFloat()
}

// GetOffsetY returns the offset Y coordinate
func (e *MouseEvent) GetOffsetY() float64 {
	return e.Value.Get("offsetY").MustFloat()
}

// GetPageX returns the page X coordinate
func (e *MouseEvent) GetPageX() float64 {
	return e.Value.Get("pageX").MustFloat()
}

// GetPageY returns the page Y coordinate
func (e *MouseEvent) GetPageY() float64 {
	return e.Value.Get("pageY").MustFloat()
}

// GetX returns the X coordinate
func (e *MouseEvent) GetX() float64 {
	return e.Value.Get("x").MustFloat()
}

// GetY returns the Y coordinate
func (e *MouseEvent) GetY() float64 {
	return e.Value.Get("y").MustFloat()
}

// GetAltKey returns true if the Alt key was pressed
func (e *MouseEvent) GetAltKey() bool {
	return e.Value.Get("altKey").MustBool()
}

// GetCtrlKey returns true if the Ctrl key was pressed
func (e *MouseEvent) GetCtrlKey() bool {
	return e.Value.Get("ctrlKey").MustBool()
}

// GetMetaKey returns true if the Meta key was pressed
func (e *MouseEvent) GetMetaKey() bool {
	return e.Value.Get("metaKey").MustBool()
}

// GetShiftKey returns true if the Shift key was pressed
func (e *MouseEvent) GetShiftKey() bool {
	return e.Value.Get("shiftKey").MustBool()
}

// GetRelatedTarget returns the related target
func (e *MouseEvent) GetRelatedTarget() *Element {
	return &Element{
		Value: e.Value.Get("relatedTarget"),
	}
}
