package main

import (
	"Mars-Rover-Coding-Challenge/internal/domain"
	"Mars-Rover-Coding-Challenge/internal/rovermotion"
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
	width, _ := strconv.Atoi(grid[0])
	height, _ := strconv.Atoi(grid[1])
	plateau := domain.Plateau{Width: width, Height: height}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Position and direction
		parts := strings.Fields(line)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		position := domain.Position{
			X: x,
			Y: y,
		}
		dir := domain.Direction(parts[2])
		rover := rovermotion.NewRoverMotionHandler(position, dir)

		// Read movement line
		scanner.Scan()
		instructions := scanner.Text()

		rover.ProcessInstructions(plateau, instructions)
		fmt.Printf("%d %d %s\n", rover.Position.X, rover.Position.Y, rover.Direction)
	}
}
