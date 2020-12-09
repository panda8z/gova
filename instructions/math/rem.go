package math

/*
取余运算指令

int 值的取余运算指令
long 值的取余运算指令
float 值的取余运算指令
double 值的取余运算指令

等一下 用jvm规范的思想来说
short char 就没有取余的权利了？

*/
import (
	"math"

	"github.com/panda8z/gova-jvm/instructions/base"
	"github.com/panda8z/gova-jvm/rtda"
)

// IREM
type IREM struct{ base.NoOperandsInstruction }

func (rem *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := staxk.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

// LREM
type LREM struct{ base.NoOperandsInstruction }

func (rem *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := staxk.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

// FREM
type FREM struct{ base.NoOperandsInstruction }

func (rem *fREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := staxk.PopFloat()
	result := math.Mod(v1, v2)
	stack.PushFloat(result)
}

// DREM
type DREM struct{ base.NoOperandsInstruction }

func (rem *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := staxk.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
