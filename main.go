package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random

	t := dealgen.MultiPbnDeal(sh, 2, 10, 0, 0, 0, "20..17.0")
	fmt.Println(t)
}
