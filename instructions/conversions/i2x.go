package conversions

import "github.com/panda8z/gova-jvm/rtda"

// Convert int to byte
type I2B struct{ base.NoOperandInstruction }

func (s *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// Convert int to char
type I2C struct{ base.NoOperandsInstruction }

func (s *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(uint16(i))
	stack.PushInt(b)
}

// Convert int to short
type I2S struct{ base.NoOperandsInstruction }

func (s *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int16(i))
	stack.PushInt(b)
}

// Convert int to long
type I2L struct{ base.NoOperandsInstruction }

func (s *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int64(i)
	stack.PushLong(b)
}

// Convert int to float
type I2F struct{ base.NoOperandsInstruction }

func (s *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := float32(i)
	stack.PushFloat(b)
}

// Convert int to double
type I2F struct{ base.NoOperandsInstruction }

func (s *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := float64(i)
	stack.PushDouble(b)
}
