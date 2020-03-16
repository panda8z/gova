package constants

import (
	"github.com/gova-jvm/instructions/base"
	"github.com/gova-jvm/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

func (cons *ACONST_NULL) Excute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (cons *DCONST_0) Excute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (cons *ICONST_M1) Excute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
