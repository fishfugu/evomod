package sim

import (
	"fmt"
	"strings"
)

type Cell struct {
	X, Y      int
	Food      int
	Creatures []*Creature
}

type Grid struct {
	Width, Height int
	Cells         [][]*Cell
}

func NewGrid(width, height int) *Grid {
	grid := &Grid{
		Width:  width,
		Height: height,
		Cells:  make([][]*Cell, height),
	}
	for i := range grid.Cells {
		grid.Cells[i] = make([]*Cell, width)
		for j := range grid.Cells[i] {
			grid.Cells[i][j] = &Cell{X: j, Y: i}
		}
	}
	return grid
}

func (g *Grid) Position(x, y int) *Cell {
	if x < 0 {
		x = 0
	} else if x >= g.Width {
		x = g.Width - 1
	}
	if y < 0 {
		y = 0
	} else if y >= g.Height {
		y = g.Height - 1
	}
	return g.Cells[y][x]
}

func (g *Grid) Update(count int) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Cells[y][x]
			cellSymbol := "."
			if len(cell.Creatures) > 0 {
				cellSymbol = "C" // a Creature
			}
			if cell.Food > 0 {
				cellSymbol = "F" // Food
			}
			fmt.Print(cellSymbol + " ")
		}
		fmt.Println()
	}
	fmt.Printf("%s %d\n", strings.Repeat("-", g.Width*2), count) // div between ticks
}
