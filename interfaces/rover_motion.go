package interfaces

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
)

type RoverMotion interface {
	RotateLeft()
	RotateRight()
	Move(plateau domain.Plateau)
	ProcessInstructions(plateau domain.Plateau, instructions string)
}
