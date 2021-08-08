# Mazes for Programmers

[Mazes for Programmers](http://www.mazesforprogrammers.com) is a book written by [Jamis Buck](http://weblog.jamisbuck.org) and published by [The Pragmatic Bookshelve](https://pragprog.com/titles/jbmaze/mazes-for-programmers/).

The alogrithms and implementations in the book are all written in Ruby. This repository contains my implementation written in Go.

# Prerequisites

1. [Go 1.16](https://golang.org/doc/install)

# Setup

```sh
$ git clone git@github.com:ChrisWilding/mazes-for-programmers.git
$ cd ./mazes-for-programmers
$ go build
```

# Usage

```sh
$ ./mazes-for-programmers -w 10 -h 10 -a sidewinder -o sidewinder.png
```

This is an example of a maze generated using the binary tree algorithm

![A maze generated using the binary tree alorithm](https://raw.githubusercontent.com/ChrisWilding/mazes-for-programmers/main/images/binary-tree.png)

This is an example of a maze generated using the sidewinder algorithm

![A maze generated using the sidewinder alorithm](https://raw.githubusercontent.com/ChrisWilding/mazes-for-programmers/main/images/binary-tree.png)

# Testing

```sh
$ go test ./....
```
