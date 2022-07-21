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
)

type (
	Random           struct{}
	ShuffleInterface interface {
		fYShuffle(int) []int
	}
)

type Shuffler interface {
	Shuffle([]int) []int
}
