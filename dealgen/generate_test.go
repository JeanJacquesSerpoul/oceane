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

var mockHand = []int{
	17, 18, 19, 20, 21, 22, 45, 24, 25, 26, 50, 28, 29,
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

type fakeRandom struct {
}

func (test *fakeRandom) fYShuffle(n int) []int {
	var random, temp int
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = i
	}
	for i := len(t) - 1; i >= 0; i-- {
		temp = t[i]
		random = i
		t[i] = t[random]
		t[random] = temp
	}
	return t
}

func Test_fYshuffle(t *testing.T) {
	t.Parallel()
	sh := new(Random)
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
