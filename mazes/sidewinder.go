package mazes

import "math/rand"

func Sidewinder(grid *Grid) {
	for row := 0; row < grid.Rows; row++ {

		var run []*Cell

		for column := 0; column < grid.Columns; column++ {
			cell := grid.Cell(row, column)
			run = append(run, cell)

			if shouldCloseOut(cell, grid.Random) {
				i := grid.Random.Intn(len(run))
				member := run[i]
				if member.North != nil {
					member.LinkBidirectional(member.North)
				}
				run = nil
			} else {
				cell.LinkBidirectional(cell.East)
			}
		}
	}
}

func shouldCloseOut(cell *Cell, random *rand.Rand) bool {
	return cell.East == nil || (cell.North != nil && random.Intn(3) == 0)
}
