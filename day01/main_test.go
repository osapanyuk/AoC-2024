package main

import (
    "testing"
)

func TestPartOneBase(t *testing.T) {
    testInput := `3   4
4   3
2   5
1   3
3   9
3   3`

    want := 11

    result := partOne(&testInput)
    if want != result {
        t.Errorf(`partOne with test input %q, returns %v; want match for %#v`, testInput, result, want)
    }
}

func TestPartOne(t *testing.T) {
    want := 1646452

    result := partOne(&input)
    if want != result {
        t.Errorf(`partOne(&input) returns %v; want match for %#v`, result, want)
    }
}

func TestPartTwoBase(t *testing.T) {
    testInput := `3   4
4   3
2   5
1   3
3   9
3   3`

    want := 31

    result := partTwo(&testInput)
    if want != result {
        t.Errorf(`partTwo with test input %q, returns %v; want match for %#v`, testInput, result, want)
    }
}

func TestPartTwo(t *testing.T) {
    want := 23609874

    result := partTwo(&input)
    if want != result {
        t.Errorf(`PartTwo(&input) returns %v; want match for %#v`, result, want)
    }
}
