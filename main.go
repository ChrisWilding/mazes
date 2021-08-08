package main

import (
	"fmt"
	"math/rand"

	"github.com/ChrisWilding/mazes-for-programmers/mazes"
)

func main() {
	g := mazes.NewGrid(4, 4)
	g.Random = rand.New(rand.NewSource(1))
	mazes.Sidewinder(g)
	fmt.Print(g)
}
