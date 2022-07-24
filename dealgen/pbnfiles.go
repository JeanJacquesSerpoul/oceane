package dealgen

import "os"

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

func pbnDeal(firstHand, dealer, vul int, deal string) string {
	r := "[Dealer \"" + position(dealer) + "\"]\n"
	r += "[Vulnerable \"" + vulnerable(vul) + "\"]\n"
	r += "[Deal \"" + position(firstHand) + ":"
	r += deal
	r += "\"]"
	return r
}

func MultiPbnDeal(sh ShuffleInterface, mode, ite, firstHand, dealer, vul int, mask string) string {
	deal := ""
	r := ""
	for i := 0; i < ite; i++ {
		if mode == 0 {
			deal = DealMaskString(sh, mask)
		}
		if mode == 1 {
			deal = DealSuitString(sh, mask)
		}
		r += pbnDeal(firstHand, dealer, vul, deal)
		r += "\n\n"
	}
	return r
}

func MultiPbnDealToFile(sh ShuffleInterface, filename string, mode, ite, firstHand, dealer, vul int, mask string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	r := MultiPbnDeal(sh, mode, ite, firstHand, dealer, vul, mask)
	_, _ = file.WriteString(r)
	defer file.Close()
	return nil
}
