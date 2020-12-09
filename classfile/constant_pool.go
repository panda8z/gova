package classfile

// ConstantPool 定义常量池，其实是一堆 ConstantInfo 组成的切片。
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 注意从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	cpNameAndType := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	cpName := cp.getUtf8(cpNameAndType.nameIndex)
	cpType := cp.getUtf8(cpNameAndType.descriptorIndex)
	return cpName, cpType
}

func (cp ConstantPool) getClassName(index uint16) string {
	cpClassInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(cpClassInfo.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	cpUtf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return cpUtf8Info.str
}
