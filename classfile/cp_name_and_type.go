package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cnati *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cnati.nameIndex = reader.readUint16()
	cnati.descriptorIndex = reader.readUint16()
}
