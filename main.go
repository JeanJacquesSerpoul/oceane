package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random
	r, err := dealgen.DealMaskString(sh, dealgen.InitDeal, "AKQJT98765432", 1, 2)
	fmt.Println(r, err)
}
