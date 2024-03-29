package dealgen

import (
	"reflect"
	"testing"
)

type fakeRandom struct{}

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

func mockAuthSuit() []int {
	return []int{1, 2}
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

func mockMaskSuitToArray() [][]int {
	r := [][]int{{5, 4, 4, 1}, {8, 7, 6, 2}, {1, 3, 6, 0}, {1, 2, 3, 1}}
	return r
}

func mockMaskConvertToArray() [][]string {
	a := make([][]string, 4)
	for i := range a {
		a[i] = make([]string, 4)
	}
	return a
}

func mockSuit() []int {
	return []int{48, 44, 40, 36, 32, 28, 24, 20, 16, 12, 8, 4, 0}
}

func (test fakeRandom) fYShuffle(n int) []int {
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
		{"Test2", args{19}, 3},
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
		{"Test2", args{mockHandWithUndef()}, "66677K888A99"},
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
			if got := pointsFromDeal(tt.args.h); got != tt.want {
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
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dealRandom(sh, tt.args.a); !reflect.DeepEqual(got, tt.want) {
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
		{"Test1", args{100}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFaceCard(tt.args.v); got != tt.want {
				t.Errorf("getFaceCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskConvertToArray(t *testing.T) {
	type args struct {
		pbn string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Test1", args{"AK4.KJ.4....KT987 62.Q6.KJT8.A53 QT8..97532.4 753.T95.6.QJ62"}, mockMaskConvertToArray()},
		{"Test2", args{"AK4.KJ.4.KT987 62.Q6.KJ T8.A53 QT8..97532.4 753.T95.6.QJ62"}, mockMaskConvertToArray()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskConvertToArray(tt.args.pbn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maskConvertToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskSuitToArray(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Test1", args{"5.4.4.1 8.7.6.2 1.3.6.0 1.2.3.1"}, mockMaskSuitToArray()},
		{"Test2", args{"5.4.4.18.7.6.2 1.3.6.0 1.2.3.1"}, nullMaskSuitToArray()},
		{"Test3", args{"5.4.4.1 8.76.2 1.3.6.0 1.2.3.1"}, nullMaskSuitToArray()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskSuitToArray(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maskSuitToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomSuitArray(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{0}, mockSuit()},
		{"Test2", args{100}, mockSuit()},
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomSuitArray(sh, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomSuitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractFromRandom(t *testing.T) {
	type args struct {
		authSuit []int
		sk       []int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{mockAuthSuit(), nil, 5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractFromRandom(tt.args.authSuit, tt.args.sk, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFromRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDealPointsString(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"16.5.0."}, ""},
		{"Test2", args{"5.3.0."}, "T9.JT9.JT9.QJT98 J543.543.Q43.432 762.762.7652.765 AKQ8.AKQ8.AK8.AK"},
		{"Test3", args{"5.3.A."}, "T9.JT9.JT9.QJT98 J543.543.Q43.432 AKQ2.AKQ2.AK2.AK 876.876.8765.765"},
		{"Test4", args{"20..17.0"}, "QJ.QJT.AQJT.AQJT K432.432.432.432 A65.AK65.K65.K65 T987.987.987.987"},
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DealPointsString(sh, tt.args.mask)
			if got != tt.want {
				t.Errorf("DealPointsString() = %v, want %v", got, tt.want)
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
		{
			"Test1",
			args{"AK4.KJ.4.KT987 62.Q6.KJT8.A53 QT8..97532.4 753.T95.6.QJ62"},
			"AK4.AKJ.A4.KT987 J62.Q6.KQJT8.A53 QT98.874.97532.4 753.T9532.6.QJ62",
		},
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DealMaskString(sh, tt.args.mask)
			if got != tt.want {
				t.Errorf("DealMaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDealSuitString(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"1.0.3. ... ... ..."}, "A..AKQ.T98765432 ..T98765432.AKQJ .KQJT98765432.J. KQJT98765432.A.."},
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DealSuitString(sh, tt.args.mask)
			if got != tt.want {
				t.Errorf("DealSuitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
