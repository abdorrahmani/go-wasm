//go:build js && wasm
// +build js,wasm

package js

import (
	"fmt"
	"syscall/js"
)

// Value represents a JavaScript value with type-safe methods
type Value struct {
	value js.Value
}

// Global returns the JavaScript global object
func Global() *Value {
	return &Value{value: js.Global()}
}

// New creates a new Value wrapper from a js.Value
func New(v js.Value) *Value {
	return &Value{value: v}
}

// Get returns a property of the JavaScript value
func (v *Value) Get(name string) *Value {
	return &Value{value: v.value.Get(name)}
}

// Set sets a property of the JavaScript value
func (v *Value) Set(name string, value interface{}) {
	v.value.Set(name, value)
}

// Call calls a method of the JavaScript value
func (v *Value) Call(method string, args ...interface{}) *Value {
	return &Value{value: v.value.Call(method, args...)}
}

// Type returns the JavaScript type of the value
func (v *Value) Type() js.Type {
	return v.value.Type()
}

// IsNull checks if the value is null
func (v *Value) IsNull() bool {
	return v.value.IsNull()
}

// IsUndefined checks if the value is undefined
func (v *Value) IsUndefined() bool {
	return v.value.IsUndefined()
}

// Bool returns the value as a bool
func (v *Value) Bool() (bool, error) {
	if v.Type() != js.TypeBoolean {
		return false, fmt.Errorf("value is not a boolean")
	}
	return v.value.Bool(), nil
}

// Int returns the value as an int
func (v *Value) Int() (int, error) {
	if v.Type() != js.TypeNumber {
		return 0, fmt.Errorf("value is not a number")
	}
	return v.value.Int(), nil
}

// Float returns the value as a float64
func (v *Value) Float() (float64, error) {
	if v.Type() != js.TypeNumber {
		return 0, fmt.Errorf("value is not a number")
	}
	return v.value.Float(), nil
}

// String returns the value as a string
func (v *Value) String() (string, error) {
	if v.Type() != js.TypeString {
		return "", fmt.Errorf("value is not a string")
	}
	return v.value.String(), nil
}

// Object returns the value as a map[string]interface{}
func (v *Value) Object() (map[string]interface{}, error) {
	if v.Type() != js.TypeObject {
		return nil, fmt.Errorf("value is not an object")
	}

	obj := make(map[string]interface{})
	keys := Global().Get("Object").Call("keys", v.value)

	length, err := keys.Length()
	if err != nil {
		return nil, fmt.Errorf("error getting keys length: %v", err)
	}

	for i := 0; i < length; i++ {
		key, err := keys.Get(fmt.Sprintf("%d", i)).String()
		if err != nil {
			return nil, fmt.Errorf("error getting key at index %d: %v", i, err)
		}
		obj[key] = v.Get(key).value
	}

	return obj, nil
}

// Array returns the value as a slice of Values
func (v *Value) Array() ([]*Value, error) {
	if v.Type() != js.TypeObject || !v.value.Get("Array").Call("isArray", v.value).Bool() {
		return nil, fmt.Errorf("value is not an array")
	}

	length := v.value.Length()
	arr := make([]*Value, length)

	for i := 0; i < length; i++ {
		arr[i] = &Value{value: v.value.Index(i)}
	}

	return arr, nil
}

// Raw returns the underlying js.Value
func (v *Value) Raw() js.Value {
	return v.value
}

// Exists checks if a property exists on the JavaScript value
func (v *Value) Exists(key string) bool {
	return !v.Get(key).IsUndefined()
}

// Length returns the length of array-like values
func (v *Value) Length() (int, error) {
	if v.Type() != js.TypeObject {
		return 0, fmt.Errorf("value is not an object")
	}
	return v.value.Length(), nil
}

// ToJSON converts the value to a Go interface{} using JSON
func (v *Value) ToJSON() (interface{}, error) {
	if v.Type() != js.TypeObject {
		return nil, fmt.Errorf("value is not an object")
	}

	// Use JSON.stringify and JSON.parse to convert to Go value
	jsonStr, err := Global().Get("JSON").Call("stringify", v.value).String()
	if err != nil {
		return nil, fmt.Errorf("error stringifying value: %v", err)
	}
	var result interface{}
	err = Global().Get("JSON").Call("parse", jsonStr).Unmarshal(&result)
	return result, err
}

// Unmarshal converts the value to the provided Go type
func (v *Value) Unmarshal(target interface{}) error {
	if v.Type() != js.TypeObject {
		return fmt.Errorf("value is not an object")
	}

	jsonStr, err := Global().Get("JSON").Call("stringify", v.value).String()
	if err != nil {
		return fmt.Errorf("error stringifying value: %v", err)
	}
	return Global().Get("JSON").Call("parse", jsonStr).Unmarshal(target)
}
