package main_test

import (
	//"fmt"
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

	Describe("Counting the neightbors of a cell", func() {
		It("Returns the number of neighbors a cell has", func() {
			var grid = NewGrid(3)

			//oxo  012
			//oxx  345
			//oxo  678

			grid.GetCell(1).Live()
			grid.GetCell(4).Live()
			grid.GetCell(7).Live()
			grid.GetCell(5).Live()
			Expect(grid.Neighbors(4)).To(Equal(3))
		})

		It("Knows where the edges of the grid are", func() {
			var grid = NewGrid(3)
			grid.GetCell(0).Live()
			grid.GetCell(3).Live()
			grid.GetCell(6).Live()
			grid.GetCell(2).Live()
			Expect(grid.Neighbors(3)).To(Equal(2))
		})

		It("Gets the neighbors that are diagonal", func() {
			var grid = NewGrid(3)

			//xox  012
			//ooo  345
			//xox  678

			grid.GetCell(0).Live()
			grid.GetCell(2).Live()
			grid.GetCell(6).Live()
			grid.GetCell(8).Live()
			Expect(grid.Neighbors(4)).To(Equal(4))
		})
	})
	Describe("Cells live or die given Conway's Rules", func() {
		It("Rule 1: Kills a cell that has less than 2 neighbors", func() {
			var grid = NewGrid(3)

			//oxx
			//oxx
			//oxx

			grid.GetCell(0).Live()
			grid.GetCell(3).Live()
			grid.GetCell(6).Live()
			grid = Bang(grid)
			Expect(grid.GetCell(0).IsAlive()).To(Equal(false))
			Expect(grid.GetCell(3).IsAlive()).To(Equal(true))
			Expect(grid.GetCell(6).IsAlive()).To(Equal(false))
		})

		It("Rule #2: Cells with 2 or 3 neighbors live", func() {
			var grid = NewGrid(3)

			// xxo
			// xoo
			// ooo
			grid.GetCell(0).Live()
			grid.GetCell(1).Live()
			grid.GetCell(3).Live()
			grid = Bang(grid)
			Expect(grid.GetCell(0).IsAlive()).To(Equal(true))
		})

		It("Rule #3: Cells with more than 3 neighbors die due to overpopulation", func() {
			var grid = NewGrid(3)

			// xxx
			// xxo
			// ooo

			grid.GetCell(0).Live()
			grid.GetCell(1).Live()
			grid.GetCell(2).Live()
			grid.GetCell(3).Live()
			grid.GetCell(4).Live()
			grid = Bang(grid)
			Expect(grid.GetCell(4).IsAlive()).To(Equal(false))
		})

		It("Rule #4: Dead cells with exactly 3 neighbors are brought back to life", func() {
			var grid = NewGrid(3)

			// ooo
			// xxx
			// ooo

			grid.GetCell(3).Live()
			grid.GetCell(4).Live()
			grid.GetCell(5).Live()
			grid = Bang(grid)
			Expect(grid.GetCell(1).IsAlive()).To(Equal(true))

		})
	})
})
