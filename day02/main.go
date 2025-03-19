package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
    if len(input) == 0 {
        log.Fatal("Input file is empty... Rebuild is necessary")
    }
    input = strings.TrimRight(input, "\n")
}

func abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}

func main() {
    part := flag.Int(
        "part",
        0,
        "Specifies which part to run. 1 for Part1, 2 for Part2",
    )

    flag.Parse()

    result := 0
    switch *part {
    case 1:
        fmt.Println("Executing Part 1")
        result = partOne(&input)
    case 2:
        fmt.Println("Executing Part 2")
        result = partTwo(&input)
    default:
        log.Fatal("'Part' argument out of supported range, needs to be either '1' or '2'")
    }

    fmt.Println("Result:", result)
}

func partOne(input *string) int {
    count := 0

    for line := range strings.SplitSeq(*input, "\n") {
        reportList, err := parseReportRow(&line)

        if err != nil {
            log.Fatal(err)
        }

        if index := validateRow(reportList); index == 0 {
            count++
        }
    }

    return count
}

func partTwo(input *string) int {
    count := 0
    for line := range strings.SplitSeq(*input, "\n") {
        reportList, err := parseReportRow(&line)

        if err != nil {
            log.Fatal(err)
        }

        isUnsafe := false

        // If we have an invalid row, re-validate the row by removing a single value
        if index := validateRow(reportList); index != 0 {
            copyRow := make([]int, len(reportList))
            copy(copyRow, reportList)

            // Case 1: first value in the row is removed.
            // This is done because the first two values define increment/decrement
            subResult1 := validateRow(reportList[1:])

            // Case 2: remove value at index
            reportList = slices.Delete(reportList, index, index + 1)

            // Case 3: remove value at index - 1
            copyRow = slices.Delete(copyRow, index - 1, index)

            subResult2 := validateRow(reportList)
            subResult3 := validateRow(copyRow)

            if subResult1 != 0 && subResult2 != 0 && subResult3 != 0 {
                isUnsafe = true
            }
        }

        if !isUnsafe {
            count++
        }
    }

    return count
}

func parseReportRow(reportRow *string) ([]int, error) {
    var reportList []int

    for _, elem := range strings.Fields(*reportRow) {
        num, err := strconv.Atoi(elem)

        if err != nil {
            return []int{}, fmt.Errorf("Input file contains non-integer value: %w", err)
        }

        reportList = append(reportList, num)
    }

    if len(reportList) < 2 {
        return []int{}, fmt.Errorf("Input file is missing the minimum number of columns of 2")
    }

    return reportList, nil
}

func validateRow(row []int) int {
        isIncrementing := row[0] < row[1]

        for i:= 1; i < len(row); i++ {

            if isIncrementing && row[i - 1] > row[i] {
                return i
            } else if !isIncrementing && row[i - 1] < row[i] {
                return i
            }

            diff := abs(row[i - 1] - row[i])

            if diff > 3 || diff < 1 {
                return i
            }
        }

        return 0
}