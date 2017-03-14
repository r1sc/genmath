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

var variables []string

func main() {
	variables = []string{"x"}
	testcases := []testcase{
		testcase{variables: map[string]float64{"x": 0}, expected: 0},
		testcase{variables: map[string]float64{"x": 1}, expected: 0},
		testcase{variables: map[string]float64{"x": 2}, expected: 1},
		testcase{variables: map[string]float64{"x": 3}, expected: 3},
		testcase{variables: map[string]float64{"x": 4}, expected: 6}}

	rand.Seed(time.Now().Unix())
	tree := &mathFunc{op: newOperator()}
	lastError := math.MaxFloat64

	organism := 1
	generations := 0
	timeToLive := 20000

	for {
		generations++
		tree = tree.evolve(testcases)
		if tree.currentError < lastError {
			fmt.Printf("Organism %d, generation %d, error=%f => %s\n", organism, generations, tree.currentError, tree.toString())
		}
		lastError = tree.currentError

		if lastError == 0 {
			break
		}
		if generations > timeToLive {
			fmt.Println("Organism has become extinct")
			tree = &mathFunc{op: newOperator()}
			generations = 0
			organism++
			if organism > 10 {
				timeToLive = int(float64(timeToLive) * 2)
				fmt.Printf("Too many organisms, adjusting time to live to %d\n", timeToLive)
				organism = 0
			}
		}
	}

	tree = tree.reduce()
	fmt.Printf("\nFound solution  => %s\n", tree.toString())
	fmt.Println("Proof:")
	for _, testcase := range testcases {
		val := tree.evaluate(testcase.variables)
		fmt.Printf("When")
		for k, v := range testcase.variables {
			fmt.Printf(" %s=%f", k, v)
		}
		fmt.Printf(" %s = %f. Expected %f\n", tree.toString(), val, testcase.expected)
	}

}
