package main

type mutator interface {
	clone() mutator
	evaluate(variableValues map[string]float64) float64
	mutate(variableValues map[string]float64)
	toString() string
}
