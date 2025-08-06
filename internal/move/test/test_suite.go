package test

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/move"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MoveTestSuite struct {
	suite.Suite

	//unit under test
	service interfaces.Move
}

func (suite *MoveTestSuite) SetupTest() {
	suite.service = move.NewRover(domain.Position{}, "")
}

func (suite *MoveTestSuite) TestRotateRight() {
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
		suite.service = move.NewRover(domain.Position{
			X: 0,
			Y: 0,
		}, c.input)
		suite.service.RotateRight()
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(c.expected, actualDir)
		suite.Assert().Equal(domain.Position{}, actualPos)
	}
}

func (suite *MoveTestSuite) TestRotateLeft() {
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
		suite.service = move.NewRover(domain.Position{
			X: 0,
			Y: 0,
		}, c.input)
		suite.service.RotateLeft()
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(c.expected, actualDir)
		suite.Assert().Equal(domain.Position{}, actualPos)
	}
}

func (suite *MoveTestSuite) TestMove() {
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
		suite.service = move.NewRover(t.startPos, t.dir)
		suite.service.Move(plateau)
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(t.expectedPos, actualPos)
		suite.Assert().Equal(t.dir, actualDir)
	}
}

func (suite *MoveTestSuite) TestMoveOutOfBounds() {
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
		suite.service = move.NewRover(t.start, t.dir)
		suite.service.Move(plateau)
		actualPos, actualDir := suite.service.Get()
		suite.Require().Equal(t.expected, actualPos)
		suite.Assert().Equal(t.dir, actualDir)
	}
}

func Move(t *testing.T) {
	t.Run("Move", func(t *testing.T) {
		suite.Run(t, &MoveTestSuite{})
	})
}
