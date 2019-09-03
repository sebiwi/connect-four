package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type application struct {
	referee  referee
	analyzer analyzer
	grid     grid
	writer   io.Writer
	reader   bufio.Reader
}

func newApplication(r referee, a analyzer, g grid, wr io.Writer, rd io.Reader) application {
	return application{referee: r,
		analyzer: a,
		grid:     g,
		writer:   wr,
		reader:   *bufio.NewReader(rd)}
}

func (a *application) play() {
	a.write()
	for !a.analyzer.win() && !a.analyzer.tie() {
		a.read()
		a.write()
	}
}

func (a *application) read() {
	column := a.formatInput()
	a.referee.play(column)
}

func (a *application) write() {
	if a.analyzer.win() || a.analyzer.tie() {
		fmt.Fprint(a.writer, "\n\n\n", a.grid.render(), a.referee.status())
	} else {
		fmt.Fprint(a.writer, "\n\n\n", a.grid.render(), a.referee.getCurrentPlayer(), " column [1-7]: ")
	}
}

func (a *application) formatInput() int {
	column, _ := a.reader.ReadString('\n')
	column = strings.TrimSuffix(column, "\n")
	columnInt, _ := strconv.Atoi(column)
	columnInt--
	return columnInt
}
