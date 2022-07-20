package main

import (
	"fmt"
	"oceane/dealgen"
)

func main() {
	var sh dealgen.Random
	r := dealgen.DealMaskString(sh, "AK4.KJ.4.KT987 62.Q6.KJT8.A53 QT8..97532.4 753.T95.6.QJ62")
	fmt.Println(r)
}
