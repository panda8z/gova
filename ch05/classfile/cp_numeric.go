package classfile

import (
	"math"
)

// ConstantIntegerInfo
type ConstantIntegerInfo struct {
	value int32
}

// readInfo 这是最终被调用读取并转换的地方，先读取一个uint32数据，然后把它转型成int32类型。
func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cii.value = int32(bytes)
}

// ConstantFloatInfo
type ConstantFloatInfo struct {
	value float32
}

func (cfi *ConstantFloatInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint32()
	cfi.value = math.Float32frombits(bytes)
}


// ConstantLoneInfo
type ConstantLongInfo struct {
	value int64
}

func (cli *ConstantLongInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint64()
	cli.value = int64(bytes)
}

// ConstantDoubleInfo
type ConstantDoubleInfo struct {
	value float64
}

func (cli *ConstantDoubleInfo) readInfo(reader *ClassReader)  {
	bytes := reader.readUint64()
	cli.value = math.Float64frombits(bytes)
}