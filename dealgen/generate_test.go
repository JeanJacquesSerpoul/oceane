package dealgen

import (
	"encoding/json"
	"reflect"
	"testing"
)

var mockResultPbn = result{
	"5432.K32.432.432 876.8765.765.765 JT9.JT9.JT98.T98 AKQ.AQ4.AKQ.AKQJ",
	`[Dealer "E"]
[Vulnerable "EW"]
[Deal "N:5432.K32.432.432 876.8765.765.765 JT9.JT9.JT98.T98 AKQ.AQ4.AKQ.AKQJ"]`,
	[4]int{3, 0, 3, 34},
	[4][4][]int{
		{{12, 8, 4, 0}, {45, 5, 1}, {10, 6, 2}, {11, 7, 3}},
		{{24, 20, 16}, {25, 21, 17, 13}, {22, 18, 14}, {23, 19, 15}},
		{{36, 32, 28}, {37, 33, 29}, {38, 34, 30, 26}, {35, 31, 27}},
		{{48, 44, 40}, {49, 41, 9}, {50, 46, 42}, {51, 47, 43, 39}},
	},
}

var mockInitDeal = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 45, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 9, 46, 47, 48, 49, 50, 51,
}

var mockRandom = []int{
	44, 39, 13, 33, 43, 37, 47, 51, 28, 0, 14, 46, 48, 35, 21, 27,
	30, 40, 42, 3, 22, 31, 17, 36, 19, 5, 25, 24, 10, 20, 26,
	50, 49, 45, 4, 38, 6, 16, 23, 32, 2, 29, 41, 34, 8, 1, 9, 18, 12, 15, 11, 7,
}

var mockResultRandom = []int{
	3, 44, 48, 19, 30, 13, 28, 9, 47, 40, 36, 39, 42, 11, 2, 43, 27,
	16, 41, 14, 35, 24, 18, 26, 37, 38, 34, 23, 4, 49, 22, 17, 50, 10, 46, 25,
	15, 5, 29, 12, 21, 51, 20, 8, 45, 31, 1, 6, 0, 32, 33, 7,
}

var mockHand = []int{
	17, 18, 19, 20, 21, 22, 45, 24, 25, 26, 50, 28, 29,
}

var mockHandWithUndef = []int{
	17, 18, 19, 20, -1, 22, 45, 24, 25, 26, 50, 28, 29,
}

var mockSuitHand = []int{
	17, 21, 45, 25, 29,
}

var mockSortHand = []int{
	50, 45, 28, 29, 24, 25, 26, 21, 22, 20, 19, 18, 17,
}

var (
	mockStringHand = "987.K9876.A876.6"
	mockPbnSimple  = "5432.K32.432.432 876.8765.765.765 JT9.JT9.JT98.T98 AKQ.AQ4.AKQ.AKQJ"
	mockPbn        = `[Dealer "E"]
[Vulnerable "EW"]
[Deal "N:5432.K32.432.432 876.8765.765.765 JT9.JT9.JT98.T98 AKQ.AQ4.AKQ.AKQJ"]`
)

var mockDealMask = []int{
	0, 1, 2, 3, -1, -1, -1, 7, 8, 9, 10, -1, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, -1, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, -1, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
}

var mockResultDealMask = []int{
	0, 1, 2, 3, 11, 40, 22, 7, 8, 9, 10, 5, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 4, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 6, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
}

var (
	mockMaskSuite       = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	mockResultMaskSuite = []int{
		6, 46, 26, 24, 34, 16, 31, 39, 51, 47, 38, 43, 40, 12, 7,
		42, 30, 20, 44, 18, 36, 3, 23, 28, 14, 15, 1, 5, 9, 13, 17,
		21, 25, 29, 33, 37, 41, 45, 49, 35, 2, 8, 50, 0, 22, 27, 11, 48, 4, 19, 10, 32,
	}
)

type FakeRandom struct{}

func extractRandom(a []int, n int) []int {
	var r []int
	for i := 0; i < N_CARDS; i++ {
		if a[i] < n {
			r = append(r, a[i])
		}
	}
	return r
}

