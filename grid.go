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
	return self.size * self.size
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

	if self.cells[position+1].IsAlive() && self.FarRight(position) != true {
		//right
		count++
	}

	if self.cells[position-1].IsAlive() && self.FarLeft(position) != true {
		//left
		count++
	}

	if self.cells[self.Below(position)].IsAlive() && self.LastRow(position) != true {
		//down
		count++
	}

	if self.cells[self.Above(position)].IsAlive() && self.TopRow(position) != true {
		//up
		count++
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
