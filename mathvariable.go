package main

import (
	"math/rand"
	"reflect"
)

type mathVariable struct {
	name string
}

func newMathVariable(variableValues map[string]float64) *mathVariable {
	newVar := &mathVariable{}
	newVar.mutateValue(variableValues)
	return newVar
}

func (mvar *mathVariable) clone() mutator {
	return &mathVariable{mvar.name}
}

func (mvar *mathVariable) evaluate(variableValues map[string]float64) float64 {
	return variableValues[mvar.name]
}

func (mvar *mathVariable) mutateValue(variableValues map[string]float64) {
	keys := reflect.ValueOf(variableValues).MapKeys()
	mvar.name = keys[rand.Intn(len(keys))].String()
}

func (mvar *mathVariable) mutate(variableValues map[string]float64) {
	if chance(100) {
		mvar.mutateValue(variableValues)
	}
}

func (mvar *mathVariable) toString() string {
	return mvar.name
}
