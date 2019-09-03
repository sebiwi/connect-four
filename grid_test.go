// Package main provides a grid
package main

import (
	"testing"
)

func TestGridIsEmptyOnCreation(t *testing.T) {
	// Given
	want := ".......\n.......\n.......\n.......\n.......\n.......\n"
	grid := newGrid()
	// When
	got := grid.render()
	// Then
	if got != want {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestGridAcceptTokenOnFirstColumn(t *testing.T) {
	// Given
	want := ".......\n.......\n.......\n.......\n.......\n*......\n"
	grid := newGrid()
	// When
	putOutput := grid.put(0, "*")
	got := grid.render()
	// Then
	if putOutput != true {
		t.Errorf("got: %v, wanted: true", putOutput)
	}
	if got != want {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestGridAcceptTokenOnSecondColumn(t *testing.T) {
	// Given
	want := ".......\n.......\n.......\n.......\n.......\n.*.....\n"
	grid := newGrid()
	// When
	putOutput := grid.put(1, "*")
	got := grid.render()
	// Then
	if putOutput != true {
		t.Errorf("got: %v, wanted: true", putOutput)
	}
	if got != want {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestGridAcceptTwoTokensOnFirstColumn(t *testing.T) {
	// Given
	want := ".......\n.......\n.......\n.......\n*......\n*......\n"
	grid := newGrid()
	// When
	grid.put(0, "*")
	putOutput := grid.put(0, "*")
	got := grid.render()
	// Then
	if putOutput != true {
		t.Errorf("got: %v, wanted: true", putOutput)
	}
	if got != want {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestGridRejectsTokenOnFullColumn(t *testing.T) {
	// Given
	grid := newGrid()
	// When
	grid.put(0, "*")
	grid.put(0, "*")
	grid.put(0, "*")
	grid.put(0, "*")
	grid.put(0, "*")
	grid.put(0, "*")
	putOutput := grid.put(0, "*")
	// Then
	if putOutput != false {
		t.Errorf("got: %v, wanted: false", putOutput)
	}
}

func TestGridRejectsTokenOutsideOnNegativeColumn(t *testing.T) {
	// Given
	grid := newGrid()
	// When
	putOutput := grid.put(-1, "*")
	// Then
	if putOutput != false {
		t.Errorf("got: %v, wanted: false", putOutput)
	}
}

func TestGridRejectsTokenOutsideOfGridSize(t *testing.T) {
	// Given
	grid := newGrid()
	// When
	putOutput := grid.put(7, "*")
	// Then
	if putOutput != false {
		t.Errorf("got: %v, wanted: false", putOutput)
	}
}

func TestGridReturnStateOfPosition(t *testing.T) {
	// Given
	grid := newGrid()
	grid.put(0, "*")
	// When
	positionOutput := grid.getPosition(0, 0)
	// Then
	if positionOutput != "*" {
		t.Errorf("got: %v, wanted: false", positionOutput)
	}
}

func TestGridReturnStateOfPositionSecondColumn(t *testing.T) {
	// Given
	grid := newGrid()
	grid.put(1, "*")
	// When
	positionOutput := grid.getPosition(1, 0)
	// Then
	if positionOutput != "*" {
		t.Errorf("got: %v, wanted: false", positionOutput)
	}
}

func TestGridReturnStateOfPositionSecondRow(t *testing.T) {
	// Given
	grid := newGrid()
	grid.put(0, "*")
	grid.put(0, "*")
	// When
	positionOutput := grid.getPosition(0, 1)
	// Then
	if positionOutput != "*" {
		t.Errorf("got: %v, wanted: false", positionOutput)
	}
}
