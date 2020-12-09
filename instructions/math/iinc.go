package math

import (
	"github.com/panda8z/gova-jvm/instructions/base"
	"github.com/panda8z/gova-jvm/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (iinc *IINC) FetchOperands(reader *base.BytecodeReader) {
	iinc.Index = uint(reader.ReadUint8())
	iinc.Const = int32(reader.ReadUint8())
}

func (iinc *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(iinc.Index)
	val += iinc.Const
	localVars.SetInt(iinc.Index, val)
}
