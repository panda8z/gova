package constants

import (
	"github.com/panda8z/gova-jvm/instructions/base"
	"github.com/panda8z/gova-jvm/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Excute(frame *rtda.Frame) {
	// nothing to
}
