package math

/*
布尔运算的与运算

int 按位与

long 按位与

1. 去除两个值
2. 用go语言的按位与后
3. 结果压进栈里
*/
import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

// Boolean AND int
type IAND struct{ base.NoOperandsIstruction }

func (and *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOperandsIstruction }

func (and *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
