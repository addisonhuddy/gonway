package main

import (
//"fmt"
)

type Grid interface {
	Length() int
	GetCell(int) Cell
}

type grid struct {
	size        int
	gridOfCells []Cell
}

func NewGrid(size int) Grid {
	g := new(grid)
	g.size = size

	for i := 0; i < size; i++ {
		g.gridOfCells = append(g.gridOfCells, NewCell())
	}

	return g
}

func (g grid) Length() int {
	return g.size
}

func (g grid) GetCell(position int) Cell {
	cell := g.gridOfCells[position]
	return cell
}
