package dealgen

const (
	N_CARDS       = 52
	N_HANDS       = 13
	N_OF_HANDS    = 4
	N_OF_SUITS    = 4
	NONE          = -1
	POINT         = "."
	SPACE         = " "
	POINTS3       = "..."
	MINUS         = "-"
	TEN           = "T"
	JOKER         = "J"
	QUEEN         = "Q"
	KING          = "K"
	ACE           = "A"
	TAB           = "\t"
	MAXPOINTSHAND = 37
	MAXPOINTSDEAL = 40
	MAXTRY        = 10
	INFINITE      = 100
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
)
