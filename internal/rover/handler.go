package rover

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"errors"
)

var _ interfaces.Rover = &handler{}

type handler struct {
	marsRover domain.Rover
}

func (h *handler) RotateLeft() {
	switch h.marsRover.Direction {
	case domain.North:
		h.marsRover.Direction = domain.West
	case domain.East:
		h.marsRover.Direction = domain.North
	case domain.South:
		h.marsRover.Direction = domain.East
	case domain.West:
		h.marsRover.Direction = domain.South
	}
}

func (h *handler) RotateRight() {
	switch h.marsRover.Direction {
	case domain.North:
		h.marsRover.Direction = domain.East
	case domain.East:
		h.marsRover.Direction = domain.South
	case domain.South:
		h.marsRover.Direction = domain.West
	case domain.West:
		h.marsRover.Direction = domain.North
	}
}

func (h *handler) Move(plateau domain.Plateau) {
	newPosition := h.calculateNewPosition()
	if h.isWithinBounds(newPosition, plateau) {
		h.marsRover.Position = newPosition
	}
}

func (h *handler) calculateNewPosition() domain.Position {
	newPosition := h.marsRover.Position
	switch h.marsRover.Direction {
	case domain.North:
		newPosition.Y++
	case domain.South:
		newPosition.Y--
	case domain.East:
		newPosition.X++
	case domain.West:
		newPosition.X--
	}

	return newPosition
}

func (h *handler) isWithinBounds(position domain.Position, plateau domain.Plateau) bool {
	return position.X >= 0 && position.X <= plateau.Width &&
		position.Y >= 0 && position.Y <= plateau.Height
}

func (h *handler) Instruct(plateau domain.Plateau, instructions string) (domain.Rover, error) {
	var hasInvalidChar bool
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			h.RotateLeft()
		case 'R':
			h.RotateRight()
		case 'M':
			h.Move(plateau)
		default:
			{
				hasInvalidChar = true
			}
		}
	}
	if hasInvalidChar {
		return h.marsRover, domain.AsBadRequestErr(errors.New("incorrect rover instructions, please use L,R,M only"))
	}

	return h.marsRover, nil
}

func (h *handler) Get() domain.Rover {
	return h.marsRover
}

func NewRover(marsRover domain.Rover) interfaces.Rover {
	return &handler{
		marsRover: marsRover,
	}
}
