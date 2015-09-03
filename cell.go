package main

type Cell struct {
	State bool
}

func NewCell() Cell {
	var c Cell
	c.State = false
	return c
}

func (self *Cell) IsAlive() bool {
	return self.State
}

func (self *Cell) Live() {
	self.State = true
}

func (self *Cell) Kill() {
	self.State = false
}
