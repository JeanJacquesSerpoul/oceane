package dealgen

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CardValueInt(tt.args.cardValue); got != tt.want {
				t.Errorf("CardValueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
