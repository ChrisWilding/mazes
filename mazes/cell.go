package mazes

func NewCell(row, column int) *Cell {
	return &Cell{
		Row:    row,
		Column: column,
		links:  make(map[*Cell]bool),
	}
}

type Cell struct {
	links map[*Cell]bool

	North *Cell
	South *Cell
	East  *Cell
	West  *Cell

	Column int
	Row    int
}

func (c *Cell) Link(cell *Cell) {
	c.links[cell] = true
}

func (c *Cell) LinkBidirectional(cell *Cell) {
	c.Link(cell)
	cell.Link(c)
}

func (c *Cell) Unlink(cell *Cell) {
	delete(c.links, cell)
}

func (c *Cell) UnlinkBidirectional(cell *Cell) {
	c.Unlink(cell)
	cell.Unlink(c)
}

func (c *Cell) Links() []*Cell {
	cells := make([]*Cell, 0, len(c.links))
	for cell := range c.links {
		cells = append(cells, cell)
	}
	return cells
}

func (c *Cell) IsLinked(cell *Cell) bool {
	_, ok := c.links[cell]
	return ok
}

func (c *Cell) Neighbours() []*Cell {
	cells := make([]*Cell, 0, len(c.links))

	if c.North != nil {
		cells = append(cells, c.North)
	}

	if c.South != nil {
		cells = append(cells, c.South)
	}

	if c.East != nil {
		cells = append(cells, c.East)
	}

	if c.West != nil {
		cells = append(cells, c.West)
	}

	return cells
}
