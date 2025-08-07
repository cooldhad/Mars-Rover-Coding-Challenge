package instruct

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"errors"
)

var _ interfaces.Instruct = &handler{}

type handler struct {
	move         interfaces.Move
	plateau      domain.Plateau
	instructions string
}

func (h handler) Instruct() (domain.Position, domain.Direction, error) {
	var hasInvalidChar bool
	for _, cmd := range h.instructions {
		switch cmd {
		case 'L':
			h.move.RotateLeft()
		case 'R':
			h.move.RotateRight()
		case 'M':
			h.move.Move(h.plateau)
		default:
			{
				hasInvalidChar = true
			}
		}
	}
	pos, dir := h.move.Get()
	if hasInvalidChar {
		return pos, dir, domain.AsBadRequestErr(errors.New("incorrect rover instructions, please use L,R,M only"))
	}

	return pos, dir, nil
}

func NewRover(moveRover interfaces.Move, plateau domain.Plateau, instructions string) interfaces.Instruct {
	return &handler{
		move:         moveRover,
		plateau:      plateau,
		instructions: instructions,
	}
}
