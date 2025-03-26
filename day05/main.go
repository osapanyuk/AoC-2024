package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	if input == "" {
		log.Fatal("Input string empty")
	}
	input = strings.TrimRight(input, "\n")
}

func main() {
	result := partOne(&input)

	fmt.Println("Result:", result)
}

func partOne(input *string) int {

	instructions := strings.Split(*input, "\n\n")

    ruleList := strings.Split(instructions[0], "\n")

    rules := make(map[int][]int, len(ruleList))
    for _, rule := range ruleList {
        ruleSplit := strings.Split(rule, "|")
        left, _ := strconv.Atoi(ruleSplit[0])
        right, _ := strconv.Atoi(ruleSplit[1])
        rules[left] = append(rules[left], right)
    }

    orderLists := strings.Split(instructions[1], "\n")

    orderList := make([][]int, len(orderLists))
    for i, order := range orderLists {
        for value := range strings.SplitSeq(order, ",") {
            val, _ := strconv.Atoi(value)
            orderList[i] = append(orderList[i], val)
        }
    }

    println("part 1", len(ruleList))
    println("part 2", len(orderLists))

    result := 0
    for x, update := range orderList {
        ordered := true
        for i, value := range update {
            for _, checkVal := range update[:i] {
                if slices.Contains(rules[value], checkVal) {
                    ordered = false
                    break
                }
            }
            if !ordered {
                break
            }
        }
        if ordered {
            fmt.Println(x, update[len(update)/2])
            result += update[len(update)/2]
        }
    }
	return result
}
