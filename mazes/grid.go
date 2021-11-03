package mazes

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math/rand"
	"time"
)

const cellSize = 10

type Grid struct {
	Random  *rand.Rand
	Grid    [][]*Cell
	Rows    int
	Columns int
}

func NewGrid(rows, columns int) *Grid {
	grid := Grid{
		Rows:    rows,
		Columns: columns,
		Grid:    make([][]*Cell, rows),
		Random:  rand.New(rand.NewSource(time.Now().UnixNano())),
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
	row := g.Random.Intn(g.Rows - 1)
	column := g.Random.Intn(g.Columns - 1)
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

		_, err := top.WriteTo(&b)
		if err != nil {
			log.Fatalln(err)
		}
		b.WriteString("\n")

		_, err = bottom.WriteTo(&b)
		if err != nil {
			log.Fatalln(err)
		}
		b.WriteString("\n")
	}

	return b.String()
}

func (g *Grid) ToPNG(w io.Writer) {
	width := cellSize * g.Columns
	height := cellSize * g.Rows

	img := image.NewRGBA(image.Rect(0, 0, width+1, height+1))
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)

	for row := 0; row < g.Rows; row++ {
		for column := 0; column < g.Columns; column++ {
			cell := g.Cell(row, column)

			x1 := cell.Column * cellSize
			y1 := cell.Row * cellSize
			x2 := (cell.Column + 1) * cellSize
			y2 := (cell.Row + 1) * cellSize

			if cell.North == nil {
				hLine(img, x1, y1, x2)
			}

			if cell.West == nil {
				vLine(img, x1, y1, y2)
			}

			if !cell.IsLinked(cell.East) {
				vLine(img, x2, y1, y2)
			}

			if !cell.IsLinked(cell.South) {
				hLine(img, x1, y2, x2)
			}
		}
	}

	err := png.Encode(w, img)
	if err != nil {
		log.Fatalln(err)
	}
}

func hLine(img *image.RGBA, x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, color.Black)
	}
}

func vLine(img *image.RGBA, x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, color.Black)
	}
}
