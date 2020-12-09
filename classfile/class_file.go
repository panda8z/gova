package classfile

import "fmt"

// ClassFile 定义 class 文件结构。
// 因为文件结构是固定的，所以所有建立在 class 文件上的读取方法调用都是有代码书写的固定顺序的。
type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse 把[]ClassFile
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		// Go语言没有异常处理机制，只有一个panic-recover机制。
		// 一旦 panic 我们就 recover 但是，如果recover失败也要打出错误❌信息。
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// read 依次调用其他各个字段的方法解析class文件，最后的到完整的 ClassFile 结构体
// 这里的方法发调用是有顺序的。
func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)               // todo
	cf.readAndCheckVersion(reader)             // todo
	cf.constantPool = readConstantPool(reader) // todo
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool) // todo
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool) // todo
}

// readAndCheckVersion 读取版本信息并校验
// 两个版本信息的读取是固定顺序的。
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return 
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return 
		}
		
	}
	panic("java.lang.UnsupportedClassVersionError")
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// MinorVersion getter
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// MagorVersion getter
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// ConstantPool getter
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// AccessFlags getter
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

// Fields getter
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

// Methods getter
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

// ClassName getter
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

// SuperClassName getter
func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return "" // 只有Object类没有父类
}

// InterfaceNames getter
func (cf *ClassFile) InterfaceNames() []string {
	interfacesNames := make([]string, len(cf.interfaces))
	for index, name := range cf.interfaces {
		interfacesNames[index] = cf.constantPool.getClassName(name)
	}
	return interfacesNames
}
