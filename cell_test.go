package main_test

import (
	. "github.com/addisonhuddy/gonway"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cell", func() {
	It("I can be created", func() {
		var cell Cell
		Expect(cell).NotTo(Equal(nil))
	})

	It("Can be alive or dead", func() {
		var cell Cell
		cell.State = true
		Expect(cell.IsAlive()).To(Equal(true))

		cell.State = false
		Expect(cell.IsAlive()).To(Equal(false))
	})

	It("Can be brough to life", func() {
		var cell = NewCell()
		Expect(cell.IsAlive()).To(Equal(false))

		cell.Live()
		Expect(cell.IsAlive()).To(Equal(true))
	})

	It("Can be killed", func() {
		var cell = NewCell()
		cell.State = true
		cell.Kill()
		Expect(cell.IsAlive()).To(Equal(false))
	})

})
