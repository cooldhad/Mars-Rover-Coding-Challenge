package test

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/instruct"
	"Mars-Rover-Coding-Challenge/internal/move"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InstructTestSuite struct {
	suite.Suite

	//unit under test
	service interfaces.Instruct

	//dependencies
	move         interfaces.Move
	plateau      domain.Plateau
	instructions string
}

func (suite *InstructTestSuite) SetupTest() {
	suite.move = move.NewRover(domain.Position{}, "")
	suite.plateau = domain.Plateau{
		Width:  5,
		Height: 5,
	}
	suite.instructions = ""

	suite.service = instruct.NewRover(suite.move, suite.plateau, suite.instructions)
}

func (suite *InstructTestSuite) TestInstruct() {
	tests := []struct {
		name         string
		startPos     domain.Position
		startDir     domain.Direction
		instructions string
		expectedPos  domain.Position
		expectedDir  domain.Direction
	}{
		{
			name:         "Rover 1 sample input",
			startPos:     domain.Position{X: 1, Y: 2},
			startDir:     domain.North,
			instructions: "LMLMLMLMM",
			expectedPos:  domain.Position{X: 1, Y: 3},
			expectedDir:  domain.North,
		},
		{
			name:         "Rover 2 sample input",
			startPos:     domain.Position{X: 3, Y: 3},
			startDir:     domain.East,
			instructions: "MMRMMRMRRM",
			expectedPos:  domain.Position{X: 5, Y: 1},
			expectedDir:  domain.East,
		},
		{
			name:         "Rover hitting boundary (North)",
			startPos:     domain.Position{Y: 5},
			startDir:     domain.North,
			instructions: "MMM",
			expectedPos:  domain.Position{Y: 5}, // can't go beyond Y=5
			expectedDir:  domain.North,
		},
		{
			name:         "Rover turning only",
			startPos:     domain.Position{X: 2, Y: 2},
			startDir:     domain.South,
			instructions: "RRLL",
			expectedPos:  domain.Position{X: 2, Y: 2},
			expectedDir:  domain.South,
		},
		{
			name:         "Rover with invalid instructions hitting the boundary (West)",
			startPos:     domain.Position{X: 1, Y: 2},
			startDir:     domain.North,
			instructions: "LMXZ1LRM", // X, Z, 1 are invalid
			expectedPos:  domain.Position{X: 0, Y: 2},
			expectedDir:  domain.West,
		},
	}

	for _, tc := range tests {
		suite.move = move.NewRover(tc.startPos, tc.startDir)
		actualPos, actualDir := suite.service.Instruct(suite.move, suite.plateau, tc.instructions)
		suite.Require().Equal(tc.expectedPos, actualPos)
		suite.Assert().Equal(tc.expectedDir, actualDir)
	}
}

func Instruct(t *testing.T) {
	t.Run("InstructRover", func(t *testing.T) {
		suite.Run(t, &InstructTestSuite{})
	})
}
