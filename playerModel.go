package main

type Player struct {
	x     int
	y     int
	board Board
}

func (p *Player) moveColumn(direction int) {
	if direction > 0 && p.x < (p.board.column-1) {
		p.x++
	} else if direction < 0 && p.x > 0 {
		p.x--
	}
}

func (p *Player) moveRow(direction int) {
	if direction > 0 && p.y < (p.board.row-1) {
		p.y++
	} else if direction < 0 && p.y > 0 {
		p.y--
	}
}
