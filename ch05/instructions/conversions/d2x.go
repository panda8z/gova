package conversions

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (d2i *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stackPopDouble()
	i := int32(d)
	stack.PushInt(i)
}
func (d2i *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stackPopDouble()
	i := int64(d)
	stack.PushInt(i)
}
func (d2i *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stackPopDouble()
	i := float64(d)
	stack.PushInt(i)
}
