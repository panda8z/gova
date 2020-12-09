package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (ea *ExceptionsAttribute) readInfo(reader *ClassReader) {
	ea.exceptionIndexTable = reader.readUint16s()
}

func (ea *ExceptionsAttribute) ExceptionIndexTable() []uint16  {
	return ea.exceptionIndexTable
}