package test

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/rover"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoverTestSuite struct {
	suite.Suite

	//unit under test
	service interfaces.Rover

	// dependencies
	plateau domain.Plateau
}

func (suite *RoverTestSuite) SetupTest() {
	suite.service = rover.NewRover(domain.Position{}, "")
	suite.plateau = domain.Plateau{Width: 5, Height: 5}
}

func (suite *RoverTestSuite) TestRotateRight() {
	cases := []struct {
		input    domain.Direction
		expected domain.Direction
	}{
		{domain.North, domain.East},
		{domain.West, domain.North},
		{domain.South, domain.West},
		{domain.East, domain.South},
	}

	for _, c := range cases {
		suite.service = rover.NewRover(domain.Position{
			X: 0,
			Y: 0,
		}, c.input)
		suite.service.RotateRight()
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(c.expected, actualDir)
		suite.Assert().Equal(domain.Position{}, actualPos)
	}
}

func (suite *RoverTestSuite) TestRotateLeft() {
	cases := []struct {
		input    domain.Direction
		expected domain.Direction
	}{
		{domain.North, domain.West},
		{domain.West, domain.South},
		{domain.South, domain.East},
		{domain.East, domain.North},
	}

	for _, c := range cases {
		suite.service = rover.NewRover(domain.Position{
			X: 0,
			Y: 0,
		}, c.input)
		suite.service.RotateLeft()
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(c.expected, actualDir)
		suite.Assert().Equal(domain.Position{}, actualPos)
	}
}

func (suite *RoverTestSuite) TestMove() {
	plateau := domain.Plateau{
		Width:  5,
		Height: 5,
	}

	tests := []struct {
		startPos    domain.Position
		dir         domain.Direction
		expectedPos domain.Position
	}{
		{startPos: domain.Position{X: 2, Y: 2}, dir: domain.North, expectedPos: domain.Position{X: 2, Y: 3}},
		{startPos: domain.Position{X: 2, Y: 2}, dir: domain.South, expectedPos: domain.Position{X: 2, Y: 1}},
		{startPos: domain.Position{X: 2, Y: 2}, dir: domain.East, expectedPos: domain.Position{X: 3, Y: 2}},
		{startPos: domain.Position{X: 2, Y: 2}, dir: domain.West, expectedPos: domain.Position{X: 1, Y: 2}},
	}

	for _, t := range tests {
		suite.service = rover.NewRover(t.startPos, t.dir)
		suite.service.Move(plateau)
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(t.expectedPos, actualPos)
		suite.Assert().Equal(t.dir, actualDir)
	}
}

func (suite *RoverTestSuite) TestMoveOutOfBounds() {
	plateau := domain.Plateau{
		Width:  5,
		Height: 5,
	}
	tests := []struct {
		start    domain.Position
		dir      domain.Direction
		expected domain.Position
	}{
		{start: domain.Position{X: 0, Y: 0}, dir: domain.South, expected: domain.Position{X: 0, Y: 0}},
		{start: domain.Position{X: 0, Y: 0}, dir: domain.West, expected: domain.Position{X: 0, Y: 0}},
		{start: domain.Position{X: 5, Y: 5}, dir: domain.East, expected: domain.Position{X: 5, Y: 5}},
		{start: domain.Position{X: 5, Y: 5}, dir: domain.North, expected: domain.Position{X: 5, Y: 5}},
	}

	for _, t := range tests {
		suite.service = rover.NewRover(t.start, t.dir)
		suite.service.Move(plateau)
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(t.expected, actualPos)
		suite.Assert().Equal(t.dir, actualDir)
	}
}

func (suite *RoverTestSuite) TestInstruct() {
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
			name:         "Rover turning only",
			startPos:     domain.Position{X: 2, Y: 2},
			startDir:     domain.South,
			instructions: "RRLL",
			expectedPos:  domain.Position{X: 2, Y: 2},
			expectedDir:  domain.South,
		},
	}

	for _, tc := range tests {
		suite.service = rover.NewRover(tc.startPos, tc.startDir)
		actualPos, actualDir, err := suite.service.Instruct(suite.plateau, tc.instructions)
		suite.Require().NoError(err)
		suite.Require().Equal(tc.expectedPos, actualPos)
		suite.Assert().Equal(tc.expectedDir, actualDir)
	}
}

func (suite *RoverTestSuite) TestInstructWhenRoverIsOutOfBounds() {
	startPos := domain.Position{Y: 5}
	startDir := domain.North
	suite.service = rover.NewRover(startPos, startDir)
	expectedPos := domain.Position{Y: 5} // can't go beyond Y=5
	expectedDir := domain.North
	actualPos, actualDir, err := suite.service.Instruct(suite.plateau, "MMM")
	suite.Require().NoError(err)
	suite.Assert().Equal(expectedPos, actualPos)
	suite.Assert().Equal(expectedDir, actualDir)
}

func (suite *RoverTestSuite) TestIncorrectInstructions() {
	startPos := domain.Position{X: 1, Y: 2}
	startDir := domain.North
	suite.service = rover.NewRover(startPos, startDir)
	instructions := "LMXZ1LRM" // X, Z, 1 are invalid
	expectedPos := domain.Position{X: 0, Y: 2}
	expectedDir := domain.West
	actualPos, actualDir, err := suite.service.Instruct(suite.plateau, instructions)
	suite.Require().Error(err)
	suite.Require().True(domain.IsBadRequestErr(err))
	suite.Require().EqualError(err, "incorrect rover instructions, please use L,R,M only")
	suite.Assert().Equal(expectedPos, actualPos)
	suite.Assert().Equal(expectedDir, actualDir)
}

func Rover(t *testing.T) {
	t.Run("Rover", func(t *testing.T) {
		suite.Run(t, &RoverTestSuite{})
	})
}
