package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	r := dealgen.FYShuffle(52)
	fmt.Println(dealgen.PbnDeal(0, 0, 0, r))
}