func (test FakeRandom) fYShuffle(n int) []int {
	return extractRandom(mockRandom, n)
}

func Test_fYshuffle(t *testing.T) {
	var sh Random
	t.Parallel()
	t1 := sh.fYShuffle(N_CARDS)
	t2 := sh.fYShuffle(N_CARDS)
	if reflect.DeepEqual(t1, t2) {
		t.Errorf("fYshuffle is not working")
	}
}

func TestCardValueInt(t *testing.T) {
	type args struct {
		cardValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{17}, 4},
		{"Test2", args{22}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardValueInt(tt.args.cardValue); got != tt.want {
				t.Errorf("CardValueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSuitInt(t *testing.T) {
	type args struct {
		cardValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{17}, 1},
		{"Test1", args{19}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardSuitInt(tt.args.cardValue); got != tt.want {
				t.Errorf("CardSuitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertCardsToString(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{mockHand}, "666777K888A99"},
		{"Test1", args{mockHandWithUndef}, "6667?7K888A99"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCardsToString(tt.args.a); got != tt.want {
				t.Errorf("convertCardsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSuitFromHand(t *testing.T) {
	type args struct {
		h         []int
		suitValue int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{mockHand, 1}, mockSuitHand},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSuitFromHand(tt.args.h, tt.args.suitValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSuitFromHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortHand(t *testing.T) {
	type args struct {
		h []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{mockHand}, mockSortHand},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortHand(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handPbn(t *testing.T) {
	type args struct {
		h []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{mockHand}, mockStringHand},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handPbn(tt.args.h); got != tt.want {
				t.Errorf("handPbn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pointsFromHand(t *testing.T) {
	type args struct {
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{mockHand}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointsFromHand(tt.args.h); got != tt.want {
				t.Errorf("pointsFromHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pbnDealSimple(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{mockInitDeal}, mockPbnSimple},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pbnDealSimple(tt.args.a); got != tt.want {
				t.Errorf("pbnDealSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pbnDeal(t *testing.T) {
	type args struct {
		firstHand int
		dealer    int
		vul       int
		a         []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{0, 1, 2, mockInitDeal}, mockPbn},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pbnDeal(tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.a); got != tt.want {
				t.Errorf("pbnDeal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structDeal(t *testing.T) {
	type args struct {
		firstHand int
		dealer    int
		vul       int
		a         []int
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{"Test1", args{0, 1, 2, mockInitDeal}, mockResultPbn},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := structDeal(tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structDeal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getStr(r result) string {
	rs, _ := json.Marshal(r)
	return string(rs)
}

func TestJsonStructDeal(t *testing.T) {
	type args struct {
		firstHand int
		dealer    int
		vul       int
		a         []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{0, 1, 2, mockInitDeal}, getStr(mockResultPbn)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonStructDeal(tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.a); got != tt.want {
				t.Errorf("JsonStructDeal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreeRandom(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{mockInitDeal}, mockResultRandom},
	}
	var sh FakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freeRandom(sh, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreeRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDealMask(t *testing.T) {
	type args struct {
		deal     []int
		maskSuit []int
		suit     int
		hand     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{mockInitDeal, mockMaskSuite, 1, 2}, mockResultMaskSuite},
	}
	var sh FakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dealMask(sh, tt.args.deal, tt.args.maskSuit, tt.args.suit, tt.args.hand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DealMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFaceCard(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{100}, ERRORMSG},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFaceCard(tt.args.v); got != tt.want {
				t.Errorf("getFaceCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaskStrToMaskInt(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{"AKQJT98765432"}, []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskStrToMaskInt(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaskStrToMaskInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDealMaskString(t *testing.T) {
	type args struct {
		mask string
		suit int
		hand int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"AKQJT98765432", 1, 2}, "KT..AQJ98753.AKQ A872..KT64.T8753 .AKQJT98765432.. QJ96543..2.J9642"},
	}
	var sh FakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DealMaskString(sh, tt.args.mask, tt.args.suit, tt.args.hand); got != tt.want {
				t.Errorf("DealMaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
