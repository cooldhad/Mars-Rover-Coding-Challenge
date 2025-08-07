package move

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
)

var _ interfaces.Move = &handler{}

type handler struct {
	Position  domain.Position
	Direction domain.Direction
}

func (h *handler) RotateLeft() {
	switch h.Direction {
	case domain.North:
		h.Direction = domain.West
	case domain.East:
		h.Direction = domain.North
	case domain.South:
		h.Direction = domain.East
	case domain.West:
		h.Direction = domain.South
	}
}

func (h *handler) RotateRight() {
	switch h.Direction {
	case domain.North:
		h.Direction = domain.East
	case domain.East:
		h.Direction = domain.South
	case domain.South:
		h.Direction = domain.West
	case domain.West:
		h.Direction = domain.North
	}
}

func (h *handler) Move(plateau domain.Plateau) {
	newPosition := h.calculateNewPosition()
	if h.isWithinBounds(newPosition, plateau) {
		h.Position = newPosition
	}
}

func (h *handler) calculateNewPosition() domain.Position {
	newPosition := h.Position
	switch h.Direction {
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

func (h *handler) Get() (domain.Position, domain.Direction) {
	return h.Position, h.Direction
}

func NewRover(pos domain.Position, dir domain.Direction) interfaces.Move {
	return &handler{
		Position:  pos,
		Direction: dir,
	}
}
