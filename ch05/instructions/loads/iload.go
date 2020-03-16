package loads

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

// Load int from local variable


type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (loads *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(loads.Index))
}

func (loads *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
