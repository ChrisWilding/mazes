package mazes

import (
	"bytes"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Grid struct {
	Rows    int
	Columns int

	Grid [][]*Cell
}

func NewGrid(rows, columns int) *Grid {
	grid := Grid{
		Rows:    rows,
		Columns: columns,

		Grid: make([][]*Cell, rows),
	}

	for row := range grid.Grid {
		grid.Grid[row] = make([]*Cell, columns)
		for column := 0; column < columns; column++ {
			grid.Grid[row][column] = NewCell(row, column)
		}
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			cell := grid.Grid[row][column]
			cell.North = grid.Cell(row-1, column)
			cell.South = grid.Cell(row+1, column)
			cell.West = grid.Cell(row, column-1)
			cell.East = grid.Cell(row, column+1)
		}
	}

	return &grid
}

func (g *Grid) Cell(row, column int) *Cell {
	if row < 0 || row >= g.Rows {
		return nil
	}
	if column < 0 || column >= g.Columns {
		return nil
	}
	return g.Grid[row][column]
}

func (g *Grid) RandomCell() *Cell {
	row := random.Intn(g.Rows - 1)
	column := random.Intn(g.Columns - 1)
	return g.Grid[row][column]
}

func (g *Grid) Size() int {
	return g.Rows * g.Columns
}

func (g *Grid) String() string {
	var b bytes.Buffer

	b.WriteString("+")
	for i := 0; i < g.Columns; i++ {
		b.WriteString("---+")
	}
	b.WriteString("\n")

	for row := 0; row < g.Rows; row++ {
		var top bytes.Buffer
		top.WriteString("|")

		var bottom bytes.Buffer
		bottom.WriteString("+")

		for column := 0; column < g.Columns; column++ {
			cell := g.Cell(row, column)
			if cell == nil {
				cell = NewCell(-1, -1)
			}

			top.WriteString("   ")

			if cell.IsLinked(cell.East) {
				top.WriteString(" ")
			} else {
				top.WriteString("|")
			}

			if cell.IsLinked(cell.South) {
				bottom.WriteString("   +")
			} else {
				bottom.WriteString("---+")
			}
		}

		top.WriteTo(&b)
		b.WriteString("\n")

		bottom.WriteTo(&b)
		b.WriteString("\n")
	}

	return b.String()
}
