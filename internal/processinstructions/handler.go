package processinstructions

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
)

var _ interfaces.ProcessInstructions = &handler{}

type handler struct {
	roverMotion  interfaces.RoverMotion
	plateau      domain.Plateau
	instructions string
}

func (h handler) ProcessInstructions(
	roverMotion interfaces.RoverMotion,
	plateau domain.Plateau,
	instructions string,
) (domain.Position, domain.Direction) {
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			roverMotion.RotateLeft()
		case 'R':
			roverMotion.RotateRight()
		case 'M':
			roverMotion.Move(plateau)
		}
	}

	return roverMotion.GetRover()
}

func NewProcessInstructions(roverMotion interfaces.RoverMotion, plateau domain.Plateau, instructions string) interfaces.ProcessInstructions {
	return &handler{
		roverMotion:  roverMotion,
		plateau:      plateau,
		instructions: instructions,
	}
}
