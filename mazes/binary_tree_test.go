package mazes

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	g := NewGrid(4, 4)
	g.Random = rand.New(rand.NewSource(1))

	BinaryTree(g)
	actual := g.String()

	expected := `+---+---+---+---+
|               |
+---+---+   +   +
|           |   |
+   +   +   +   +
|   |   |   |   |
+---+   +---+   +
|       |       |
+---+---+---+---+
`

	assert.Equal(t, expected, actual)
}
