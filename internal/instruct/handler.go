package instruct

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
)

var _ interfaces.Instruct = &handler{}

type handler struct {
	move         interfaces.Move
	plateau      domain.Plateau
	instructions string
}

func (h handler) Instruct(
	move interfaces.Move,
	plateau domain.Plateau,
	instructions string,
) (domain.Position, domain.Direction) {
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			move.RotateLeft()
		case 'R':
			move.RotateRight()
		case 'M':
			move.Move(plateau)
		default:
			// ensuring that invalid comments are ignored
		}
	}

	return move.Get()
}

func NewRover(moveRover interfaces.Move, plateau domain.Plateau, instructions string) interfaces.Instruct {
	return &handler{
		move:         moveRover,
		plateau:      plateau,
		instructions: instructions,
	}
}
