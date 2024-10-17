package main

import (
	"fmt"

	"github.com/fishfugu/evomod/pkg/sim" // Ensure the import path matches your module setup
)

func main() {
	fmt.Println("Starting Evolution Modeller...")

	// Create a new grid
	grid := sim.NewGrid(100, 100) // Example size: 100x100

	// Create creatures and place them on the grid
	startCell := grid.Position(50, 50) // Starting position for the creature
	creature := sim.NewCreature(1, 5, 10, 180, 10, startCell, grid)
	startCell.Creatures = append(startCell.Creatures, creature)

	// Optionally, create more creatures and configure them similarly

	// Example of starting the simulation (you'll need to implement this loop)
	for {
		grid.Update() // This method needs to be implemented to progress the simulation
	}
}
