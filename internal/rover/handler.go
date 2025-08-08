package rover

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"errors"
)

var _ interfaces.Rover = &rover{}

type rover struct {
	marsRover domain.Rover
}

func (r *rover) RotateLeft() {
	switch r.marsRover.Direction {
	case domain.North:
		r.marsRover.Direction = domain.West
	case domain.East:
		r.marsRover.Direction = domain.North
	case domain.South:
		r.marsRover.Direction = domain.East
	case domain.West:
		r.marsRover.Direction = domain.South
	}
}

func (r *rover) RotateRight() {
	switch r.marsRover.Direction {
	case domain.North:
		r.marsRover.Direction = domain.East
	case domain.East:
		r.marsRover.Direction = domain.South
	case domain.South:
		r.marsRover.Direction = domain.West
	case domain.West:
		r.marsRover.Direction = domain.North
	}
}

func (r *rover) Move(plateau domain.Plateau) {
	newPosition := r.calculateNewPosition()
	if r.isWithinBounds(newPosition, plateau) {
		r.marsRover.Position = newPosition
	}
}

func (r *rover) calculateNewPosition() domain.Position {
	newPosition := r.marsRover.Position
	switch r.marsRover.Direction {
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

func (r *rover) Instruct(plateau domain.Plateau, instructions string) (domain.Rover, error) {
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
		return r.marsRover, domain.AsBadRequestErr(errors.New("incorrect rover instructions, please use L,R,M only"))
	}

	return r.marsRover, nil
}

func (r *rover) Get() domain.Rover {
	return r.marsRover
}

func NewRover(marsRover domain.Rover) interfaces.Rover {
	return &rover{
		marsRover: marsRover,
	}
}
