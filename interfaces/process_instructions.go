package interfaces

import "Mars-Rover-Coding-Challenge/internal/domain"

type ProcessInstructions interface {
	ProcessInstructions(roverMotion RoverMotion, plateau domain.Plateau, instructions string) (domain.Position, domain.Direction)
}
