package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ChrisWilding/mazes-for-programmers/mazes"
)

func main() {
	algorithm := flag.String("a", "sidewinder", "algorithm")
	width := flag.Int("w", 5, "width")
	height := flag.Int("h", 5, "height")
	output := flag.String("o", "stdout", "output")
	flag.Parse()

	if *width < 0 {
		fmt.Fprintf(os.Stderr, "-w was %d, width must be greater than 0\n", *width)
		os.Exit(2)
	}

	if *height < 0 {
		fmt.Fprintf(os.Stderr, "-h was %d, height must be greater than 0\n", *height)
		os.Exit(2)
	}

	g := mazes.NewGrid(*width, *height)

	switch *algorithm {
	case "binary-tree":
		mazes.BinaryTree(g)
	case "sidewinder":
		mazes.Sidewinder(g)
	default:
		fmt.Fprintf(os.Stderr, "-a was %s, algorithm must be one of: binary-tree or sidewinder\n", *algorithm)
		os.Exit(2)
	}

	if *output == "stdout" {
		g.ToPNG(os.Stdout)
		os.Exit(0)
	}

	f, err := os.OpenFile(*output, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	g.ToPNG(f)
}
