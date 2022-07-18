package dealgen

type Shuffler interface {
	Shuffle([]int) []int
}
type CardList struct {
	CardList []int
}
