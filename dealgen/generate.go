package dealgen

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
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

func freeRandom(sh ShuffleInterface, a []int) []int {
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

func getInitDeal() []int {
	return initDeal
}

func maskToArray(pbn string) ([]int, []int) {
	s := maskConvertToArray(pbn)
	var r, w, d []int
	for i := 0; i < 4; i++ {
		r = nil
		for j := 0; j < 4; j++ {
			b := s[i][j]
			if b != "" {
				r = append(r, maskStrToMaskInt(b, j)...)
			}
		}
		l := len(r)
		for k := 0; k < N_HANDS-l; k++ {
			r = append(r, -1)
		}
		w = append(w, r...)
	}
	d = delta(getInitDeal(), w)
	return w, d
}

func maskConvertToArray(pbn string) [4][4]string {
	var a [4][4]string
	var hand []string
	hand = strings.Split(pbn, SPACE)
	for i, v := range hand {
		if v != MINUS {
			suit := strings.Split(v, POINT)
			for j, w := range suit {
				a[i][j] = w
			}
		}
	}
	return a
}

func DealMaskString(sh ShuffleInterface, mask string) string {
	deal, delta := maskToArray(mask)
	s := freeRandom(sh, delta)
	k := 0
	for i, value := range deal {
		if value == -1 {
			deal[i] = s[k]
			k++
		}
	}
	return pbnDealSimple(deal)
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

func maskStrToMaskInt(v string, suit int) []int {
	var a []int
	for i := 0; i < len(v); i++ {
		x := string(v[i])
		if (x == "2") || (x == "3") || (x == "4") || (x == "5") || (x == "6") || (x == "7") || (x == "8") || (x == "9") {
			t, _ := strconv.Atoi(x)
			t -= 2
			a = append(a, (t<<2)+suit)
		}
		if x == "T" {
			a = append(a, (8<<2)+suit)
		}
		if x == "J" {
			a = append(a, (9<<2)+suit)
		}
		if x == "Q" {
			a = append(a, (10<<2)+suit)
		}
		if x == "K" {
			a = append(a, (11<<2)+suit)
		}
		if x == "A" {
			a = append(a, (12<<2)+suit)
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
			r += POINT
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

func pbnDealSimple(a []int) string {
	var h []int
	r := ""
	for i := 0; i <= 3; i++ {
		h = a[i*N_HANDS : i*N_HANDS+N_HANDS]
		r += handPbn(h)
		if i < 3 {
			r += SPACE
		}
	}
	return r
}
