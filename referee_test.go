package main

import (
	"testing"
)

func TestRefereeStartsWithYellow(t *testing.T) {
	// Given
	want := "Yellow"
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	got := referee.getCurrentPlayer()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestRefereeContinuesWithRed(t *testing.T) {
	// Given
	want := "Red"
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(1)
	got := referee.getCurrentPlayer()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestRefereeContinuesWithYellowOnThirdTurn(t *testing.T) {
	// Given
	want := "Yellow"
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(1)
	referee.play(1)
	got := referee.getCurrentPlayer()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestRefereeAddsYellowTokenOnFirstPlay(t *testing.T) {
	// Given
	want := ".......\n.......\n.......\n.......\n.......\no......\n"
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(0)
	got := grid.render()
	// Then
	if got != want {
		t.Errorf("got\n%v\nwant\n%v", got, want)
	}
}

func TestRefereeAddsRedTokenOnNextPlay(t *testing.T) {
	// Given
	want := ".......\n.......\n.......\n.......\n.......\no*.....\n"
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(0)
	referee.play(1)
	got := grid.render()
	// Then
	if got != want {
		t.Errorf("got\n%v\nwant\n%v", got, want)
	}
}

func TestRefereePlaysSameTokenAgainWhenColumnIsFull(t *testing.T) {
	// Given
	want := "*......\no......\n*......\no......\n*......\noo.....\n"
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(0)
	referee.play(0)
	referee.play(0)
	referee.play(0)
	referee.play(0)
	referee.play(0)
	// Rejected
	referee.play(0)
	// Accepted
	referee.play(1)
	got := grid.render()
	// Then
	if got != want {
		t.Errorf("got\n%v\nwant\n%v", got, want)
	}
}

func TestRefereeKnowsGameIsOngoing(t *testing.T) {
	// Given
	want := ongoingMatch
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	got := referee.status()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestRefereeKnowsYellowWon(t *testing.T) {
	// Given
	want := yellowWins
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(0)
	referee.play(0)
	referee.play(1)
	referee.play(1)
	referee.play(2)
	referee.play(2)
	referee.play(3)
	got := referee.status()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestRefereeKnowsRedWon(t *testing.T) {
	// Given
	want := redWins
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(6)
	referee.play(0)
	referee.play(0)
	referee.play(1)
	referee.play(1)
	referee.play(2)
	referee.play(2)
	referee.play(3)
	got := referee.status()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestRefereeKnowsTie(t *testing.T) {
	// Given
	want := tieMatch
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	// When
	referee.play(0)
	referee.play(1)
	referee.play(0)
	referee.play(1)
	referee.play(0)
	referee.play(1)

	referee.play(2)
	referee.play(3)
	referee.play(2)
	referee.play(3)
	referee.play(2)
	referee.play(3)

	referee.play(4)
	referee.play(5)
	referee.play(4)
	referee.play(5)
	referee.play(4)
	referee.play(5)

	referee.play(6)
	referee.play(0)
	referee.play(6)
	referee.play(0)
	referee.play(6)
	referee.play(0)

	referee.play(1)
	referee.play(2)
	referee.play(1)
	referee.play(2)
	referee.play(1)
	referee.play(2)

	referee.play(3)
	referee.play(4)
	referee.play(3)
	referee.play(4)
	referee.play(3)
	referee.play(4)

	referee.play(5)
	referee.play(6)
	referee.play(5)
	referee.play(6)
	referee.play(5)
	referee.play(6)

	got := referee.status()
	// Then
	if got != want {
		t.Errorf("got '%v', want '%v', given:\n%v", got, want, referee.grid.render())
	}
}
