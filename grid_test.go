package main_test

import (
	. "github.com/addisonhuddy/gonway"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {
	It("can accept an initial gride size", func() {
		grid := NewGrid(3)
		Expect(grid.Size()).To(Equal(3))
	})

	It("can return the number of cells in the grid", func() {
		grid := NewGrid(3)
		Expect(grid.NumerOfCells()).To(Equal(9))
	})

	It("returns a cell at a given position", func() {
		grid := NewGrid(9)
		cell := grid.GetCell(5)
		cell.Live()
		Expect(cell.IsAlive()).To(Equal(true))
	})

	It("Prints the grid", func() {
		var grid = NewGrid(2)
		cell := grid.GetCell(2)
		cell.Live()
		expectedGrid := "  \no \n"
		Expect(grid.PrintGrid()).To(Equal(expectedGrid))
	})

	It("Returns the number of neighbors a cell has", func() {
		var grid = NewGrid(9)
		grid.GetCell(1).Live()
		grid.GetCell(2).Live()
		grid.GetCell(4).Live()
		Expect(grid.Neighbors(2)).To(Equal(2))
	})
})
