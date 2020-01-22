package main

import (
	"fmt"
	"syscall/js"
)

var killSignal chan bool

func init() {
	killSignal = make(chan bool)
}

func main() {
	fmt.Println("Hello, WASM !!")
	js.Global().Set("helloAgain", js.FuncOf(helloAgain))
	<-killSignal
}

func helloAgain(this js.Value, inputs []js.Value) interface{} {

	n := js.Global().Get("noneExistingVar")
	printValue("noneExistingVar", n)

	js.Global().Set("noneExistingVar", "1")
	n = js.Global().Get("noneExistingVar")
	printValue("noneExistingVar", n)

	js.Global().Set("noneExistingVar", nil)
	n = js.Global().Get("noneExistingVar")
	printValue("noneExistingVar", n)

	js.Global().Set("noneExistingVar", true)
	n = js.Global().Get("noneExistingVar")
	printValue("noneExistingVar", n)

	a := js.Global().Get("alert")
	a.Invoke("Long leave the Gopher !")

	j := js.Global().Get("JSON")
	ret := j.Call("parse", `{"name":"gopher", "age" : 10}`)

	return ret
}

func printValue(name string, v js.Value) {
	var ret string
	if v == js.Undefined() {
		ret = " is undefined"
	} else if v == js.Null() {
		ret = " is null"
	} else if v.Type() == js.TypeBoolean {
		ret = " bool value = " + fmt.Sprint(v.Bool())
	} else if v.Type() == js.TypeNumber {
		ret = " num value = " + fmt.Sprint(v.Float())
	} else if v.Type() == js.TypeString {
		ret = " str value = " + v.String()
	} else {
		//we can go on ...
	}
	fmt.Println(name + ret)
}
