//go:build js && wasm
// +build js,wasm

package js

import (
	"fmt"
	"syscall/js"
	"testing"
)

func TestValueWrapper(t *testing.T) {
	fmt.Println("Starting WebAssembly tests...")

	tests := []struct {
		name     string
		setup    func() *Value
		validate func(*Value) error
	}{
		{
			name: "Global object access",
			setup: func() *Value {
				return Global()
			},
			validate: func(v *Value) error {
				if v.IsNull() || v.IsUndefined() {
					return fmt.Errorf("global object should not be null or undefined")
				}
				return nil
			},
		},
		{
			name: "Object property access and type checking",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("string", "value")
				obj.Set("number", 42)
				obj.Set("boolean", true)
				obj.Set("null", nil)
				obj.Set("undefined", js.Undefined())
				return obj
			},
			validate: func(v *Value) error {
				// Test string property
				if val, err := v.Get("string").String(); err != nil || val != "value" {
					return fmt.Errorf("string property error: %v", err)
				}

				// Test number property
				if val, err := v.Get("number").Int(); err != nil || val != 42 {
					return fmt.Errorf("number property error: %v", err)
				}

				// Test boolean property
				if val, err := v.Get("boolean").Bool(); err != nil || !val {
					return fmt.Errorf("boolean property error: %v", err)
				}

				// Test null property
				if !v.Get("null").IsNull() {
					return fmt.Errorf("null property should be null")
				}

				// Test undefined property
				if !v.Get("undefined").IsUndefined() {
					return fmt.Errorf("undefined property should be undefined")
				}

				return nil
			},
		},
		{
			name: "Array operations and type checking",
			setup: func() *Value {
				arr := Global().Get("Array").Call("from", []interface{}{
					"string", 42, true, nil, js.Undefined(),
					map[string]interface{}{"key": "value"},
					[]interface{}{1, 2, 3},
				})
				return arr
			},
			validate: func(v *Value) error {
				// First check if it's an array
				isArray, err := Global().Get("Array").Call("isArray", v.value).Bool()
				if err != nil {
					return fmt.Errorf("error checking if value is array: %v", err)
				}
				if !isArray {
					return fmt.Errorf("value is not an array")
				}

				// Get array length
				length := v.value.Length()
				if length != 7 {
					return fmt.Errorf("expected array length 7, got %d", length)
				}

				// Test each element type
				if str, err := New(v.value.Index(0)).String(); err != nil || str != "string" {
					return fmt.Errorf("array[0] string error: %v", err)
				}

				if num, err := New(v.value.Index(1)).Int(); err != nil || num != 42 {
					return fmt.Errorf("array[1] number error: %v", err)
				}

				if bool, err := New(v.value.Index(2)).Bool(); err != nil || !bool {
					return fmt.Errorf("array[2] boolean error: %v", err)
				}

				if !New(v.value.Index(3)).IsNull() {
					return fmt.Errorf("array[3] should be null")
				}

				if !New(v.value.Index(4)).IsUndefined() {
					return fmt.Errorf("array[4] should be undefined")
				}

				// Test nested object
				obj := New(v.value.Index(5))
				if val, err := obj.Get("key").String(); err != nil || val != "value" {
					return fmt.Errorf("array[5] object error: key should be 'value', got error: %v", err)
				}

				// Test nested array
				nestedArr := New(v.value.Index(6))
				if length, err := nestedArr.Length(); err != nil || length != 3 {
					return fmt.Errorf("array[6] nested array error: expected length 3, got %d, error: %v", length, err)
				}

				// Verify nested array elements
				for i := 0; i < 3; i++ {
					if val, err := New(nestedArr.value.Index(i)).Int(); err != nil || val != i+1 {
						return fmt.Errorf("array[6][%d] should be %d, got %d, error: %v", i, i+1, val, err)
					}
				}

				return nil
			},
		},
		{
			name: "Error handling for type conversions",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("string", "not a number")
				obj.Set("number", 42)
				obj.Set("boolean", true)
				return obj
			},
			validate: func(v *Value) error {
				// Test invalid number conversion
				if _, err := v.Get("string").Int(); err == nil {
					return fmt.Errorf("expected error when converting string to number")
				}

				// Test invalid boolean conversion
				if _, err := v.Get("number").Bool(); err == nil {
					return fmt.Errorf("expected error when converting number to boolean")
				}

				// Test invalid string conversion
				if _, err := v.Get("boolean").String(); err == nil {
					return fmt.Errorf("expected error when converting boolean to string")
				}

				return nil
			},
		},
		{
			name: "JSON conversion with complex objects",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("string", "test")
				obj.Set("number", 42)
				obj.Set("boolean", true)
				obj.Set("null", nil)
				obj.Set("array", []interface{}{1, "two", true, nil})
				obj.Set("nested", map[string]interface{}{
					"key":    "value",
					"number": 123,
					"array":  []interface{}{"nested", "array"},
				})
				return obj
			},
			validate: func(v *Value) error {
				result, err := v.ToJSON()
				if err != nil {
					return fmt.Errorf("JSON conversion error: %v", err)
				}

				m, ok := result.(map[string]interface{})
				if !ok {
					return fmt.Errorf("expected map result, got %T", result)
				}

				// Verify all properties
				if m["string"] != "test" {
					return fmt.Errorf("string property mismatch")
				}
				if m["number"] != float64(42) {
					return fmt.Errorf("number property mismatch")
				}
				if m["boolean"] != true {
					return fmt.Errorf("boolean property mismatch")
				}
				if m["null"] != nil {
					return fmt.Errorf("null property mismatch")
				}

				// Verify array
				arr, ok := m["array"].([]interface{})
				if !ok || len(arr) != 4 {
					return fmt.Errorf("array property mismatch")
				}

				// Verify nested object
				nested, ok := m["nested"].(map[string]interface{})
				if !ok {
					return fmt.Errorf("nested object type mismatch")
				}
				if nested["key"] != "value" {
					return fmt.Errorf("nested key property mismatch")
				}

				return nil
			},
		},
		{
			name: "Callback handling with different argument types",
			setup: func() *Value {
				obj := Global().Call("Object")
				var callbackCalled bool
				callback := NewCallback(func(args []*Value) {
					callbackCalled = true
					if len(args) != 5 {
						t.Errorf("expected 5 arguments, got %d", len(args))
					}
					// Test string argument
					if str, err := args[0].String(); err != nil || str != "test" {
						t.Errorf("string argument error: %v", err)
					}
					// Test number argument
					if num, err := args[1].Int(); err != nil || num != 42 {
						t.Errorf("number argument error: %v", err)
					}
					// Test boolean argument
					if bool, err := args[2].Bool(); err != nil || !bool {
						t.Errorf("boolean argument error: %v", err)
					}
					// Test null argument
					if !args[3].IsNull() {
						t.Errorf("null argument error")
					}
					// Test object argument
					objArg := args[4]
					if objArg.Type() != js.TypeObject {
						t.Errorf("object argument error: expected object type, got %v", objArg.Type())
					}
					if val, err := objArg.Get("key").String(); err != nil || val != "value" {
						t.Errorf("object argument error: expected key 'value', got error: %v, value: %v", err, val)
					}
				})
				obj.Set("callback", callback)
				obj.Call("callback", "test", 42, true, nil, map[string]interface{}{"key": "value"})
				obj.Set("callbackCalled", callbackCalled)
				return obj
			},
			validate: func(v *Value) error {
				called := v.Get("callbackCalled").MustBool()
				if !called {
					return fmt.Errorf("callback was not called")
				}
				return nil
			},
		},
		{
			name: "Must and Try methods with error cases",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("validString", "test")
				obj.Set("invalidString", 42)
				obj.Set("validNumber", 42)
				obj.Set("invalidNumber", "not a number")
				obj.Set("validBoolean", true)
				obj.Set("invalidBoolean", "not a boolean")
				return obj
			},
			validate: func(v *Value) error {
				// Test Must methods with valid values
				if v.Get("validString").MustString() != "test" {
					return fmt.Errorf("MustString failed with valid value")
				}
				if v.Get("validNumber").MustInt() != 42 {
					return fmt.Errorf("MustInt failed with valid value")
				}
				if !v.Get("validBoolean").MustBool() {
					return fmt.Errorf("MustBool failed with valid value")
				}

				// Test Try methods with invalid values
				if v.Get("invalidString").TryString("fallback") != "fallback" {
					return fmt.Errorf("TryString failed with invalid value")
				}
				if v.Get("invalidNumber").TryInt(0) != 0 {
					return fmt.Errorf("TryInt failed with invalid value")
				}
				if v.Get("invalidBoolean").TryBool(false) != false {
					return fmt.Errorf("TryBool failed with invalid value")
				}

				// Test Must methods with invalid values (should panic)
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("MustString should have panicked with invalid value")
					}
				}()
				v.Get("invalidString").MustString()

				return nil
			},
		},
		{
			name: "Object property existence and type checking",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("exists", "value")
				return obj
			},
			validate: func(v *Value) error {
				if !v.Exists("exists") {
					return fmt.Errorf("property 'exists' should exist")
				}
				if v.Exists("nonexistent") {
					return fmt.Errorf("property 'nonexistent' should not exist")
				}
				if v.Get("nonexistent").Type() != js.TypeUndefined {
					return fmt.Errorf("nonexistent property should be undefined")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", tt.name)
			value := tt.setup()
			if err := tt.validate(value); err != nil {
				t.Errorf("test failed: %v", err)
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
