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
	variables = []string{"x", "y"}
	testcases := []testcase{
		testcase{variables: map[string]float64{"x": 0, "y": 0}, expected: 0},
		testcase{variables: map[string]float64{"x": 1, "y": 1}, expected: 2},
		testcase{variables: map[string]float64{"x": 2, "y": 4}, expected: 6},
		testcase{variables: map[string]float64{"x": 3, "y": 9}, expected: 12}}

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
			fmt.Printf("Organism %d, generation %d, error=%f\n", organism, generations, tree.currentError)
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
				timeToLive = int(float64(timeToLive) * 1.5)
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
