package dealgen

import (
	"encoding/json"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Shuffle Implement Fisher and Yates shuffle method
func (rd Random) fYShuffle(n int) []int {
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

func FreeRandom(sh ShuffleInterface, a []int) []int {
	r := sh.fYShuffle(len(a))
	t := make([]int, len(a))
	for i, value := range r {
		t[value] = a[i]
	}
	return t
}

func intInSlice(a int, list []int) int {
	for _, vlist := range list {
		if vlist == a {
			return a
		}
	}
	return -1
}

func delta(slice []int, ToRemove []int) []int {
	var diff []int

	var n int

	for _, vslice := range slice {
		n = intInSlice(vslice, ToRemove)
		if n < 0 {
			diff = append(diff, vslice)
		}
	}
	return diff
}

func cardSuitToValue(cardValue, suit int) int {
	return (cardValue << 2) + suit
}

func DealMaskSuit(maskSuit []int, suit int) []int {
	var r []int
	for _, value := range maskSuit {
		v := cardSuitToValue(value, suit)
		r = append(r, v)
	}
	return r
}

func DealMask(sh ShuffleInterface, deal, maskSuit []int, suit, hand int) []int {
	var r, d, mask []int
	dm := DealMaskSuit(maskSuit, suit)
	for range deal {
		mask = append(mask, -1)
	}
	for i, value := range dm {
		mask[hand*N_HANDS+i] = value
	}
	d = delta(deal, mask)
	s := FreeRandom(sh, d)
	k := 0
	for i, value := range mask {
		if value >= 0 && (i >= hand*N_HANDS && i < hand*N_HANDS+N_HANDS) {
			r = append(r, value)
		} else {
			r = append(r, s[k])
			k++
		}
	}
	return r
}

func cardValueInt(cardValue int) int { return cardValue >> 2 }

func cardSuitInt(cardValue int) int { return cardValue & 3 }

func getSuitFromHand(h []int, suitValue int) []int {
	var r []int
	for _, value := range h {
		if cardSuitInt(value) == suitValue {
			r = append(r, value)
		}
	}
	return r
}

func MaskStrToMaskInt(v string) []int {
	var a []int
	for i := 0; i < len(v); i++ {
		x := string(v[i])
		if (x == "2") || (x == "3") || (x == "4") || (x == "5") || (x == "6") || (x == "7") || (x == "8") || (x == "9") {
			t, _ := strconv.Atoi(x)
			t -= 2
			a = append(a, t)
		}
		if x == "T" {
			a = append(a, 8)
		}
		if x == "J" {
			a = append(a, 9)
		}
		if x == "Q" {
			a = append(a, 10)
		}
		if x == "K" {
			a = append(a, 11)
		}
		if x == "A" {
			a = append(a, 12)
		}
	}
	return a
}

func getFaceCard(v int) string {
	if v <= 7 {
		return strconv.Itoa(v + 2)
	} else {
		if v == 8 {
			return "T"
		}
		if v == 9 {
			return "J"
		}
		if v == 10 {
			return "Q"
		}
		if v == 11 {
			return "K"
		}
		if v == 12 {
			return "A"
		}
	}
	return ERRORMSG
}

func convertCardsToString(a []int) string {
	r := ""
	for _, value := range a {
		if value == -1 {
			r += UNDEF
		} else {
			v := cardValueInt(value)
			r += getFaceCard(v)
		}
	}
	return r
}

func sortHand(h []int) []int {
	sort.Slice(h, func(i, j int) bool {
		return cardValueInt(h[i]) > cardValueInt(h[j])
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
		v += valueCards[cardValueInt(value)]
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

func jsonStructDeal(firstHand, dealer, vul int, a []int) string {
	result := structDeal(firstHand, dealer, vul, a)
	r, _ := json.Marshal(result)
	return string(r)
}
