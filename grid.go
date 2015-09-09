package main

import (
//"fmt"
)

type Grid interface {
	Length() int
	GetCell(int) Cell
}

type grid struct {
	cells []Cell
}

func NewGrid(size int) *grid {
	g := new(grid)

	// fill the grid with dead cells
	for i := 0; i < size; i++ {
		g.cells = append(g.cells, NewCell())
	}

	return g
}

func (self *grid) Length() int {
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
		if self.Length()%(index+1) == 0 && index != 0 {
			s = s + "\n"
		}
	}
	return s
}
