package rovermotion

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
)

var _ interfaces.RoverMotion = &handler{}

type handler struct {
	Position  domain.Position
	Direction domain.Direction
}

func (r *handler) RotateLeft() {
	switch r.Direction {
	case domain.North:
		r.Direction = domain.West
	case domain.East:
		r.Direction = domain.North
	case domain.South:
		r.Direction = domain.East
	case domain.West:
		r.Direction = domain.South
	}
}

func (r *handler) RotateRight() {
	switch r.Direction {
	case domain.North:
		r.Direction = domain.East
	case domain.East:
		r.Direction = domain.South
	case domain.South:
		r.Direction = domain.West
	case domain.West:
		r.Direction = domain.North
	}
}

func (r *handler) Move(plateau domain.Plateau) {
	newPosition := r.calculateNewPosition()
	if r.isWithinBounds(newPosition, plateau) {
		r.Position = newPosition
	}
}

func (r *handler) calculateNewPosition() domain.Position {
	newPosition := r.Position
	switch r.Direction {
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

func (r *handler) isWithinBounds(position domain.Position, plateau domain.Plateau) bool {
	return position.X >= 0 && position.X <= plateau.Width &&
		position.Y >= 0 && position.Y <= plateau.Height
}

func (r *handler) GetRover() (domain.Position, domain.Direction) {
	return r.Position, r.Direction
}

func NewRover(pos domain.Position, dir domain.Direction) interfaces.RoverMotion {
	return &handler{
		Position:  pos,
		Direction: dir,
	}
}
