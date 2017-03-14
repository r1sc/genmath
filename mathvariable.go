package main

import "math/rand"

type mathVariable struct {
	name string
}

func newMathVariable() *mathVariable {
	newVar := &mathVariable{}
	newVar.mutateValue()
	return newVar
}

func (mvar *mathVariable) clone() mutator {
	return &mathVariable{mvar.name}
}

func (mvar *mathVariable) evaluate(variableValues map[string]float64) float64 {
	return variableValues[mvar.name]
}

func (mvar *mathVariable) mutateValue() {
	mvar.name = variables[rand.Intn(len(variables))]
}

func (mvar *mathVariable) mutate() {
	if chance(100) {
		mvar.mutateValue()
	}
}

func (mvar *mathVariable) toString() string {
	return mvar.name
}

func (mvar *mathVariable) reduce() mutator {
	return mvar
}
