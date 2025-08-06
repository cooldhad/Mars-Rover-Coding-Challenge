package rovermotion

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
)

var _ interfaces.RoverMotion = &MotionHandler{}

type MotionHandler struct {
	Position  domain.Position
	Direction domain.Direction
}

func (r *MotionHandler) RotateLeft() {
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

func (r *MotionHandler) RotateRight() {
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

func (r *MotionHandler) Move(plateau domain.Plateau) {
	newPosition := r.calculateNewPosition()
	if r.isWithinBounds(newPosition, plateau) {
		r.Position = newPosition
	}
}

func (r *MotionHandler) calculateNewPosition() domain.Position {
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

func (r *MotionHandler) isWithinBounds(position domain.Position, plateau domain.Plateau) bool {
	return position.X >= 0 && position.X <= plateau.Width &&
		position.Y >= 0 && position.Y <= plateau.Height
}

func (r *MotionHandler) ProcessInstructions(plateau domain.Plateau, instructions string) {
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			r.RotateLeft()
		case 'R':
			r.RotateRight()
		case 'M':
			r.Move(plateau)
		}
	}
}

func NewRoverMotionHandler(pos domain.Position, dir domain.Direction) MotionHandler {
	return MotionHandler{
		Position:  pos,
		Direction: dir,
	}
}
