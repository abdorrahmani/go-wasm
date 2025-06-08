//go:build js && wasm
// +build js,wasm

package dom

import (
	"github.com/abdorrahmani/go-wasm/js"
)

// Style represents the CSS style of an element
type Style struct {
	Value *js.Value
}

// SetProperty sets a CSS property
func (s *Style) SetProperty(name, value string) {
	s.Value.Call("setProperty", name, value)
}

// GetPropertyValue gets the value of a CSS property
func (s *Style) GetPropertyValue(name string) string {
	return s.Value.Call("getPropertyValue", name).MustString()
}

// RemoveProperty removes a CSS property
func (s *Style) RemoveProperty(name string) {
	s.Value.Call("removeProperty", name)
}

// GetPropertyPriority gets the priority of a CSS property
func (s *Style) GetPropertyPriority(name string) string {
	return s.Value.Call("getPropertyPriority", name).MustString()
}

// SetPropertyWithPriority sets a CSS property with priority
func (s *Style) SetPropertyWithPriority(name, value, priority string) {
	s.Value.Call("setProperty", name, value, priority)
}

// GetCSSText gets all CSS properties as a string
func (s *Style) GetCSSText() string {
	return s.Value.Get("cssText").MustString()
}

// SetCSSText sets all CSS properties from a string
func (s *Style) SetCSSText(text string) {
	s.Value.Set("cssText", text)
}

// GetLength returns the number of CSS properties
func (s *Style) GetLength() int {
	return s.Value.Get("length").MustInt()
}

// GetItem returns the name of a CSS property by index
func (s *Style) GetItem(index int) string {
	return s.Value.Call("item", index).MustString()
}

// GetParentRule returns the parent CSS rule
func (s *Style) GetParentRule() *CSSRule {
	return &CSSRule{
		Value: s.Value.Get("parentRule"),
	}
}

// Common style properties
func (s *Style) SetColor(value string) {
	s.SetProperty("color", value)
}

func (s *Style) GetColor() string {
	return s.GetPropertyValue("color")
}

func (s *Style) SetBackgroundColor(value string) {
	s.SetProperty("background-color", value)
}

func (s *Style) GetBackgroundColor() string {
	return s.GetPropertyValue("background-color")
}

func (s *Style) SetWidth(value string) {
	s.SetProperty("width", value)
}

func (s *Style) GetWidth() string {
	return s.GetPropertyValue("width")
}

func (s *Style) SetHeight(value string) {
	s.SetProperty("height", value)
}

func (s *Style) GetHeight() string {
	return s.GetPropertyValue("height")
}

func (s *Style) SetMargin(value string) {
	s.SetProperty("margin", value)
}

func (s *Style) GetMargin() string {
	return s.GetPropertyValue("margin")
}

func (s *Style) SetPadding(value string) {
	s.SetProperty("padding", value)
}

func (s *Style) GetPadding() string {
	return s.GetPropertyValue("padding")
}

func (s *Style) SetBorder(value string) {
	s.SetProperty("border", value)
}

func (s *Style) GetBorder() string {
	return s.GetPropertyValue("border")
}

func (s *Style) SetDisplay(value string) {
	s.SetProperty("display", value)
}

func (s *Style) GetDisplay() string {
	return s.GetPropertyValue("display")
}

func (s *Style) SetPosition(value string) {
	s.SetProperty("position", value)
}

func (s *Style) GetPosition() string {
	return s.GetPropertyValue("position")
}

func (s *Style) SetTop(value string) {
	s.SetProperty("top", value)
}

func (s *Style) GetTop() string {
	return s.GetPropertyValue("top")
}

func (s *Style) SetRight(value string) {
	s.SetProperty("right", value)
}

func (s *Style) GetRight() string {
	return s.GetPropertyValue("right")
}

func (s *Style) SetBottom(value string) {
	s.SetProperty("bottom", value)
}

func (s *Style) GetBottom() string {
	return s.GetPropertyValue("bottom")
}

func (s *Style) SetLeft(value string) {
	s.SetProperty("left", value)
}

func (s *Style) GetLeft() string {
	return s.GetPropertyValue("left")
}

// CSSRule represents a CSS rule
type CSSRule struct {
	Value *js.Value
}

// GetType returns the type of the CSS rule
func (r *CSSRule) GetType() int {
	return r.Value.Get("type").MustInt()
}

// GetCSSText gets the CSS rule as a string
func (r *CSSRule) GetCSSText() string {
	return r.Value.Get("cssText").MustString()
}

// SetCSSText sets the CSS rule from a string
func (r *CSSRule) SetCSSText(text string) {
	r.Value.Set("cssText", text)
}

// GetParentStyleSheet returns the parent style sheet
func (r *CSSRule) GetParentStyleSheet() *StyleSheet {
	return &StyleSheet{
		Value: r.Value.Get("parentStyleSheet"),
	}
}

// StyleSheet represents a CSS style sheet
type StyleSheet struct {
	Value *js.Value
}

// GetType returns the type of the style sheet
func (s *StyleSheet) GetType() string {
	return s.Value.Get("type").MustString()
}

// GetHref returns the URL of the style sheet
func (s *StyleSheet) GetHref() string {
	return s.Value.Get("href").MustString()
}

// GetOwnerNode returns the node that owns the style sheet
func (s *StyleSheet) GetOwnerNode() *Node {
	return &Node{
		Value: s.Value.Get("ownerNode"),
	}
}

// GetParentStyleSheet returns the parent style sheet
func (s *StyleSheet) GetParentStyleSheet() *StyleSheet {
	return &StyleSheet{
		Value: s.Value.Get("parentStyleSheet"),
	}
}

// GetTitle returns the title of the style sheet
func (s *StyleSheet) GetTitle() string {
	return s.Value.Get("title").MustString()
}

// GetMedia returns the media list of the style sheet
func (s *StyleSheet) GetMedia() *MediaList {
	return &MediaList{
		Value: s.Value.Get("media"),
	}
}

// GetDisabled returns true if the style sheet is disabled
func (s *StyleSheet) GetDisabled() bool {
	return s.Value.Get("disabled").MustBool()
}

// SetDisabled sets whether the style sheet is disabled
func (s *StyleSheet) SetDisabled(disabled bool) {
	s.Value.Set("disabled", disabled)
}

// MediaList represents a list of media queries
type MediaList struct {
	Value *js.Value
}

// GetLength returns the number of media queries
func (m *MediaList) GetLength() int {
	return m.Value.Get("length").MustInt()
}

// GetItem returns a media query by index
func (m *MediaList) GetItem(index int) string {
	return m.Value.Call("item", index).MustString()
}

// GetMediaText gets all media queries as a string
func (m *MediaList) GetMediaText() string {
	return m.Value.Get("mediaText").MustString()
}

// SetMediaText sets all media queries from a string
func (m *MediaList) SetMediaText(text string) {
	m.Value.Set("mediaText", text)
}

// AppendMedium adds a media query
func (m *MediaList) AppendMedium(medium string) {
	m.Value.Call("appendMedium", medium)
}

// DeleteMedium removes a media query
func (m *MediaList) DeleteMedium(medium string) {
	m.Value.Call("deleteMedium", medium)
}

// SetBorderRadius sets the border-radius property
func (s *Style) SetBorderRadius(value string) {
	s.SetProperty("border-radius", value)
}

// SetCursor sets the cursor property
func (s *Style) SetCursor(value string) {
	s.SetProperty("cursor", value)
}

// SetFontSize sets the font-size property
func (s *Style) SetFontSize(value string) {
	s.SetProperty("font-size", value)
}
