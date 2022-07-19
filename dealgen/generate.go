package dealgen

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
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

func convertCardsToString(a []int) string {
	r := ""
	for _, value := range a {
		r += faceCards[cardValueInt(value)]
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

func pbnDealSimple(a []int) string {
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
	r := "% Index: " + ArrayToIndex(a) + "\n"
	r += "[Dealer \"" + position[dealer] + "\"]\n"
	r += "[Vulnerable \"" + vulnerable[vul] + "\"]\n"
	r += "[Deal \"" + position[firstHand] + ":"
	r += pbnDealSimple(a)
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
func getHandFromDist(index int) int {
	return index / N_HANDS
}
func convertDistToIndexArray(content []int, index *[N_CARDS]int) {
	var suit, height int
	for i, v := range content {
		suit = 3 - cardSuitInt(v)
		height = cardValueInt(v)
		index[suit*N_HANDS+height] = getHandFromDist(i)
	}
}

// From java code Copyright (@)1999, Thomas Andrews
//http://bridge.thomasoandrews.com/impossible/
// Free for non-commercial use

func fraction(val *big.Int, num, den int) *big.Int {
	numer := int64(num)
	denom := int64(den)
	v := new(big.Int)
	n := big.NewInt(numer)
	d := big.NewInt(denom)
	v.Mul(val, n)
	v.Div(v, d)
	return v
}

func pages() *big.Int {
	v := new(big.Int)
	v.SetString((NbDist), 10)

	return v
}

func ArrayToIndex(content []int) string {
	var r string
	var cardsNeeded = [FOUR]int{N_HANDS, N_HANDS, N_HANDS, N_HANDS}
	var hand, skipped, goesTo int
	var index [N_CARDS]int
	convertDistToIndexArray(content, &index)
	width := pages()
	minimum := big.NewInt(0)
	for c := N_CARDS; c > 0; c-- {
		hand = 0
		skipped = 0
		goesTo = index[c-1]
		for hand < goesTo {
			skipped += cardsNeeded[hand]
			hand++
		}
		minimum.Add(minimum, fraction(width, skipped, c))
		width = fraction(width, cardsNeeded[goesTo], c)
		cardsNeeded[goesTo]--
	}
	r = fmt.Sprintf("%v", minimum)
	return r
}

func structDeal(firstHand, dealer, vul int, a []int) result {
	var r result
	r.PbnSimple = pbnDealSimple(a)
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
