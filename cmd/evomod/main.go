package main

import (
	"fmt"
	"time"

	"github.com/fishfugu/evomod/pkg/sim"
)

func main() {
	fmt.Println("Starting Evolution Modeller...")
	grid := sim.NewGrid(100, 100)
	reporter := sim.NewReporter(grid) // Initialise reporter with the grid

	// Setup initial creatures
	for i := 0; i < 100; i++ { // Run 100 ticks
		grid.Update()
		time.Sleep(500 * time.Millisecond) // Slow output for visibility
	}

	reporter.GenerateReport() // generate end of simulation report
}
