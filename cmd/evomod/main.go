package main

import (
	"fmt"
	"time"

	"github.com/fishfugu/evomod/pkg/sim"
)

func main() {
	fmt.Println("Starting Evolution Modeller...")
	grid := sim.NewGrid(32, 32)
	reporter := sim.NewReporter(grid) // Initialise reporter with the grid

	// Setup initial creatures
	for i := 0; i < 10; i++ { // Run 100 ticks
		grid.Update(i)
		time.Sleep(500 * time.Millisecond) // Slow output for visibility
	}

	reporter.GenerateReport() // generate end of simulation report
}
