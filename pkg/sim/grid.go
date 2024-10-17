package sim

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

func (g *Grid) Update() {
	// Example of a simple update loop for each cell in the grid
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Cells[y][x]
			// Update each creature in this cell
			for _, creature := range cell.Creatures {
				creature.Act() // Act method to be defined in creature.go
			}
			// Consider adding code to handle food regeneration or depletion here
		}
	}
	// Additional logic to handle environment changes or global events
}
