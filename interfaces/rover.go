package interfaces

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
)

type Rover interface {
	RotateLeft()
	RotateRight()
	Move(plateau domain.Plateau)
	Instruct(plateau domain.Plateau, instructions string) (domain.Position, domain.Direction, error)
	Get() (domain.Position, domain.Direction)
}
