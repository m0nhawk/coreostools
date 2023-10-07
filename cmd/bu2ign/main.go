package main

import (
	"fmt"
	"github.com/coreos/butane/config"
	"github.com/coreos/butane/config/common"
	"syscall/js"
)

// https://medium.com/vacatronics/how-to-write-a-webassembly-app-in-go-fd769fa2b64b
// https://github.com/coreos/butane
// https://coreos.github.io/butane/getting-started/

func translateButane(input []byte) string {
	options := common.TranslateBytesOptions{}
	var err error
	fmt.Println("%v", input)
	dataOut, _, err := config.TranslateBytes(input, options)
	fmt.Println("%v", string(dataOut))
	if err != nil {
		println("Error translating config: %v\n", err)
	}
	return string(dataOut)
}

func jsTranslateButane(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid no of args passed"
	}
	square := translateButane([]byte(args[0].String()))
	return square
}

func main() {
	c := make(chan bool)
	js.Global().Set("translateButane", js.FuncOf(jsTranslateButane))
	<-c
}
