package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random
	r := "6.4.1.2 ..0. ... ..."
	t := dealgen.MultiPbnDealToFile(sh, "test.pbn", 1, 100, 1, 0, 0, r)
	fmt.Println(t)
}
