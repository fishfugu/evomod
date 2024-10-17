# Evolution Modeller (Evomod)

## Overview

Evolution Modeller (Evomod) is a grid-based simulation environment where virtual creatures interact, search for food, and exhibit behaviors based on evolutionary algorithms. This project is designed to explore the principles of evolutionary biology and artificial life in a controlled setting.

## Features

- **Grid-based Environment**: A simple but expansive grid where all interactions take place.
- **Creature Behaviors**: Creatures in the simulation can move, search for food, and interact with each other.
- **Dynamic Simulation**: Each tick of the simulation updates the states of all creatures based on predefined behaviors.

## Installation

### Prereqs

- Go 1.15 or higher

### Set Up

```bash
git clone https://github.com/fishfugu/evomod.git
cd evomod
go build -o evomod ./cmd/evomod
```

## Usage

```bash
./evomod
```

Watch inline. For now, the sim runs in text and outputs updates each tick (future versions will include visual tools for better interaction and observation).

## Contributing

Contributions to the Evolution Modeller welcome!

[] TODO: make Contributing.md

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Inspired by classic models of artificial life and evolutionary computation... and life.

### TODO

- **CONTRIBUTING.md**: guidelines on how contributors can help with the project.
- **Screenshots** or **GIFs** of the simulation in action.
- **Section on future features**.
- **Tech details** about the algos / methods used in the simulation.
