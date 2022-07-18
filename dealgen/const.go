package dealgen

const (
	N_CARDS = 52
	N_HANDS = 13
	FOUR    = 4
)

var InitHand = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
}
var (
	faceCards  = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	suitCards  = []string{"S", "H", "D", "C"}
	position   = []string{"N", "E", "S", "W"}
	valueCards = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4}
	vulnerable = []string{"ALL", "NS", "EW", "NONE"}
)