package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
	"time"
)

var killSignal chan bool

func init() {
	killSignal = make(chan bool)
}

func main() {
	fmt.Println("Hello, WASM !!")
	js.Global().Set("helloAgain", js.FuncOf(helloAgain))
	js.Global().Set("callgithub", js.FuncOf(callgithub))
	<-killSignal
}

func callgithub(this js.Value, inputs []js.Value) interface{} {
	url := "https://api.github.com"
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err.Error())
	} else {
		defer resp.Body.Close()
		if bt, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Println(err.Error())
		} else {
			d := js.Global().Get("document")
			elem := d.Call("getElementById", "replyDiv")
			elem.Set("innerHTML", string(bt))
		}
	}
	return nil
}

var times int

func helloAgain(this js.Value, inputs []js.Value) interface{} {
	times++
	type object = map[string]interface{}
	o := object{"createdByGo": true, "creator": "hello.go", "timeCreated": time.Now().Format(time.RFC3339)}
	subs := object{}
	for i, inp := range inputs {
		subs[inp.String()] = i * times
	}
	o["elements"] = subs
	js.Global().Set("helloObject", js.ValueOf(o)) //json.Marshal a struct could be considered
	js.Global().Call("updateState")
	return nil
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
	} else if v.Type() == js.TypeFunction {
		ret = " is a func "
	} else if v.Type() == js.TypeObject {
		ret = " is an object "
	} else {
		ret = " is TBD "
	}
	fmt.Println(name + ret)
}
