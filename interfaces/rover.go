package interfaces

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
)

type Rover interface {
	RotateLeft()
	RotateRight()
	Move(plateau domain.Plateau)
	Instruct(plateau domain.Plateau, instructions string) (domain.Rover, error)
	Get() domain.Rover
}
