package main

const (
	gridHeight int = 6
	gridWidth  int = 7
)

type grid struct {
	columns [][]string
}

func newGrid() grid {
	grid := grid{}
	grid.init()
	return grid
}

func (g *grid) init() {
	columns := make([][]string, gridWidth)
	for column := 0; column < gridWidth; column++ {
		columns[column] = make([]string, gridHeight)
	}
	g.columns = columns
	g.empty()
}

func (g *grid) empty() {
	for column := 0; column < gridWidth; column++ {
		for row := 0; row < gridHeight; row++ {
			g.columns[column][row] = "."
		}
	}
}

func (g grid) render() (gridString string) {
	for row := gridHeight - 1; row >= 0; row-- {
		for column := 0; column < gridWidth; column++ {
			gridString += g.columns[column][row]
		}
		gridString += "\n"
	}
	return gridString
}

func (g *grid) put(column int, token string) bool {
	if column < 0 || column >= 7 {
		return false
	}
	for row := 0; row < gridHeight; row++ {
		if g.columns[column][row] == "." {
			g.columns[column][row] = token
			return true
		}
	}
	return false
}

func (g grid) getPosition(column, row int) string {
	return g.columns[column][row]
}
