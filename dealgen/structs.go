package dealgen

const (
	N_CARDS    = 52
	N_HANDS    = 13
	N_OF_HANDS = 4
	N_OF_SUITS = 4
	NONE       = -1
	POINT      = "."
	SPACE      = " "
	POINTS3    = "..."
	MINUS      = "-"
	TEN        = "T"
	JOKER      = "J"
	QUEEN      = "Q"
	KING       = "K"
	ACE        = "A"
	TAB        = "\t"
	MAXPOINTS  = 37
)

type dataPoints struct {
	points int
	dist   []int
}

type (
	Random           struct{}
	ShuffleInterface interface {
		fYShuffle(int) []int
	}
	Shuffler interface {
		Shuffle([]int) []int
	}
)
