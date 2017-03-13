package main

import (
	"fmt"
	"math/rand"
)

const (
	add = iota
	sub
	div
	mul
)

type mathBinaryOp struct {
	op    int
	left  mutator
	right mutator
}

func newOperator(variableValues map[string]float64) *mathBinaryOp {
	op := add + rand.Intn(mul+1)

	var left mutator
	if chance(10) {
		left = newMathVariable(variableValues)
	} else {
		left = newMathValue()
	}

	var right mutator
	if chance(10) {
		right = newMathVariable(variableValues)
	} else {
		right = newMathValue()
	}

	return &mathBinaryOp{op, left, right}
}

func (op *mathBinaryOp) clone() mutator {
	return &mathBinaryOp{op: op.op, left: op.left.clone(), right: op.right.clone()}
}

func (op *mathBinaryOp) evaluate(variableValues map[string]float64) float64 {
	switch op.op {
	case add:
		return op.left.evaluate(variableValues) + op.right.evaluate(variableValues)
	case sub:
		return op.left.evaluate(variableValues) - op.right.evaluate(variableValues)
	case div:
		return op.left.evaluate(variableValues) / op.right.evaluate(variableValues)
	case mul:
		return op.left.evaluate(variableValues) * op.right.evaluate(variableValues)
	}
	panic(fmt.Sprintf("Undefined operation %d\n", op.op))
}

func (op *mathBinaryOp) mutate(variableValues map[string]float64) {
	if chance(100) {
		newOp := newOperator(variableValues)
		newOp.left = op.left
		op.left = newOp
	}
	if chance(100) {
		newOp := newOperator(variableValues)
		newOp.right = op.right
		op.right = newOp
	}
	if chance(100) {
		op.op = add + rand.Intn(mul+1)
	}
	if chance(50) {
		op.left, op.right = op.right, op.left
	}
	if chance(100) {
		op.collapse(variableValues)
	}
	op.left.mutate(variableValues)
	op.right.mutate(variableValues)
}

func (op *mathBinaryOp) toString() string {
	switch op.op {
	case add:
		return fmt.Sprintf("(%s + %s)", op.left.toString(), op.right.toString())
	case sub:
		return fmt.Sprintf("(%s - %s)", op.left.toString(), op.right.toString())
	case div:
		return fmt.Sprintf("(%s / %s)", op.left.toString(), op.right.toString())
	case mul:
		return fmt.Sprintf("(%s * %s)", op.left.toString(), op.right.toString())
	}
	panic(fmt.Sprintf("Undefined operation %d\n", op.op))
}

func (op *mathBinaryOp) collapse(variableValues map[string]float64) {
	left := mathValue(op.left.evaluate(variableValues))
	op.left = &left
	right := mathValue(op.right.evaluate(variableValues))
	op.right = &right
}
