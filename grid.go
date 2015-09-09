package main

import (
//"fmt"
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
	g := new(grid)
	g.size = size

	// fill the grid with dead cells
	for i := 0; i < (size * size); i++ {
		g.cells = append(g.cells, NewCell())
	}

	return g
}

func (self *grid) Size() int {
	return self.size
}

func (self *grid) NumerOfCells() int {
	return len(self.cells)
}

func (self *grid) GetCell(position int) *Cell {
	cell := &self.cells[position]
	return cell
}

func (self *grid) PrintGrid() string {
	var s string
	for index, element := range self.cells {

		if element.IsAlive() == true {
			s = s + "o"
		} else {
			s = s + " "
		}
		//TODO make this more readable
		if self.NumerOfCells()%(index+1) == 0 && index != 0 {
			s = s + "\n"
		}
	}
	return s
}

func (self *grid) Neighbors(position int) int {
	count := 0

	if self.cells[position+1].IsAlive() {
		//right
		count++
	} else if self.cells[position-1].IsAlive() {
		//left
		count++
	} else if self.cells[position-1].IsAlive() {
		//down
		count++
	}
	return count
}

func (self *grid) Below(index int) int {
	return -1
}
