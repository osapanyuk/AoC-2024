package main

import (
	_ "embed"
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
    listOne := []int{}
    listTwo := []int{}
    for _, numLine := range strings.Split(input, "\n") {
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

    fmt.Println("The distance is ", diff)
}
