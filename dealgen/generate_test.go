package dealgen

import (
	"reflect"
	"testing"
)

func mockRandom() []int {
	return []int{
		44, 39, 13, 33, 43, 37, 47, 51, 28, 0, 14, 46, 48, 35, 21, 27,
		30, 40, 42, 3, 22, 31, 17, 36, 19, 5, 25, 24, 10, 20, 26,
		50, 49, 45, 4, 38, 6, 16, 23, 32, 2, 29, 41, 34, 8, 1, 9, 18, 12, 15, 11, 7,
	}
}

func mockDeal() []int {
	return []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 45, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
		36, 37, 38, 39, 40, 41, 42, 43, 44, 9, 46, 47, 48, 49, 50, 51,
	}
}

func mockResultRandom() []int {
	return []int{
		51, 50, 49, 48, 47, 46, 9, 44, 43, 42, 41, 40, 39, 38, 37,
		36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20,
		19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 45, 8, 7, 6, 5, 4, 3, 2, 1, 0,
	}
}

func mockHand() []int {
	return []int{
		17, 18, 19, 20, 21, 22, 45, 24, 25, 26, 50, 28, 29,
	}
}

func mockHandWithUndef() []int {
	return []int{
		17, 18, 19, 20, -1, 22, 45, 24, 25, 26, 50, 28, 29,
	}
}

func mockSuitHand() []int {
	return []int{
		17, 21, 45, 25, 29,
	}
}

func mockSortHand() []int {
	return []int{
		50, 45, 28, 29, 24, 25, 26, 21, 22, 20, 19, 18, 17,
	}
}

func mockStringHand() string {
	return "987.K9876.A876.6"
}

func mockPbnSimple() string {
	return "5432.K32.432.432 876.8765.765.765 JT9.JT9.JT98.T98 AKQ.AQ4.AKQ.AKQJ"
}

func mockPbn() string {
	return `[Dealer "E"]
	[Vulnerable "EW"]
	[Deal "N:5432.K32.432.432 876.8765.765.765 JT9.JT9.JT98.T98 AKQ.AQ4.AKQ.AKQJ"]`
}

func mockDealMask() []int {
	return []int{
		0, 1, 2, 3, -1, -1, -1, 7, 8, 9, 10, -1, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, -1, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
		36, 37, 38, 39, -1, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
}

func mockResultDealMask() []int {
	return []int{
		0, 1, 2, 3, 11, 40, 22, 7, 8, 9, 10, 5, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 4, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 6, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
}

func mockMaskSuite() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
}

func mockResultMaskSuite() []int {
	return []int{
		51, 50, 48, 47, 46, 44, 43, 42, 40, 39, 38, 36, 35, 34, 32, 31, 30, 28, 27, 26, 24,
		23, 22, 20, 19, 18, 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45,
		49, 16, 15, 14, 12, 11, 10, 8, 7, 6, 4, 3, 2, 0,
	}
}

func mockMaskToArray() [4][4]string {
	return [4][4]string{
		{"KQ", "954", "Q92", "A"},
		{"764", "76", "A6", "Q72"},
		{"", "", "", ""},
		{"AT8", "KT3", "KJ54", ""},
	}
}

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
		{"Test1", args{mockHand()}, "666777K888A99"},
		{"Test1", args{mockHandWithUndef()}, "6667?7K888A99"},
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
		{"Test1", args{mockHand(), 1}, mockSuitHand()},
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
		{"Test1", args{mockHand()}, mockSortHand()},
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
		{"Test1", args{mockHand()}, mockStringHand()},
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
		{"Test1", args{mockHand()}, 7},
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
		{"Test1", args{mockDeal()}, mockPbnSimple()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pbnDealSimple(tt.args.a); got != tt.want {
				t.Errorf("pbnDealSimple() = %v, want %v", got, tt.want)
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
		{"Test1", args{mockDeal()}, mockResultRandom()},
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

func TestDealMaskString(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"AK4.KJ.4.KT987 62.Q6.KJT8.A53 QT8..97532.4 753.T95.6.QJ62"},
			"AK4.AKJ.A4.KT987 J62.Q6.KQJT8.A53 QT98.874.97532.4 753.T9532.6.QJ62"},
	}
	var sh FakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DealMaskString(sh, tt.args.mask); got != tt.want {
				t.Errorf("DealMaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
