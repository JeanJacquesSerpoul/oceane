package main

import (
	_ "crypto/sha512"
	"syscall/js"

	"oceane/dealgen"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("wasmDealGen", js.FuncOf(dealGen))
	<-done
}

func dealGen(this js.Value, args []js.Value) interface{} {
	var sh dealgen.Random

	return dealgen.PbnDeal(sh, 2, 1, 0, 0, 0, "...")
}
