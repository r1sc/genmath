package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type testcase struct {
	variables map[string]float64
	expected  float64
}

func main() {
	variables := []string{"x"}
	testcases := []testcase{
		testcase{variables: map[string]float64{"x": 1}, expected: 2},
		testcase{variables: map[string]float64{"x": 2}, expected: 3},
		testcase{variables: map[string]float64{"x": 3}, expected: 4}}

	rand.Seed(time.Now().Unix())
	tree := &mathFunc{op: newOperator(variables)}
	lastError := math.MaxFloat64
	for {
		tree = tree.evolve(variables, testcases)
		if tree.currentError < lastError {
			fmt.Printf("Error=%f of %s\n", tree.currentError, tree.toString())
		}
		lastError = tree.currentError

		if lastError == 0 {
			break
		}
	}
	fmt.Println("Found solution")
}
