package sim

import (
	"math"
	"math/rand"
)

type Creature struct {
	ID               int
	Position         *Cell
	Speed            int
	Strength         int
	SenseOfDirection float64
	SearchRadius     int
	Grid             *Grid // Reference to the Grid
}

// Constructor for Creature
func NewCreature(id, speed, strength int, senseOfDirection, searchRadius float64, startCell *Cell, grid *Grid) *Creature {
	return &Creature{
		ID:               id,
		Position:         startCell,
		Speed:            speed,
		Strength:         strength,
		SenseOfDirection: senseOfDirection,
		SearchRadius:     int(searchRadius),
		Grid:             grid,
	}
}

func (c *Creature) Move() {
	if c.SenseOfDirection < 360 {
		direction := rand.Float64()*360 - c.SenseOfDirection
		radians := direction * math.Pi / 180 // Convert direction to radians for cosine and sine functions
		newX := c.Position.X + int(float64(c.Speed)*math.Cos(radians))
		newY := c.Position.Y + int(float64(c.Speed)*math.Sin(radians))
		c.Position = c.Grid.Position(newX, newY)
	} else {
		newX := c.Position.X + rand.Intn(2*c.Speed) - c.Speed
		newY := c.Position.Y + rand.Intn(2*c.Speed) - c.Speed
		c.Position = c.Grid.Position(newX, newY)
	}
}

func (c *Creature) SearchFood() {
	for dx := -c.SearchRadius; dx <= c.SearchRadius; dx++ {
		for dy := -c.SearchRadius; dy <= c.SearchRadius; dy++ {
			cell := c.Grid.Position(c.Position.X+dx, c.Position.Y+dy)
			if cell.Food > 0 {
				c.MoveTowards(cell.X, cell.Y)
				return
			}
		}
	}
	c.Move()
}

func (c *Creature) MoveTowards(x, y int) {
	if c.Position.X < x {
		c.Position.X++
	} else if c.Position.X > x {
		c.Position.X--
	}
	if c.Position.Y < y {
		c.Position.Y++
	} else if c.Position.Y > y {
		c.Position.Y--
	}
}

func (c *Creature) Act() {
	// Define what the creature does in one tick of the simulation
	// This example includes searching for food and potentially moving
	if c.NeedsToEat() {
		c.SearchFood()
	} else {
		c.Move()
	}
}

func (c *Creature) NeedsToEat() bool {
	// Placeholder for decision logic, assuming always true for simplicity
	return true
}
