package rtda

import "math"

type OperandStack struct {
	size  int
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (oStack *OperandStack) PushInt(value int32) {
	oStack.slots[oStack.size].num = value
	oStack.size++
}

func (oStack *OperandStack) PopInt() int32 {
	oStack.size--
	return oStack.slots[oStack.size].num
}

func (oStack *OperandStack) PushFloat(value float32) {
	bits := math.Float32bits(value)
	oStack.slots[oStack.size].num = int32(bits)
	oStack.size++
}

func (oStack *OperandStack) PopFloat() float32 {
	oStack.size--
	bits := uint32(oStack.slots[oStack.size].num)
	return math.Float32frombits(bits)
}

func (oStack *OperandStack) PushLong(value int64) {
	oStack.slots[oStack.size].num = int32(value)
	oStack.slots[oStack.size+1].num = int32(value >> 32)
	oStack.size += 2
}

func (oStack *OperandStack) PopLong() int64 {
	oStack.size -= 2
	low := uint32(oStack.slots[oStack.size].num)
	high := uint32(oStack.slots[oStack.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (oStack *OperandStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	oStack.PushLong(int64(bits))
}

func (oStack *OperandStack) PopDouble() float64 {
	bits := uint64(oStack.PopLong())
	return math.Float64frombits(bits)
}

func (oStack *OperandStack) PushRef(ref *Object) {
	oStack.slots[oStack.size].ref = ref
	oStack.size++
}

func (oStack *OperandStack) PopRef() *Object {
	oStack.size--
	ref := oStack.slots[oStack.size].ref
	oStack.slots[oStack.size].ref = nil
	return ref
}
