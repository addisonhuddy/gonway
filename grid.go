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

func (self *grid) Bang() {
	for i := range self.cells {
		cell := self.GetCell(i)
		if self.Neighbors(i) < 2 {
			cell.Kill()
		}

		if self.Neighbors(i) == 2 || self.Neighbors(i) == 3 {
			if cell.IsAlive() {
				continue
			}
		}

		if self.Neighbors(i) > 3 {
			cell.Kill()
		}

		if self.Neighbors(i) == 2 {
			if !cell.IsAlive() {
				cell.Live()
			}
		}
	}
}
