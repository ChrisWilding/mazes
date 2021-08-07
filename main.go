package main

import (
	"fmt"
	"math/rand"

	"github.com/ChrisWilding/mazes/mazes"
)

func main() {
	g := mazes.NewGrid(4, 4)
	g.Random = rand.New(rand.NewSource(1))
	mazes.BinaryTree(g)
	fmt.Print(g)
}
