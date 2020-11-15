package main

import "errors"

// ErrInvalidB 除数为0
var ErrInvalidB = errors.New("invalid b: zero")

// Operation 支持的操作
type Operation int

// 操作列表
const (
	OpAdd Operation = iota
	OpSub
	OpMul
	OpDiv
)

// Calculator 计算器
type Calculator struct{}

// Do 计算机进行计算
func (c Calculator) Do(op Operation, a, b int) (int, error) {
	switch op {
	case OpAdd:
		return c.add(a, b), nil
	case OpSub:
		return c.sub(a, b), nil
	case OpMul:
		return c.mul(a, b), nil
	case OpDiv:
		return c.div(a, b)
	default:
	}

	// unknown op?
	return 0, errors.New("unknown error")
}

func (Calculator) add(a, b int) int { return a + b }

func (Calculator) sub(a, b int) int { return a - b }

func (Calculator) mul(a, b int) int { return a * b }

func (Calculator) div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrInvalidB
	}

	return a / b, nil
}
