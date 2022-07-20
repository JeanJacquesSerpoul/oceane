package dealgen

const (
	N_CARDS  = 52
	N_HANDS  = 13
	UNDEF    = "?"
	ERRORMSG = "ERROR!"
	POINT    = "."
	SPACE    = " "
	POINTS3  = "..."
	MINUS    = "-"
)

var initDeal = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
}

var valueCards = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4}

type (
	Random           struct{}
	ShuffleInterface interface {
		fYShuffle(int) []int
	}
)

type Shuffler interface {
	Shuffle([]int) []int
}
