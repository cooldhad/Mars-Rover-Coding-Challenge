package test

import (
	"Mars-Rover-Coding-Challenge/interfaces"
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/rovermotion"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoverMotionTestSuite struct {
	suite.Suite

	service interfaces.RoverMotion
}

func (suite *RoverMotionTestSuite) SetupTest() {
	suite.service = rovermotion.NewRover(domain.Position{
		X: 0,
		Y: 0,
	}, "")
}

func (suite *RoverMotionTestSuite) TestRotateRight() {
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
		suite.service = rovermotion.NewRover(domain.Position{
			X: 0,
			Y: 0,
		}, c.input)
		suite.service.RotateRight()
		_, dir := suite.service.GetRover()
		suite.Equal(dir, c.expected)
	}
}

func (suite *RoverMotionTestSuite) TestRotateLeft() {
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
		suite.service = rovermotion.NewRover(domain.Position{
			X: 0,
			Y: 0,
		}, c.input)
		suite.service.RotateLeft()
		_, dir := suite.service.GetRover()
		suite.Assert().Equal(dir, c.expected)
	}
}

func (suite *RoverMotionTestSuite) TestMove() {
	plateau := domain.Plateau{
		Width:  5,
		Height: 5,
	}

	tests := []struct {
		start    domain.Position
		dir      domain.Direction
		expected domain.Position
	}{
		{start: domain.Position{X: 2, Y: 2}, dir: domain.North, expected: domain.Position{X: 2, Y: 3}},
		{start: domain.Position{X: 2, Y: 2}, dir: domain.South, expected: domain.Position{X: 2, Y: 1}},
		{start: domain.Position{X: 2, Y: 2}, dir: domain.East, expected: domain.Position{X: 3, Y: 2}},
		{start: domain.Position{X: 2, Y: 2}, dir: domain.West, expected: domain.Position{X: 1, Y: 2}},
	}

	for _, t := range tests {
		suite.service = rovermotion.NewRover(t.start, t.dir)
		suite.service.Move(plateau)
		newPos, newDir := suite.service.GetRover()
		suite.Assert().Equal(newPos, t.expected)
		suite.Assert().Equal(newDir, t.dir)
	}
}

func RoverMotion(t *testing.T) {
	t.Run("RoverMotion", func(t *testing.T) {
		suite.Run(t, &RoverMotionTestSuite{})
	})
}
