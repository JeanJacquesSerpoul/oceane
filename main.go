package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random
	r := dealgen.DealMaskString(sh, "AKQJT98765432", 1, 3)
	fmt.Println(r)
}
