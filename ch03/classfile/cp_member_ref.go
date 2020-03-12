package classfile

type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (cmi *ConstantMemberrefInfo) readInfo(reader *ClassReader)  {
	cmi.classIndex = reader.readUint16()
	cmi.nameAndTypeIndex = reader.readUint16()
}

func (cmi *ConstantMemberrefInfo) ClassName() string {
	return cmi.cp.getClassName(cmi.classIndex)
}

func (cmi *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return cmi.cp.getNameAndType(cmi.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}