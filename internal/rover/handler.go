package rover

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"errors"
)

var _ interfaces.Rover = &rover{}

type rover struct {
	Position  domain.Position
	Direction domain.Direction
}

func (r *rover) RotateLeft() {
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

func (r *rover) RotateRight() {
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

func (r *rover) Move(plateau domain.Plateau) {
	newPosition := r.calculateNewPosition()
	if r.isWithinBounds(newPosition, plateau) {
		r.Position = newPosition
	}
}

func (r *rover) calculateNewPosition() domain.Position {
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

func (r *rover) isWithinBounds(position domain.Position, plateau domain.Plateau) bool {
	return position.X >= 0 && position.X <= plateau.Width &&
		position.Y >= 0 && position.Y <= plateau.Height
}

func (r *rover) Instruct(plateau domain.Plateau, instructions string) (domain.Position, domain.Direction, error) {
	var hasInvalidChar bool
	for _, cmd := range instructions {
		switch cmd {
		case 'L':
			r.RotateLeft()
		case 'R':
			r.RotateRight()
		case 'M':
			r.Move(plateau)
		default:
			{
				hasInvalidChar = true
			}
		}
	}
	if hasInvalidChar {
		return r.Position, r.Direction, domain.AsBadRequestErr(errors.New("incorrect rover instructions, please use L,R,M only"))
	}

	return r.Position, r.Direction, nil
}

func (r *rover) Get() (domain.Position, domain.Direction) {
	return r.Position, r.Direction
}

func NewRover(pos domain.Position, dir domain.Direction) interfaces.Rover {
	return &rover{
		Position:  pos,
		Direction: dir,
	}
}
