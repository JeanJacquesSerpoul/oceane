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

func mockMaskPbn() string {
	v := `[Dealer "N"]
[Vulnerable "ALL"]
[Deal "N:AK4.AKJ.A4.KT987 J62.Q6.KQJT8.A53 QT98.874.97532.4 753.T9532.6.QJ62"]

`
	return v
}

func TestMultiPbnDeal(t *testing.T) {
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
	}
	var sh fakeRandom
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiPbnDeal(sh, tt.args.mode, tt.args.ite, tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.mask); got != tt.want {
				t.Errorf("MultiPbnDeal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPbnDealToFile(t *testing.T) {
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
			if err := MultiPbnDealToFile(sh, tt.args.filename, tt.args.mode, tt.args.ite, tt.args.firstHand, tt.args.dealer, tt.args.vul, tt.args.mask); (err != nil) != tt.wantErr {
				t.Errorf("MultiPbnDealToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
