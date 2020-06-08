package conversions

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

// Convert long to double
type L2D struct{ base.NoOperandsInstruction }

func (s *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	value := float64(l)
	stack.PushDouble(value)
}

// Convert long to float
type L2F struct{ base.NoOperandsInstruction }

func (s *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	value := float32(l)
	stack.PushFloat(value)
}

// Convert long to int
type L2I struct{ base.NoOperandsInstruction }

func (s *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	value := int(l)
	stack.PushInt(value)
}
