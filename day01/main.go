package main

import (
	_ "embed"
    "flag"
    "fmt"
	"log"
    "sort"
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
        "Specifies the section of code to run. 1 for part1, 2 for part2",
    )

    flag.Parse()

    result := 0
    switch *part {
    case 1:
        result = partOne(&input)
    case 2:
        result = partTwo(&input)
    default:
        log.Fatal("'Part' argument out of supported range, needs to be either '1' or '2'")
    }

    fmt.Println("The distance is", result)
}

func partOne(input *string) int {
    listOne := []int{}
    listTwo := []int{}

    for numLine := range strings.SplitSeq(*input, "\n") {
        splitNums := strings.Fields(numLine)

        numOne, err := strconv.Atoi(splitNums[0])
        if err != nil {
            log.Fatal("Input file contains non-numerical values")
        }

        listOne = append(listOne, numOne)

        numTwo, err := strconv.Atoi(splitNums[1])
        if err != nil {
            log.Fatal("Input file contains non-numerical values")
        }

        listTwo = append(listTwo, numTwo)
    }

    sort.Ints(listOne)
    sort.Ints(listTwo)

    diff := 0

    for i := range len(listOne) {
        diff += abs(listOne[i] - listTwo[i])
    }

    return diff
}

func partTwo(input *string) int {
    listOne := []int{}
    setTwo := make(map[int]int)

    for numLine := range strings.SplitSeq(*input, "\n") {
        splitNums := strings.Fields(numLine)

        numOne, err := strconv.Atoi(splitNums[0])
        if err != nil {
            log.Fatal("Input file contains non-numerical values")
        }

        listOne = append(listOne, numOne)

        numTwo, err := strconv.Atoi(splitNums[1])
        if err != nil {
            log.Fatal("Input file contains non-numerical values")
        }

        _, ok := setTwo[numTwo]
        if !ok {
            setTwo[numTwo] = 1
        } else {
            setTwo[numTwo]++
        }
    }

    sum := 0

    for _, num := range listOne {
        count, ok := setTwo[num]
        if !ok {
            continue
        }
        sum += count * num
    }

    return sum
}
