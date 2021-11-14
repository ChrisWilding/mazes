package mazes

func NewDistances(root *Cell) *Distances {
	d := &Distances{
		root:  root,
		cells: make(map[*Cell]int),
	}
	d.cells[root] = 0
	return d
}

type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

func (d *Distances) GetCells() []*Cell {
	keys := make([]*Cell, 0, len(d.cells))
	for c := range d.cells {
		keys = append(keys, c)
	}
	return keys
}

func (d *Distances) GetDistance(cell *Cell) (int, bool) {
	distance, ok := d.cells[cell]
	return distance, ok
}

func (d *Distances) SetDistance(cell *Cell, distance int) {
	d.cells[cell] = distance
}
