package main

import (
	"fmt"

	"github.com/ChrisWilding/mazes/mazes"
)

func main() {
	g := mazes.NewGrid(4, 4)
	mazes.BinaryTree(g)
	fmt.Print(g)
}
