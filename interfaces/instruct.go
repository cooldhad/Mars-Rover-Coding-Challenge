package interfaces

import "Mars-Rover-Coding-Challenge/internal/domain"

type Instruct interface {
	Instruct() (domain.Position, domain.Direction, error)
}
