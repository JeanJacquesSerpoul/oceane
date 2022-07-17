package dealgen

type Hand [N_CARDS_IN_HAND]int
type CardGames [N_CARDS]int
type Desk struct {
	Deal CardGames
}
type Shuffler interface {
	Shuffle() []int
}
