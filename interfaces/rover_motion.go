package interfaces

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
)

type RoverMotion interface {
	RotateLeft()
	RotateRight()
	Move(plateau domain.Plateau)
	GetRover() (domain.Position, domain.Direction)
}
