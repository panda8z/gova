package constants

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

type BIPUSH struct {
	value int8
} // push byte

type SIPUSH struct {
	value 16
} // push short


func (push *BIPUSH) FetchOperands(reader *base.BytecodeReader) { 
	push.value = reader.ReadInt8()
}

func (push *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(push.value)
	frame.OperandStack().PushInt(i) 
}


func (push *SIPUSH) FetchOperands(reader *base.BytecodeReader) { 
	push.value = reader.ReadInt16()
}

func (push *SIPUSH) Execute(frame *rtda.Frame) {
	i := int64(push.value)
	frame.OperandStack().PushInt(i) 
}