package mazes

func BinaryTree(grid *Grid) {
	for row := 0; row < grid.Rows; row++ {
		for column := 0; column < grid.Columns; column++ {
			cell := grid.Cell(row, column)

			var neighbour *Cell

			if random.Int()%2 == 0 {
				neighbour = cell.North
			} else {
				neighbour = cell.East
			}

			if neighbour != nil {
				cell.LinkBidirectional(neighbour)
			}
		}
	}
}
