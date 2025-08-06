package interfaces

import "Mars-Rover-Coding-Challenge/internal/domain"

type Instruct interface {
	Instruct(move Move, plateau domain.Plateau, instructions string) (domain.Position, domain.Direction)
}
