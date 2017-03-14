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

func newOperator() *mathBinaryOp {
	op := add + rand.Intn(mul+1)

	var left mutator
	if chance(100) {
		left = newMathVariable()
	} else {
		v := newMathValue(0)
		v.mutateValue()
		left = v
	}

	var right mutator
	if chance(100) {
		right = newMathVariable()
	} else {
		v := newMathValue(0)
		v.mutateValue()
		right = v
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

func (op *mathBinaryOp) mutate() {
	op.left.mutate()
	op.right.mutate()

	if chance(500) {
		op.left = newOperator()
	}
	if chance(500) {
		op.right = newOperator()
	}
	if chance(500) {
		op.left = newMathValue(0)
	}
	if chance(500) {
		op.right = newMathValue(0)
	}
	if chance(500) {
		op.left = newMathVariable()
	}
	if chance(500) {
		op.right = newMathVariable()
	}
	if chance(5000) {
		newOp := newOperator()
		newOp.left = op.left
		op.left = newOp
	}
	if chance(5000) {
		newOp := newOperator()
		newOp.right = op.right
		op.right = newOp
	}
	if chance(100) {
		op.op = add + rand.Intn(mul+1)
	}
	if chance(100) {
		op.left, op.right = op.right, op.left
	}
	if chance(10) {
		op.reduce()
	}
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

func isMathValue(m mutator) bool {
	switch m.(type) {
	case *mathValue:
		return true
	}
	return false
}

func isMathVariable(m mutator) bool {
	switch m.(type) {
	case *mathVariable:
		return true
	}
	return false
}

func isValue(m mutator, what float64) bool {
	return isMathValue(m) && m.evaluate(nil) == what
}

func isValueSame(left mutator, right mutator) bool {
	if isMathValue(left) && isMathValue(right) {
		return left.evaluate(nil) == right.evaluate(nil)
	}
	return false
}

func (op *mathBinaryOp) reduce() mutator {
	op.left = op.left.reduce()
	op.right = op.right.reduce()

	if isMathValue(op.left) && isMathValue(op.right) {
		return newMathValue(op.evaluate(nil))
	}

	switch op.op {
	case add:
		if isValue(op.left, 0) {
			return op.right
		} else if isValue(op.right, 0) {
			return op.left
		}
	case sub:
		if isValue(op.right, 0) {
			return op.left
		}
	case div:
		if isValue(op.left, 0) {
			return newMathValue(0)
		} else if isValue(op.right, 0) {
			v := newMathValue(1)
			v.mutateValue()
			op.right = v
			return op
		} else if isValue(op.right, 1) {
			return op.left
		} else if isMathVariable(op.left) && isMathVariable(op.right) {
			v1 := op.left.(*mathVariable)
			v2 := op.right.(*mathVariable)
			if v1.name == v2.name {
				return newMathValue(1)
			}
		} else if isValueSame(op.left, op.right) {
			return newMathValue(1)
		}
	case mul:
		if isValue(op.left, 1) {
			return op.right
		} else if isValue(op.right, 1) {
			return op.left
		} else if isValue(op.left, 0) || isValue(op.right, 0) {
			return newMathValue(0)
		}
	}

	return op
}
