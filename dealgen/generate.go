package dealgen

import (
	"encoding/json"
	"math/rand"
	"sort"
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

func (c CardList) Shuffle(a []int) []int {
	v := make([]int, len(a))
	v = FYShuffle(len(a))
	return v
}

func CardValueInt(cardValue int) int { return cardValue >> 2 }

func CardSuitInt(cardValue int) int { return cardValue & 3 }

func getSuitFromHand(h []int, suitValue int) []int {
	var r []int
	for _, value := range h {
		if CardSuitInt(value) == suitValue {
			r = append(r, value)
		}
	}
	return r
}

func convertCardsToString(a []int) string {
	r := ""
	for _, value := range a {
		r += faceCards[CardValueInt(value)]
	}
	return r
}

func sortHand(h []int) []int {
	sort.Slice(h, func(i, j int) bool {
		return CardValueInt(h[i]) > CardValueInt(h[j])
	})
	return h
}

func handPbn(h []int) string {
	r := ""

	for i := 0; i <= 3; i++ {
		var v []int
		v = append(v, getSuitFromHand(h, i)...)
		v = sortHand(v)
		r += convertCardsToString(v)
		if i < 3 {
			r += "."
		}
	}
	return r
}

func pointsFromHand(h []int) int {
	v := 0
	for _, value := range h {
		v += valueCards[CardValueInt(value)]
	}
	return v
}

func PbnDealSimple(a []int) string {
	var h []int
	r := ""
	for i := 0; i <= 3; i++ {
		h = a[i*N_HANDS : i*N_HANDS+N_HANDS]
		r += handPbn(h)
		if i < 3 {
			r += " "
		}
	}
	return r
}

func pbnDeal(firstHand, dealer, vul int, a []int) string {
	r := "[Dealer \"" + position[dealer] + "\"]\n"
	r += "[Vulnerable \"" + vulnerable[vul] + "\"]\n"
	r += "[Deal \"" + position[firstHand] + ":"
	r += PbnDealSimple(a)
	r += "\"]"
	return r
}

func getHandPoints(r result, a []int) result {
	for i := 0; i <= 3; i++ {
		r.HandPoints[i] = pointsFromHand(a[i*N_HANDS : i*N_HANDS+N_HANDS])
	}
	return r
}

func getSuitPoints(r result, a []int) result {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			s := make([]int, len(getSuitFromHand(a[i*N_HANDS:i*N_HANDS+N_HANDS], j)))
			s = getSuitFromHand(sortHand(a[i*N_HANDS:i*N_HANDS+N_HANDS]), j)
			r.Suit[i][j] = append(r.Suit[i][j], s...)
		}
	}
	return r
}

func structDeal(firstHand, dealer, vul int, a []int) result {
	var r result
	r.PbnSimple = PbnDealSimple(a)
	r.Pbn = pbnDeal(firstHand, dealer, vul, a)
	r = getHandPoints(r, a)
	r = getSuitPoints(r, a)
	return r
}
func JsonStructDeal(firstHand, dealer, vul int, a []int) string {
	result := structDeal(firstHand, dealer, vul, a)
	r, _ := json.Marshal(result)
	return string(r)
}
