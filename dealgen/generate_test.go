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

var mockRandom = []int{
	44, 39, 13, 33, 43, 37, 47, 51, 28, 0, 14, 46, 48, 35, 21, 27,
	30, 40, 42, 3, 22, 31, 17, 36, 19, 5, 25, 24, 10, 20, 26,
	50, 49, 45, 4, 38, 6, 16, 23, 32, 2, 29, 41, 34, 8, 1, 9, 18, 12, 15, 11, 7,
}

var mockDeal = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 45, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 9, 46, 47, 48, 49, 50, 51,
}

//  "Obsure Bug if reuse mockDeal
var mockDealB = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 45, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 9, 46, 47, 48, 49, 50, 51,
}

var mockResultRandom = []int{
	51, 50, 49, 48, 47, 46, 9, 44, 43, 42, 41, 40, 39, 38, 37,
	36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20,
	19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 45, 8, 7, 6, 5, 4, 3, 2, 1, 0,
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
		51, 50, 48, 47, 46, 44, 43, 42, 40, 39, 38, 36, 35, 34, 32, 31, 30, 28, 27, 26, 24,
		23, 22, 20, 19, 18, 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45,
		49, 16, 15, 14, 12, 11, 10, 8, 7, 6, 4, 3, 2, 0,
	}
)

type FakeRandom struct{}

func (test FakeRandom) fYShuffle(n int) []int {
	var r []int
	for i := 0; i < n; i++ {
		r = append(r, n-i-1)
	}
	return r
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
		{"Test1", args{mockDeal}, mockPbnSimple},
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
		{"Test1", args{0, 1, 2, mockDeal}, mockPbn},
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
		{"Test1", args{0, 1, 2, mockDeal}, mockResultPbn},
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
		{"Test1", args{0, 1, 2, mockDeal}, getStr(mockResultPbn)},
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
		{"Test1", args{mockDealB}, mockResultRandom},
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
		{"Test1", args{mockDealB, mockMaskSuite, 2, 1}, mockResultMaskSuite},
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
		deal []int
		suit int
		hand int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{"AKQJT98765432", mockDealB, 2, 1}, "AKQJ..AKQJ.AKQJT T987..T9876.9876 .AKQJT98765432.. 65432..5432.5432", false},
		{"Test1", args{"AKQJT98765432", mockDealB, 2, 8}, "", true},
		{"Test1", args{"AKQJT98765432", mockDealB, 8, 1}, "", true},
	}
	var sh FakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DealMaskString(sh, tt.args.deal, tt.args.mask, tt.args.suit, tt.args.hand)
			if (err != nil) != tt.wantErr {
				t.Errorf("DealMaskString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DealMaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
