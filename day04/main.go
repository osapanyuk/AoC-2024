package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

const SEARCHWORD = "XMAS"

type Vector struct {
	x int
	y int
}

var (
	NORTH     = Vector{x: 0, y: -1}
	NORTHEAST = Vector{x: 1, y: -1}
	EAST      = Vector{x: 1, y: 0}
	SOUTHEAST = Vector{x: 1, y: 1}
	SOUTH     = Vector{x: 0, y: 1}
	SOUTHWEST = Vector{x: -1, y: 1}
	WEST      = Vector{x: -1, y: 0}
	NORTHWEST = Vector{x: -1, y: -1}
)

func init() {
	if len(input) == 0 {
		log.Fatal("input.txt is empty")
	}
	input = strings.TrimRight(input, "\n")
}

func main() {
	part := flag.Int(
		"part",
		0,
		"Specifies which part to execute. '1' for part1, '2' for part2",
	)

	flag.Parse()

	result := 0
	switch *part {
	case 1:
		fmt.Println("Running Part 1")
		result = partOne(&input)
	case 2:
		fmt.Println("Running Part 2")
		result = partTwo(&input)
	default:
		log.Fatal("Part argument is out of supported range. '1' for part1, '2' for part2")
	}

	fmt.Println("Result:", result)
}

func partOne(input *string) int {
	parsedInput := parseInput(input)
	maxY := len(parsedInput)
	maxX := len(parsedInput[0])

	count := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if parsedInput[y][x] == rune(SEARCHWORD[0]) {
				directions := pruneDirections(x, y, maxX, maxY)
				for _, direction := range directions {
					if findWord(parsedInput, SEARCHWORD[1:], x+direction.x, y+direction.y, direction) {
						count++
					}
				}
			}
		}
	}

	return count
}

func partTwo(input *string) int {
	parsedInput := parseInput(input)
	maxY := len(parsedInput)
	maxX := len(parsedInput[0])

	count := 0
	for y := 1; y < maxY-1; y++ {
		for x := 1; x < maxX-1; x++ {
			if parsedInput[y][x] == rune('A') {
				diagRunes := []rune{}
				diagRunes = append(diagRunes, parsedInput[y+NORTHWEST.y][x+NORTHWEST.x])
				diagRunes = append(diagRunes, parsedInput[y+NORTHEAST.y][x+NORTHEAST.x])
                diagRunes = append(diagRunes, parsedInput[y+SOUTHEAST.y][x+SOUTHEAST.x])
				diagRunes = append(diagRunes, parsedInput[y+SOUTHWEST.y][x+SOUTHWEST.x])
				if string(diagRunes) == "MMSS" ||
					string(diagRunes) == "SSMM" ||
					string(diagRunes) == "SMMS" ||
					string(diagRunes) == "MSSM" {
					count++
				}
			}
		}
	}

	return count
}

func parseInput(input *string) [][]rune {
	var result [][]rune

	for line := range strings.SplitSeq(*input, "\n") {
		result = append(result, []rune(line))
	}

	return result
}

func pruneDirections(x, y, maxX, maxY int) []Vector {
	directions := make([]Vector, 0)
	distance := len(SEARCHWORD) - 1

	if y-distance >= 0 {
		directions = append(directions, NORTH)
	}

	if x+distance < maxX {
		directions = append(directions, EAST)
	}

	if y+distance < maxY {
		directions = append(directions, SOUTH)
	}

	if x-distance >= 0 {
		directions = append(directions, WEST)
	}

	if x-distance >= 0 && y-distance >= 0 {
		directions = append(directions, NORTHWEST)
	}

	if x+distance < maxX && y-distance >= 0 {
		directions = append(directions, NORTHEAST)
	}

	if x+distance < maxX && y+distance < maxY {
		directions = append(directions, SOUTHEAST)
	}

	if x-distance >= 0 && y+distance < maxY {
		directions = append(directions, SOUTHWEST)
	}

	return directions
}

func findWord(grid [][]rune, word string, x, y int, direction Vector) bool {
	if word == "" {
		return true
	}

	return grid[y][x] == rune(word[0]) &&
		findWord(grid, word[1:], x+direction.x, y+direction.y, direction)
}
