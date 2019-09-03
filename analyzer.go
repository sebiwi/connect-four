package main

import (
	"strings"
)

const (
	rightDirection int = 1
	upDirection    int = 1
	leftDirection  int = -1
)

type analyzer struct {
	grid grid
}

func newAnalyzer(g grid) analyzer {
	return analyzer{grid: g}
}

func (a analyzer) win() bool {
	for row := 0; row < gridHeight; row++ {
		for column := 0; column < gridWidth; column++ {
			if a.winnerAtPosition(column, row) {
				return true
			}
		}
	}
	return false
}

func (a analyzer) winnerAtPosition(column, row int) bool {
	return a.connectFour(column, row, rightDirection, 0) ||
		a.connectFour(column, row, 0, upDirection) ||
		a.connectFour(column, row, rightDirection, upDirection) ||
		a.connectFour(column, row, leftDirection, upDirection)
}

func (a analyzer) connectFour(column, row, horizontalDirection, verticalDirection int) (ret bool) {
	defer func() {
		if err := recover(); err != nil {
			ret = false
		}
	}()
	if a.grid.getPosition(column, row) != "." &&
		a.grid.getPosition(column, row) == a.grid.getPosition(column+1*horizontalDirection, row+1*verticalDirection) &&
		a.grid.getPosition(column, row) == a.grid.getPosition(column+2*horizontalDirection, row+2*verticalDirection) &&
		a.grid.getPosition(column, row) == a.grid.getPosition(column+3*horizontalDirection, row+3*verticalDirection) {
		return true
	}
	return false
}

func (a analyzer) tie() bool {
	return !a.win() && !strings.Contains(a.grid.render(), ".")
}
