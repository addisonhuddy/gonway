package main_test

import (
	. "github.com/addisonhuddy/gonway"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {
	It("can accept an initial gride size", func() {
		grid := NewGrid(3)
		Expect(grid.Length()).To(Equal(3))
	})

	It("returns a cell at a given position", func() {
		grid := NewGrid(9)
		cell := grid.GetCell(5)
		Expect(cell.IsAlive()).To(Equal(false))
	})
})
