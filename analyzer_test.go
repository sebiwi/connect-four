package main

import (
	"testing"
)

// .......
// .......
// .......
// .......
// .......
// ****...
func TestAnalyzerDetectsWinOn4HorizontalFirstColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .......
// .......
// .......
// .***...
func TestAnalyzerDoesntDetectWinOn3HorizontalFirstColumnFirstRow(t *testing.T) {
	// Given
	wanted := false
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .......
// .......
// .......
// .****..
func TestAnalyzerDetectsWinOn4HorizontalSecondColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(4, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .......
// .......
// .......
// .......
func TestAnalyzerDoesntDetectWinOnEmptyGrid(t *testing.T) {
	// Given
	wanted := false
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .......
// .......
// .......
// .****..
func TestAnalyzerDoesntDetectWinOn4DifferentHorizontalFirstColumnFirstRow(t *testing.T) {
	// Given
	wanted := false
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(3, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .......
// .......
// ****...
// *o*o...
func TestAnalyzerDetectsWinOn4HoritontalFirstColumnSecondRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// *......
// *......
// *......
// *......
func TestAnalyzerDetectsWinOn4VerticalFirstColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// *......
// *......
// *......
// *......
// o......
func TestAnalyzerDetectsWinOn4VerticalFirstColumnSecondRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "o")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .*.....
// .*.....
// .*.....
// .*.....
func TestAnalyzerDetectsWinOn4VerticalSecondColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// ...*...
// ..*o...
// .*oo...
// *ooo...
func TestAnalyzerDetectsWinOn4DiagonalFirstColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(3, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// ....*..
// ...*o..
// ..*oo..
// .*ooo..
func TestAnalyzerDetectsWinOn4DiagonalSecondColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(4, "o")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(4, "o")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(4, "o")
	analyzer.grid.put(4, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// ......*
// .....*o
// ....*oo
// ...*ooo
func TestAnalyzerDetectsWinOn4DiagonalFourthColumnFirstRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(3, "*")
	analyzer.grid.put(4, "o")
	analyzer.grid.put(5, "o")
	analyzer.grid.put(6, "o")
	analyzer.grid.put(4, "*")
	analyzer.grid.put(5, "o")
	analyzer.grid.put(6, "o")
	analyzer.grid.put(5, "*")
	analyzer.grid.put(6, "o")
	analyzer.grid.put(6, "*")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// ...o...
// ..o*...
// .oo*...
// o*o*...
// *o*o...
// o*o*...
func TestAnalyzerDetectsWinOn4DiagonalFirstColumnThirdRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "o")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "o")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(3, "o")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// o......
// *o.....
// o*o....
// *o*o...
func TestAnalyzerDetectsWinOn4ReverseDiagonalFirstColumnFourthRow(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "o")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "o")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(3, "o")
	got := analyzer.win()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// .......
// .......
// .......
// .......
// .......
// ****...
func TestAnalyzerDoesntDetectTie(t *testing.T) {
	// Given
	wanted := false
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(3, "*")
	got := analyzer.tie()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}

// *o*o*o*
// o*o*o*o
// *o*o*o*
// *o*o*o*
// o*o*o*o
// *o*o*o*
func TestAnalyzerDetectsTie(t *testing.T) {
	// Given
	wanted := true
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	// When
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "o")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "*")
	analyzer.grid.put(0, "o")
	analyzer.grid.put(0, "*")

	analyzer.grid.put(1, "o")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(1, "o")
	analyzer.grid.put(1, "*")
	analyzer.grid.put(1, "o")

	analyzer.grid.put(2, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(2, "*")
	analyzer.grid.put(2, "o")
	analyzer.grid.put(2, "*")

	analyzer.grid.put(3, "o")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(3, "o")
	analyzer.grid.put(3, "*")
	analyzer.grid.put(3, "o")

	analyzer.grid.put(4, "*")
	analyzer.grid.put(4, "o")
	analyzer.grid.put(4, "*")
	analyzer.grid.put(4, "*")
	analyzer.grid.put(4, "o")
	analyzer.grid.put(4, "*")

	analyzer.grid.put(5, "o")
	analyzer.grid.put(5, "*")
	analyzer.grid.put(5, "o")
	analyzer.grid.put(5, "o")
	analyzer.grid.put(5, "*")
	analyzer.grid.put(5, "o")

	analyzer.grid.put(6, "*")
	analyzer.grid.put(6, "o")
	analyzer.grid.put(6, "*")
	analyzer.grid.put(6, "*")
	analyzer.grid.put(6, "o")
	analyzer.grid.put(6, "*")

	got := analyzer.tie()
	// Then
	if got != wanted {
		t.Errorf("got: %v, wanted: %v, given:\n%v", got, wanted, analyzer.grid.render())
	}
}
