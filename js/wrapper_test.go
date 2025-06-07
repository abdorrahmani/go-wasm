//go:build js && wasm
// +build js,wasm

package js

import (
	"fmt"
	"testing"
)

func TestValueWrapper(t *testing.T) {
	// Note: These tests can only run in a WebAssembly environment
	// They are here as documentation and for when running in WASM
	t.Skip("Skipping tests that require WebAssembly environment")

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
				arr := Global().Call("Array", 1, 2, 3)
				return arr
			},
			validate: func(v *Value) error {
				arr, err := v.Array()
				if err != nil {
					return err
				}
				if len(arr) != 3 {
					return fmt.Errorf("expected array length 3, got %d", len(arr))
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := tt.setup()
			if err := tt.validate(value); err != nil {
				t.Errorf("test failed: %v", err)
			}
		})
	}
}
