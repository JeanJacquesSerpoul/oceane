package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random
	r := dealgen.DealSuitString(sh, "0.0.0. ... ... ...")
	fmt.Println(r)
}
