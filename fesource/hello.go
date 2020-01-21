package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WASM !!")
	js.Global().Set("helloAgain", js.FuncOf(helloAgain))
}

func helloAgain(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("Hello again, WASM !!")
	return nil
}
