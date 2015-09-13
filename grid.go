package main

import (
	"math/rand"
	"time"
)

type Grid interface {
	Size() int
	NumerOfCells() int
	GetCell(int) Cell
	Neighbors(int) int
}

type grid struct {
	size  int
	cells []Cell
}

func NewGrid(size int) *grid {
	rand.Seed(time.Now().UTC().UnixNano())
	g := new(grid)
	g.size = size

	// fill the grid with dead cells
	for i := 0; i < (size * size); i++ {
		g.cells = append(g.cells, NewCell())
	}

	return g
}

func CloneGrid(g *grid) *grid {
	newGrid := new(grid)
	newGrid.size = g.size

	for i := range g.cells {
		newGrid.cells = append(newGrid.cells, NewCell())
		if g.GetCell(i).IsAlive() {
			newGrid.GetCell(i).Live()
		}
	}

	return newGrid
}

func (self *grid) Size() int {
	return self.size
}

func (self *grid) NumerOfCells() int {
	return self.size * self.size
}

func (self *grid) GetCell(position int) *Cell {
	cell := &self.cells[position]
	return cell
}

func (self *grid) PrintGrid() string {
	var s string

	for i := range self.cells {

		cell := self.GetCell(i)

		if cell.IsAlive() == true {
			s = s + "o"
		} else {
			s = s + " "
		}

		if self.FarRight(i) {
			s = s + "\n"
		}
	}
	return s
}

func (self *grid) Neighbors(position int) int {
	count := 0

	//right
	if !self.FarRight(position) {
		if self.cells[position+1].IsAlive() {
			count++
		}
	}

	//left
	if !self.FarLeft(position) {
		if self.cells[position-1].IsAlive() {
			count++
		}
	}

	//down
	if !self.LastRow(position) {
		if self.cells[self.Below(position)].IsAlive() {
			count++
		}

		// bottom right
		if !self.FarRight(position) {
			if self.cells[self.Below(position+1)].IsAlive() {
				count++
			}
		}

		// bottom left
		if !self.FarLeft(position) {
			if self.cells[self.Below(position-1)].IsAlive() {
				count++
			}
		}
	}

	//up
	if !self.TopRow(position) {
		if self.cells[self.Above(position)].IsAlive() {
			count++
		}

		// upper right
		if !self.FarRight(position) {
			if self.cells[self.Above(position+1)].IsAlive() {
				count++
			}
		}

		//upper left
		if !self.FarLeft(position) {
			if self.cells[self.Above(position-1)].IsAlive() {
				count++
			}
		}
	}
	return count
}

func (self *grid) Below(index int) int {
	return index + self.size
}

func (self *grid) Above(index int) int {
	return index - self.size
}

func (self *grid) TopRow(index int) bool {
	return index < self.size
}

func (self *grid) LastRow(index int) bool {
	return index >= self.NumerOfCells()-self.size
}

func (self *grid) FarRight(index int) bool {
	return (index+1)%self.size == 0
}

func (self *grid) FarLeft(index int) bool {
	return index%self.size == 0
}

func Bang(grid *grid) *grid {

	tempGrid := CloneGrid(grid)

	for i := range grid.cells {
		cell := tempGrid.GetCell(i)

		// Rule 1
		if grid.Neighbors(i) < 2 {
			cell.Kill()
		}

		// Rule 2
		if grid.Neighbors(i) == 2 || grid.Neighbors(i) == 3 {
			if cell.IsAlive() {
				continue
			}
		}

		// Rule 3
		if grid.Neighbors(i) > 3 {
			cell.Kill()
		}

		// Rule 4
		if grid.Neighbors(i) == 3 {
			if !cell.IsAlive() {
				cell.Live()
			}
		}
	}

	return tempGrid
}

func (self *grid) Randomize() {
	for i := range self.cells {
		randomNumber := randInt(0, 100)
		print(randomNumber)
		cell := self.GetCell(i)
		if randomNumber > 75 {
			cell.Live()
		}
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func ClearScreen() {
	print("\033[H\033[2J")
}
