package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
    if len(input) == 0 {
        log.Fatal("Input file empty. Rebuild with populated input.txt file")
    }
    input = strings.TrimRight(input, "\n")
}

func main() {
    part := flag.Int(
        "part",
        0,
        "Specifies the part to be run. '1' for part 1, '2' for part 2",
    )

    flag.Parse()

    result := 0
    switch *part {
    case 1:
        fmt.Println("Running part 1")
        result = partOne(&input)
    case 2:
        fmt.Println("Running part 2")
        result = partTwo(&input)
    default:
        log.Fatal("Invalid part selected. Must be either '1' or '2'")
    }

    fmt.Println("Result:", result)
}

func partOne(input *string) int {
    re := regexp.MustCompile(`mul\(\d*,\d*\)`)
    matches := re.FindAllString(*input, -1)

    total := 0
    for _, match := range matches {
        mult, err := multiply(match)
        if err != nil {
            log.Fatal("Issue with regexp: %v", err)
        }

        total += mult
    }

    return total
}

func partTwo(input *string) int {
    re := regexp.MustCompile(`mul\(\d*,\d*\)|do\(\)|don't\(\)`)
    matches := re.FindAllString(*input, -1)

    total := 0
    doState := true
    for _, match := range matches {
        if match == "do()" {
            doState = true
        } else if match == "don't()" {
            doState = false
        } else if doState {
            mult, err := multiply(match)
            if err != nil {
                log.Fatal("Issue with regexp: %v", err)
            }
            total += mult
        }
    }

    return total
}

func multiply(op string) (int, error) {
    commaIdx := strings.Index(op, ",")

    val1, err := strconv.Atoi(op[4:commaIdx])
    if err != nil {
        return 0, fmt.Errorf("First value is not a number: %v", op[:commaIdx])
    }

    val2, err := strconv.Atoi(op[commaIdx + 1 : len(op) - 1])
    if err != nil {
        return 0, fmt.Errorf("Second value is not a number: %v", op[:commaIdx])
    }

    return val1 * val2, nil
}
