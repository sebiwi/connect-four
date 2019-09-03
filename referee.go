package main

const (
	ongoingMatch string = "Ongoing\n"
	yellowWins   string = "Yellow wins!\n"
	redWins      string = "Red wins!\n"
	tieMatch     string = "Tie!\n"
)

type referee struct {
	analyzer      analyzer
	grid          grid
	currentPlayer string
}

func newReferee(a analyzer, g grid) referee {
	return referee{analyzer: a, grid: g, currentPlayer: "o"}
}

func (r referee) getCurrentPlayer() string {
	if r.currentPlayer == "o" {
		return "Yellow"
	}
	return "Red"
}

func (r *referee) play(column int) {
	if r.grid.put(column, r.currentPlayer) {
		if r.status() == ongoingMatch {
			r.switchPlayer()
		}
	}
}

func (r *referee) switchPlayer() {
	if r.currentPlayer == "*" {
		r.currentPlayer = "o"
	} else {
		r.currentPlayer = "*"
	}
}

func (r referee) status() string {
	if r.analyzer.win() {
		if r.currentPlayer == "*" {
			return redWins
		}
		return yellowWins
	}
	if r.analyzer.tie() {
		return tieMatch
	}
	return ongoingMatch
}
