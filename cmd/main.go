package main

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/instruct"
	"Mars-Rover-Coding-Challenge/internal/move"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	grid := strings.Fields(scanner.Text())
	width, err := strconv.Atoi(grid[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error converting width to int: %s\n", err)
		os.Exit(1)
	}
	height, err := strconv.Atoi(grid[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error converting height to int: %s\n", err)
		os.Exit(1)
	}
	plateau := domain.Plateau{Width: width, Height: height}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Position and direction
		parts := strings.Fields(line)
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error converting xCoordinate to int: %s\n", err)
			os.Exit(1)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error converting yCoordinate to int: %s\n", err)
			os.Exit(1)
		}
		position := domain.Position{
			X: x,
			Y: y,
		}
		dir := domain.Direction(parts[2])
		if dir != "N" && dir != "E" && dir != "S" && dir != "W" {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing input direction %s\n", dir)
		}
		moveRover := move.NewRover(position, dir)

		// Read movement (instructions) line
		scanner.Scan()
		instructions := scanner.Text()

		roverInstructions := instruct.NewRover(moveRover, plateau, instructions)
		newPosition, newDirection := roverInstructions.Instruct(moveRover, plateau, instructions)
		fmt.Printf("%d %d %s\n", newPosition.X, newPosition.Y, newDirection)
	}
}
