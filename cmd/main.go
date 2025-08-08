package main

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/rover"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RoverInput struct {
	index        int
	initialPos   string
	instructions string
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	// reading the plateau
	scanner.Scan()
	grid := strings.Fields(scanner.Text())
	plateau, err := parsePlateau(grid)
	if err != nil {
		return err
	}
	// collecting all rover inputs first
	var inputs []RoverInput
	roverIndex := 0
	for {
		if !scanner.Scan() {
			break
		}
		posLine := strings.TrimSpace(scanner.Text())
		if posLine == "" {
			break
		}
		// read the corresponding rover's instruction line
		if !scanner.Scan() {
			return fmt.Errorf("missing instructions for rover %d", roverIndex)
		}
		cmdLine := strings.TrimSpace(scanner.Text())
		inputs = append(inputs, RoverInput{
			index:        roverIndex,
			initialPos:   posLine,
			instructions: cmdLine,
		})
		roverIndex++
	}

	// process each rover after collecting all inputs
	for _, input := range inputs {
		parts := strings.Fields(input.initialPos)
		position, dir, err := parsePositionAndDirection(parts)
		if err != nil {
			return err
		}
		marsRover := rover.NewRover(position, dir)
		// Read movement (instructions) line
		if !scanner.Scan() {
			return fmt.Errorf("missing rover instructions")
		}
		instructions := strings.TrimSpace(scanner.Text())
		newPosition, newDirection, err := marsRover.Instruct(plateau, instructions)
		if err != nil {
			return err
		}
		fmt.Printf("%d %d %s\n", newPosition.X, newPosition.Y, newDirection)
	}

	return scanner.Err()
}
func parsePlateau(parts []string) (domain.Plateau, error) {
	if len(parts) != 2 {
		return domain.Plateau{}, fmt.Errorf("invalid input, expected 2 values, got %d", len(parts))
	}
	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return domain.Plateau{}, err
	}
	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return domain.Plateau{}, err
	}
	plateau := domain.Plateau{Width: width, Height: height}

	return plateau, nil
}

func parsePositionAndDirection(parts []string) (domain.Position, domain.Direction, error) {
	if len(parts) != 3 {
		return domain.Position{}, "", fmt.Errorf("invalid input, expected 3 values, got %d", len(parts))
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return domain.Position{}, "", err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return domain.Position{}, "", err
	}
	dir := domain.Direction(parts[2])
	position := domain.Position{
		X: x,
		Y: y,
	}

	return position, dir, nil
}
