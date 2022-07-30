package dealgen

import (
	"os"
	"sort"
)

func position(i int) string {
	if i < 0 || i > N_OF_HANDS {
		i = 0
	}
	v := []string{"N", "E", "S", "W"}
	return v[i]
}

func vulnerable(i int) string {
	if i < 0 || i > N_OF_HANDS {
		i = 0
	}
	v := []string{"ALL", "NS", "EW", "NONE"}
	return v[i]
}

func simplePbnDeal(firstHand, dealer, vul int, deal string) string {
	r := "[Dealer \"" + position(dealer) + "\"]\n"
	r += "[Vulnerable \"" + vulnerable(vul) + "\"]\n"
	r += "[Deal \"" + position(firstHand) + ":"
	r += deal
	r += "\"]"
	return r
}

func checkDeal(deal []int) int {
	if len(deal) != N_CARDS {
		return 1
	}
	dealCopy := make([]int, len(deal))
	copy(dealCopy, deal)
	sort.Slice(dealCopy, func(i, j int) bool {
		return dealCopy[i] > dealCopy[j]
	})
	for i := 0; i < len(dealCopy); i++ {
		if dealCopy[i] != len(dealCopy)-i-1 {
			return 2
		}
	}
	return 0
}

func PbnDeal(sh ShuffleInterface, mode, ite, firstHand, dealer, vul int, mask string) string {
	deal := ""
	rs := ""
	var r []int
	i := 0
	k := 0
	for i < ite {
		if mode == 0 {
			r = DealMaskArray(sh, mask)
		}
		if mode == 1 {
			r = DealSuitArray(sh, mask)
		}
		if mode == 2 {
			r, _ = DealPointArrayPlus(sh, mask)
		}
		if checkDeal(r) == 0 {
			deal = pbnDealSimple(r)
			rs += simplePbnDeal(firstHand, dealer, vul, deal)
			rs += "\n\n"
			i++
		}
		k++
		if k == INFINITE {
			return ""
		}
	}
	return rs
}

func PbnDealToFile(sh ShuffleInterface, filename string, mode, ite, firstHand, dealer, vul int, mask string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	r := PbnDeal(sh, mode, ite, firstHand, dealer, vul, mask)
	_, _ = file.WriteString(r)
	defer file.Close()
	return nil
}
