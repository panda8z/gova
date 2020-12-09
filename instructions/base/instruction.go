package base

import (
	"github.com/panda8z/gova-jvm/rtda"
)

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (ins *NoOperandsInstruction) FetchOperands(reder *BytecodeReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (ins *BranchInstruction) FetchOperands(reder *BytecodeReader) {
	ins.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (ins *Index8Instruction) FetchOperands(reder *BytecodeReader) {
	ins.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (ins *Index8Instruction) FetchOperands(reder *BytecodeReader) {
	ins.Index = uint(reader.ReadUint16())
}
