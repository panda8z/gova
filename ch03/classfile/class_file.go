package classfile

import "fmt"

// ClassFile 定义 class 文件结构。
type ClassFile struct {
	//magic uint32
	minorVersion uint16
	magorVersion uint16
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
	cf := &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {

}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

// MinorVersion getter
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// MagorVersion getter
func (cf *ClassFile) MagorVersion() uint16 {
	return cf.magorVersion
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

}

// SuperClassName getter
func (cf *ClassFile) SuperClassName() string {

}
