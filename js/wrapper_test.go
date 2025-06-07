//go:build js && wasm
// +build js,wasm

package js

import (
	"fmt"
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
			name: "Object property access",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("test", "value")
				return obj
			},
			validate: func(v *Value) error {
				val, err := v.Get("test").String()
				if err != nil {
					return err
				}
				if val != "value" {
					return fmt.Errorf("expected 'value', got '%s'", val)
				}
				return nil
			},
		},
		{
			name: "Array operations",
			setup: func() *Value {
				// Create array using Array constructor
				arr := Global().Get("Array").Call("from", []interface{}{1, 2, 3})
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
				if length != 3 {
					return fmt.Errorf("expected array length 3, got %d", length)
				}

				// Check each element
				for i := 0; i < length; i++ {
					val := v.value.Index(i).Int()
					if val != i+1 {
						return fmt.Errorf("expected value %d at index %d, got %d", i+1, i, val)
					}
				}

				return nil
			},
		},
		{
			name: "Property existence check",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("test", "value")
				return obj
			},
			validate: func(v *Value) error {
				if !v.Exists("test") {
					return fmt.Errorf("property 'test' should exist")
				}
				if v.Exists("nonexistent") {
					return fmt.Errorf("property 'nonexistent' should not exist")
				}
				return nil
			},
		},
		{
			name: "Length method",
			setup: func() *Value {
				return Global().Get("Array").Call("from", []interface{}{1, 2, 3, 4, 5})
			},
			validate: func(v *Value) error {
				length, err := v.Length()
				if err != nil {
					return fmt.Errorf("error getting length: %v", err)
				}
				if length != 5 {
					return fmt.Errorf("expected length 5, got %d", length)
				}
				return nil
			},
		},
		{
			name: "JSON conversion",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("name", "test")
				obj.Set("value", 42)
				obj.Set("nested", map[string]interface{}{
					"key": "value",
				})
				return obj
			},
			validate: func(v *Value) error {
				result, err := v.ToJSON()
				if err != nil {
					return fmt.Errorf("error converting to JSON: %v", err)
				}

				// Type assert to map
				m, ok := result.(map[string]interface{})
				if !ok {
					return fmt.Errorf("expected map result, got %T", result)
				}

				// Check values
				if m["name"] != "test" {
					return fmt.Errorf("expected name 'test', got %v", m["name"])
				}
				if m["value"] != float64(42) {
					return fmt.Errorf("expected value 42, got %v", m["value"])
				}

				nested, ok := m["nested"].(map[string]interface{})
				if !ok {
					return fmt.Errorf("expected nested map, got %T", m["nested"])
				}
				if nested["key"] != "value" {
					return fmt.Errorf("expected nested key 'value', got %v", nested["key"])
				}

				return nil
			},
		},
		{
			name: "Must and Try methods",
			setup: func() *Value {
				obj := Global().Call("Object")
				obj.Set("string", "test")
				obj.Set("number", 42)
				obj.Set("boolean", true)
				obj.Set("array", []interface{}{1, 2, 3})
				return obj
			},
			validate: func(v *Value) error {
				// Test Must methods
				if v.Get("string").MustString() != "test" {
					return fmt.Errorf("MustString failed")
				}
				if v.Get("number").MustInt() != 42 {
					return fmt.Errorf("MustInt failed")
				}
				if !v.Get("boolean").MustBool() {
					return fmt.Errorf("MustBool failed")
				}
				if v.Get("array").MustLength() != 3 {
					return fmt.Errorf("MustLength failed")
				}

				// Test Try methods with valid values
				if v.Get("string").TryString("fallback") != "test" {
					return fmt.Errorf("TryString failed with valid value")
				}
				if v.Get("number").TryInt(0) != 42 {
					return fmt.Errorf("TryInt failed with valid value")
				}
				if !v.Get("boolean").TryBool(false) {
					return fmt.Errorf("TryBool failed with valid value")
				}
				if v.Get("array").TryLength(0) != 3 {
					return fmt.Errorf("TryLength failed with valid value")
				}

				// Test Try methods with invalid values
				if v.Get("nonexistent").TryString("fallback") != "fallback" {
					return fmt.Errorf("TryString failed with invalid value")
				}
				if v.Get("nonexistent").TryInt(0) != 0 {
					return fmt.Errorf("TryInt failed with invalid value")
				}
				if v.Get("nonexistent").TryBool(false) != false {
					return fmt.Errorf("TryBool failed with invalid value")
				}
				if v.Get("nonexistent").TryLength(0) != 0 {
					return fmt.Errorf("TryLength failed with invalid value")
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
