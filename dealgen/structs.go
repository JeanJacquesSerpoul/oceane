package dealgen

type Shuffler interface {
	Shuffle([]int) []int
}

type CardList struct {
	CardList []int
}

type result struct {
	PbnSimple  string            `json:"pbnsimple"`
	Pbn        string            `json:"pbn"`
	HandPoints [FOUR]int         `json:"handpoints"`
	Suit       [FOUR][FOUR][]int `json:"suit"`
}
