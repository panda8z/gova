package stack

import (
	"github.com/panda8z/gova-jvm/instructions/base"
	"github.com/panda8z/gova-jvm/rtda"
)

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (dup *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stackPopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
