package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random
	r, err := dealgen.DealMaskString(sh, "AKQJT98765432", 0, 0)
	fmt.Println(r, err)
}
