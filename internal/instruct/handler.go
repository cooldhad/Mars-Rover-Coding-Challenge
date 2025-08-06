package instruct

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
)

var _ interfaces.Instruct = &handler{}

type handler struct {
	moveRover    interfaces.Move
	plateau      domain.Plateau
	instructions string
}

func (h handler) Instruct(
	moveRover interfaces.Move,
	plateau domain.Plateau,
	instructions string,
) (domain.Position, domain.Direction) {
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			moveRover.RotateLeft()
		case 'R':
			moveRover.RotateRight()
		case 'M':
			moveRover.Move(plateau)
		}
	}

	return moveRover.Get()
}

func NewRover(moveRover interfaces.Move, plateau domain.Plateau, instructions string) interfaces.Instruct {
	return &handler{
		moveRover:    moveRover,
		plateau:      plateau,
		instructions: instructions,
	}
}
