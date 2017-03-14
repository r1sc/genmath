package main

import (
	"fmt"
	"math/rand"
)

type mathValue float64

func newMathValue(value float64) *mathValue {
	v := mathValue(value)
	return &v
}

func (value *mathValue) clone() mutator {
	newVal := mathValue(*value)
	return &newVal
}

func (value *mathValue) evaluate(variables map[string]float64) float64 {
	return float64(*value)
}

func (value *mathValue) mutateValue() {
	*value = mathValue((rand.Float64() - 0.5) * 10)
}

func (value *mathValue) mutate() {
	if chance(100) {
		value.mutateValue()
	}
	if chance(10) {
		*value = mathValue(int(*value))
	}
}

func (value *mathValue) toString() string {
	return fmt.Sprintf("%f", float64(*value))
}

func (value *mathValue) reduce() mutator {
	return value
}
