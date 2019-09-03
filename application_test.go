package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestApplicationRendersEmptyGridAndYellowPrompt(t *testing.T) {
	// Given
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	writer := &bytes.Buffer{}
	reader := bufio.NewReader(strings.NewReader("1\n1\n2\n2\n3\n3\n4\n"))
	application := newApplication(referee, analyzer, grid, writer, reader)
	want := "\n\n\n.......\n.......\n.......\n.......\n.......\n.......\nYellow column [1-7]: \n"
	// When
	application.play()
	got := writer.String()
	// Then
	if !strings.Contains(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestApplicationTakesInputUntilYellowWins(t *testing.T) {
	// Given
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	writer := &bytes.Buffer{}
	reader := bufio.NewReader(strings.NewReader("1\n1\n2\n2\n3\n3\n4\n"))
	application := newApplication(referee, analyzer, grid, writer, reader)
	want := "\n\n\n.......\n.......\n.......\n.......\n***....\noooo...\nYellow wins!\n"
	// When
	application.play()
	got := writer.String()
	// Then
	if !strings.Contains(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestApplicationTakesInputUntilRedWins(t *testing.T) {
	// Given
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	writer := &bytes.Buffer{}
	reader := bufio.NewReader(strings.NewReader("5\n1\n1\n2\n2\n3\n3\n4\n"))
	application := newApplication(referee, analyzer, grid, writer, reader)
	want := "\n\n\n.......\n.......\n.......\n.......\nooo....\n****o..\nRed wins!\n"
	// When
	application.play()
	got := writer.String()
	// Then
	if !strings.Contains(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}

func TestApplicationTakesInputUntilTie(t *testing.T) {
	// Given
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	writer := &bytes.Buffer{}
	reader := bufio.NewReader(strings.NewReader("1\n2\n1\n2\n1\n2\n3\n4\n3\n4\n3\n4\n5\n6\n5\n6\n5\n6\n7\n1\n7\n1\n7\n1\n2\n3\n2\n3\n2\n3\n4\n5\n4\n5\n4\n5\n6\n7\n6\n7\n6\n7\n"))
	application := newApplication(referee, analyzer, grid, writer, reader)
	want := "\n\n\n*o*o*o*\n*o*o*o*\n*o*o*o*\no*o*o*o\no*o*o*o\no*o*o*o\nTie!\n"
	// When
	application.play()
	got := writer.String()
	// Then
	if !strings.Contains(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v", got, want)
	}
}
