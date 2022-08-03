package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random

	t := dealgen.PbnDeal(sh, 2, 1, 0, 0, 0, "20..17.0")
	fmt.Println(t)
}
