package dealgen

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* Copyright 2020 Jean-Jacques Serpoul
under APACHE 2.0 license*
https://github.com/JeanJacquesSerpoul/oceane

Card Suit Card Value(D) Height Value(D) Suit Value(D) Card Value(B) Height Value(B) Suit Value(B)

2 S 0 0 0 000000 0000 00
2 H 1 0 1 000001 0000 01
2 D 2 0 2 000010 0000 10
2 C 3 0 3 000011 0000 11
3 S 4 1 0 000100 0001 00
3 H 5 1 1 000101 0001 01
3 D 6 1 2 000110 0001 10
3 C 7 1 3 000111 0001 11
4 S 8 2 0 001000 0010 00
4 H 9 2 1 001001 0010 01
4 D 10 2 2 001010 0010 10
4 C 11 2 3 001011 0010 11
5 S 12 3 0 001100 0011 00
5 H 13 3 1 001101 0011 01
5 D 14 3 2 001110 0011 10
5 C 15 3 3 001111 0011 11
6 S 16 4 0 010000 0100 00
6 H 17 4 1 010001 0100 01
6 D 18 4 2 010010 0100 10
6 C 19 4 3 010011 0100 11
7 S 20 5 0 010100 0101 00
7 H 21 5 1 010101 0101 01
7 D 22 5 2 010110 0101 10
7 C 23 5 3 010111 0101 11
8 S 24 6 0 011000 0110 00
8 H 25 6 1 011001 0110 01
8 D 26 6 2 011010 0110 10
8 C 27 6 3 011011 0110 11
9 S 28 7 0 011100 0111 00
9 H 29 7 1 011101 0111 01
9 D 30 7 2 011110 0111 10
9 C 31 7 3 011111 0111 11
10 S 32 8 0 100000 1000 00
10 H 33 8 1 100001 1000 01
10 D 34 8 2 100010 1000 10
10 C 35 8 3 100011 1000 11
J S 36 9 0 100100 1001 00
J H 37 9 1 100101 1001 01
J D 38 9 2 100110 1001 10
J C 39 9 3 100111 1001 11
Q S 40 10 0 101000 1010 00
Q H 41 10 1 101001 1010 01
Q D 42 10 2 101010 1010 10
Q C 43 10 3 101011 1010 11
K S 44 11 0 101100 1011 00
K H 45 11 1 101101 1011 01
K D 46 11 2 101110 1011 10
K C 47 11 3 101111 1011 11
A S 48 12 0 110000 1100 00
A H 49 12 1 110001 1100 01
A D 50 12 2 110010 1100 10
A C 51 12 3 110011 1100 11
*/
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
	return []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
		36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
}
func maskSuitToArray(s string) [][]int {
	a := make([][]int, N_OF_HANDS)
	for i := range a {
		a[i] = make([]int, N_OF_SUITS)
	}
	hand := strings.Split(s, SPACE)
	if len(hand) != N_OF_HANDS {
		return a
	}
	for i, v := range hand {
		if v != MINUS {
			suit := strings.Split(v, POINT)
			if len(suit) != N_OF_SUITS {
				return a
			}
			w1, err := strconv.Atoi(suit[0])
			if err == nil {
				a[i][0] = w1
			}
			w2, err := strconv.Atoi(suit[1])
			if err == nil {
				a[i][1] = w2
			}
			w3, err := strconv.Atoi(suit[2])
			if err == nil {
				a[i][2] = w3
			}
			w4, err := strconv.Atoi(suit[3])
			if err == nil {
				a[i][3] = w4
			}
		}
	}
	return a
}

func maskToArray(pbn string) ([]int, []int) {
	s := maskConvertToArray(pbn)
	var r, w, d []int
	for i := 0; i < N_OF_HANDS; i++ {
		r = nil
		for j := 0; j < N_OF_SUITS; j++ {
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

func maskConvertToArray(pbn string) [][]string {
	a := make([][]string, N_OF_HANDS)
	for i := range a {
		a[i] = make([]string, 4)
	}
	hand := strings.Split(pbn, SPACE)
	if len(hand) != N_OF_HANDS {
		return a
	}
	for i, v := range hand {
		if v != MINUS {
			suit := strings.Split(v, POINT)
			if len(suit) != N_OF_SUITS {
				return a
			}
			copy(a[i], suit)
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
	return ""
}

func convertCardsToString(a []int) string {
	r := ""
	for _, value := range a {
		if value == -1 {
			r += ""
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

	for i := 0; i < N_OF_SUITS; i++ {
		var v []int
		v = append(v, getSuitFromHand(h, i)...)
		v = sortHand(v)
		r += convertCardsToString(v)
		if i < N_OF_SUITS-1 {
			r += POINT
		}
	}
	return r
}

func getValueCards(i int) int {
	valueCards := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4}
	return valueCards[i]
}

func pointsFromHand(h []int) int {
	v := 0
	for _, value := range h {
		v += getValueCards(cardValueInt(value))
	}
	return v
}

func pbnDealSimple(a []int) string {
	var h []int
	r := ""
	for i := 0; i < N_OF_HANDS; i++ {
		h = a[i*N_HANDS : i*N_HANDS+N_HANDS]
		r += handPbn(h)
		if i < 3 {
			r += SPACE
		}
	}
	return r
}
