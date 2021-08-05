package main

import (
	"fmt"
	"testing"

	"github.com/ChrisWilding/mazes/mazes"
)

func TestMain(t *testing.T) {
	g := mazes.NewGrid(4, 4)
	mazes.BinaryTree(g)
	fmt.Print(g)
}
