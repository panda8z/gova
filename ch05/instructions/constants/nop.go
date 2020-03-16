package constants

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Excute(frame *rtda.Frame) {
	// nothing to
}
