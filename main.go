package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to connect-four! \\o/")
	grid := newGrid()
	analyzer := newAnalyzer(grid)
	referee := newReferee(analyzer, grid)
	application := newApplication(referee, analyzer, grid, os.Stdout, os.Stdin)
	application.play()
}
