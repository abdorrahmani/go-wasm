//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"

	wrapper "github.com/abdorrahmani/go-wasm/js"
)

func main() {
	// Get the global object
	global := wrapper.Global()

	// Create a new object
	obj := global.Call("Object")
	obj.Set("name", "John")
	obj.Set("age", 30)

	// Get properties with type safety
	name, err := obj.Get("name").String()
	if err != nil {
		fmt.Printf("Error getting name: %v\n", err)
		return
	}

	age, err := obj.Get("age").Int()
	if err != nil {
		fmt.Printf("Error getting age: %v\n", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Create an array
	arr := global.Call("Array", 1, 2, 3, 4, 5)

	// Get array values
	values, err := arr.Array()
	if err != nil {
		fmt.Printf("Error getting array: %v\n", err)
		return
	}

	fmt.Println("Array values:")
	for i, v := range values {
		num, err := v.Int()
		if err != nil {
			fmt.Printf("Error getting value at index %d: %v\n", i, err)
			continue
		}
		fmt.Printf("  [%d] = %d\n", i, num)
	}

	// Keep the program running
	select {}
}
