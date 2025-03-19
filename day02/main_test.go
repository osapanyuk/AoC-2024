package main

import (
    "testing"
)

type test struct {
    name    string
    input   string
    want    int
}

var sample = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPartOne(t *testing.T) {
    tests := []test {
        {
            name: "sample",
            input: sample,
            want: 2,
        },
        {
            name: "actual",
            input: input,
            want: 564,
        },
    }

    for _, testRoutine := range tests {
        t.Run(testRoutine.name, func(t *testing.T) {
            t.Parallel()
            if result := partOne(&testRoutine.input); result != testRoutine.want {
                t.Errorf("partOne() = %v, want %v", result, testRoutine.want)
            }
        })
    }
}

func TestPartTwo(t *testing.T) {
    tests := []test {
        {
            name: "sample",
            input: sample,
            want: 4,
        },
        {
            name: "actual",
            input: input,
            want: 604,
        },
    }

    for _, testRoutine := range tests {
        t.Run(testRoutine.name, func(t *testing.T) {
            t.Parallel()
            if result := partTwo(&testRoutine.input); result != testRoutine.want {
                t.Errorf("partTwo() = %v, want %v", result, testRoutine.want)
            }
        })
    }
}
