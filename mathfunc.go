package main

import (
	"math"
)

type mathFunc struct {
	op           *mathBinaryOp
	currentError float64
}

func (mf *mathFunc) clone() *mathFunc {
	newFunc := &mathFunc{op: mf.op.clone().(*mathBinaryOp)}
	return newFunc
}

func (mf *mathFunc) evaluate(variables map[string]float64) float64 {
	return mf.op.evaluate(variables)
}

func (mf *mathFunc) mutate(variables []string) {
	if chance(100) {
		newop := newOperator(variables)
		newop.left = mf.op
		mf.op = newop
	}
	if chance(50) {
		mf.op.mutate(variables)
	}
}

func (mf *mathFunc) evolve(variables []string, testcases []testcase) *mathFunc {
	mf.currentError = 0

	clone := mf.clone()
	clone.mutate(variables)

	for _, testcase := range testcases {
		val := mf.evaluate(testcase.variables)
		mf.currentError += math.Abs(val - testcase.expected)

		val = clone.evaluate(testcase.variables)
		clone.currentError += math.Abs(val - testcase.expected)
	}
	if clone.currentError < mf.currentError {
		return clone
	}
	return mf
}

func (mf *mathFunc) toString() string {
	return mf.op.toString()
}
