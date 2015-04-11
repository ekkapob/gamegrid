package main

type Player struct {
	x int
	y int
}

func (p *Player) moveColumn(direction int) {
	p.x = p.x + direction
}

func (p *Player) moveRow(direction int) {
	p.y = p.y + direction
}
