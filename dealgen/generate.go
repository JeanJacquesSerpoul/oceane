package dealgen

func (d *Desk) Shuffle() CardGames {
	for i, _ := range d.Deal {
		d.Deal[i] = 1
	}
	return d.Deal
}
func (d *Desk) Init() {
	d.Deal = initHand
}
