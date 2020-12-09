package math
/*
位移指令
1. int 左移
2. int 右移
3. int 逻辑右移
4. long 左移
5. long 右移
6. long 逻辑右移

我没明白逻辑右移
我也没明白 Pop出来的数值类型不同。
s 的 计算方式也不同是为什么。
*/
import (
	"math"

	"github.com/panda8z/gova-jvm/instructions/base"
	"github.com/panda8z/gova-jvm/rtda"
)

// int left move
type ISHL sturct{ base.NoOperandsInstruction } 

func (sh *ISHL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(reslut)
}

// int right move
type ISHR sturct{ base.NoOperandsInstruction } 

func (sh *ISHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(reslut)
}

// int logic right move
type IUSHR sturct{ base.NoOperandsInstruction } 

func (sh *LSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLang(reslut)
}

// long left move
type LSHL sturct{ base.NoOperandsInstruction } 

func (sh *IUSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f 
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// long right move
type LSHR sturct{ base.NoOperandsInstruction } 

func (sh *LSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushInt(reslut)
}

// long logic right move
type LUSHR sturct{ base.NoOperandsInstruction } 

func (sh *LUSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64( uint64(v1) >> s)
	stack.PushLang(reslut)
}
