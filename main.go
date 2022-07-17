package main

import (
	"fmt"

	"oceane/dealgen"
)

func main() {
	d := new(dealgen.Desk)
	d.Init()
	t := make([]int, dealgen.N_CARDS)
	t = dealgen.FYShuffle(dealgen.N_CARDS)
	fmt.Println(t)
}
