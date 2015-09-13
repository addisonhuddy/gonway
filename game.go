package main

import (
	"time"
)

func BigBang() {
	var grid = NewGrid(60)
	grid.Randomize()
	for {
		print(grid.PrintGrid())
		grid = Bang(grid)
		time.Sleep(200 * time.Millisecond)
	}
}
