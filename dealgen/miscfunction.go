package dealgen

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Shuffle Implement Fisher and Yates shuffle method
func FYShuffle(n int) []int {
	var random, temp int
	t := make([]int, n)
	for i := range t {
		t[i] = i
	}
	for i := len(t) - 1; i >= 0; i-- {
		temp = t[i]
		random = rand.Intn(i + 1)
		t[i] = t[random]
		t[random] = temp
	}
	return t
}
