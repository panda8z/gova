package stack

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

// SWAP ： swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

// Execute 交换栈顶两个元素
func Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
