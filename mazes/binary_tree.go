package mazes

func BinaryTree(grid *Grid) {
	for row := 0; row < grid.Rows; row++ {
		for column := 0; column < grid.Columns; column++ {
			cell := grid.Cell(row, column)

			var neighbours []*Cell

			if cell.North != nil {
				neighbours = append(neighbours, cell.North)
			}

			if cell.East != nil {
				neighbours = append(neighbours, cell.East)
			}

			l := len(neighbours)
			if l == 0 {
				break
			}

			i := grid.Random.Intn(l)
			neighbour := neighbours[i]

			if neighbour != nil {
				cell.LinkBidirectional(neighbour)
			}
		}
	}
}
