package interfaces

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
)

type Move interface {
	RotateLeft()
	RotateRight()
	Move(plateau domain.Plateau)
	Get() (domain.Position, domain.Direction)
}
