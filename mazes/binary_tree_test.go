package mazes

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	g := NewGrid(4, 4)
	g.Random = rand.New(rand.NewSource(1))

	BinaryTree(g)
	actual := g.String()

	expected := `
+---+---+---+---+
|               |
+---+---+   +   +
|           |   |
+   +   +   +   +
|   |   |   |   |
+---+   +---+   +
|       |       |
+---+---+---+---+
`

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual))
}

func TestBinaryTreeWithDistances(t *testing.T) {
	g := NewGrid(4, 4)

	g.Random = rand.New(rand.NewSource(1))

	BinaryTree(g)

	start := g.Cell(0, 0)
	distances := start.Distances()
	contentsOfCellWithDistances := func(cell *Cell) string {
		if distance, ok := distances.GetDistance(cell); ok {
			return fmt.Sprintf(" %d ", distance)
		} else {
			return "   "
		}
	}
	g.ContentsOfCell = contentsOfCellWithDistances

	actual := g.String()

	expected := `
+---+---+---+---+
| 0   1   2   3 |
+---+---+   +   +
| 5   4   3 | 4 |
+   +   +   +   +
| 6 | 5 | 4 | 5 |
+---+   +---+   +
| 7   6 | 7   6 |
+---+---+---+---+
`

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual))
}
