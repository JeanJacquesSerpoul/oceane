package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random

	t, err := dealgen.DealPointsString(sh, "16.5.0.")
	fmt.Println(t, err)
}
