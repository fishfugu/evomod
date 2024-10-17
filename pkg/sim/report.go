package sim

import (
	"fmt"
)

// Reporter holds reporting methods / properties
type Reporter struct {
	Grid *Grid
}

// NewReporter creates new instance of Reporter
func NewReporter(grid *Grid) *Reporter {
	return &Reporter{Grid: grid}
}

// GenerateReport summarises sim outcomes
func (r *Reporter) GenerateReport() {
	creatureCount := 0
	for y := 0; y < r.Grid.Height; y++ {
		for x := 0; x < r.Grid.Width; x++ {
			cell := r.Grid.Cells[y][x]
			creatureCount += len(cell.Creatures)
		}
	}
	fmt.Printf("End of Simulation Report:\n")
	fmt.Printf("Total Creatures Surviving: %d\n", creatureCount)
	// TODO add stats
}
