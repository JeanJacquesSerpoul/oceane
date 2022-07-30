package dealgen

import (
	"testing"
)

func mockPbn() string {
	v := `[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:AKQJT9.AKQJ.A.AK ..32.QJT98765432 .432.KQJT987654. 8765432.T98765.."]

`
	return v
}

func mockPbnPoint() string {
	v := `[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:QJ.QJT.AQJT.AQJT K432.432.432.432 A65.AK65.K65.K65 T987.987.987.987"]

`
	return v
}

func mockMaskPbn() string {
	v := `[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:AK4.AKJ.A4.KT987 J62.Q6.KQJT8.A53 QT98.874.97532.4 753.T9532.6.QJ62"]

`
	return v
}

func mockArrayCards() []int {
	return []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
		36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
}

func mockArrayAll() []int {
	return []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 33, 34, 35,
		36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
}

func TestPbnDeal(t *testing.T) {
	type args struct {
		mode      int
		ite       int
		firstHand int
		dealer    int
		vul       int
		mask      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{1, 1, 0, 0, 0, "6.4.1.2 ... ... ..."}, mockPbn()},
		{
			"Test2",
			args{0, 1, 0, 0, 0, "AK4.KJ.4.KT987 62.Q6.KJT8.A53 QT8..97532.4 753.T95.6.QJ62"},
			mockMaskPbn(),
		},
		{"Test3", args{1, 1, 8, 0, 0, "6.4.1.2 ... ... ..."}, mockPbn()},
		{"Test4", args{1, 1, 0, 0, 8, "6.4.1.2 ... ... ..."}, mockPbn()},
		{"Test5", args{2, 1, 0, 0, 0, "20..17.0"}, mockPbnPoint()},
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PbnDeal(sh, tt.args.mode, tt.args.ite, tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.mask); got != tt.want {
				t.Errorf("PbnDeal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbnDealToFile(t *testing.T) {
	type args struct {
		filename  string
		mode      int
		ite       int
		firstHand int
		dealer    int
		vul       int
		mask      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test1", args{"test.pbn", 1, 1, 0, 0, 0, "6.4.1.2 ... ... ..."}, false},
		{"Test2", args{"aaaa/test.pbn", 1, 1, 0, 0, 0, "6.4.1.2 ... ... ..."}, true},
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PbnDealToFile(sh, tt.args.filename, tt.args.mode, tt.args.ite, tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.mask); (err != nil) != tt.wantErr {
				t.Errorf("PbnDealToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkDeal(t *testing.T) {
	type args struct {
		deal []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{mockArrayCards()}, 1},
		{"Test2", args{mockArrayAll()}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDeal(tt.args.deal); got != tt.want {
				t.Errorf("checkDeal() = %v, want %v", got, tt.want)
			}
		})
	}
}
