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
