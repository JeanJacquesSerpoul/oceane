package dealgen

import "sort"

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

func handPbn(h []int) string {
	r := ""

	for i := 0; i <= 3; i++ {
		var v []int
		v = append(v, getSuitFromHand(h, i)...)
		sort.Slice(v, func(i, j int) bool {
			return CardValueInt(v[i]) > CardValueInt(v[j])
		})
		r += convertCardsToString(v)
		if i < 3 {
			r += "."
		}
	}
	return r
}
func PbnDealSimple(a []int) string {
	var h []int
	r := ""
	for i := 0; i <= 3; i++ {
		h = a[i*13 : i*13+13]
		r += handPbn(h)
		if i < 3 {
			r += " "
		}
	}
	return r
}

func PbnDeal(firstHand, dealer, vul int, a []int) string {
	r := "[Dealer \"" + position[dealer] + "\"]\n"
	r += "[Vulnerable \"" + vulnerable[vul] + "\"]\n"
	r += "[Deal \"" + position[firstHand] + ":"
	r += PbnDealSimple(a)
	r += "\"]"
	return r
}
