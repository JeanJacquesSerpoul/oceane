package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var hand = []int{
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	}
	var sh dealgen.Random
	fmt.Println(dealgen.FreeRandom(sh, hand))
}
