package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// u1类型数据
func (cr *ClassReader) readUint8() uint8 {
	value = cr.data[0]    // 字节数组的 第一个就是8位
	cr.data = cr.data[1:] // 字节数组剩余部分
	return value
}
// u2类型数据
func (cr *ClassReader) readUint16() uint16 {
	value = binary.BigEndian.Uint16(cr.data) // 用Go标准库读取多字节。
	cr.data = cr.data[2:]                    // 字节数组剩余部分。
	return value
}

// u4类型数据
func (cr *ClassReader) readUint32() uint32 {
	value = binary.BigEndian.Uint32(cr.data) // 用Go标准库读取多字节。
	cr.data = cr.data[4:]                    // 字节数组剩余部分。
	return value
}

// 读取uint64(Java虚拟机规范并没有定义u8)类型数据
func (cr *ClassReader) readUint64() uint64 {
	value = binary.BigEndian.Uint64(cr.data) // 用Go标准库读取多字节。
	cr.data = cr.data[8:]                    // 字节数组剩余部分。
	return value
}

// 读取uint16表，表的大小由开头的uint16数据指出
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	s := make([]uint16, n)
	// should omit 2nd value from range; this loop is equivalent to `for index := range ...`
	for index := range s {
		s[index] = cr.readUint16()
	}
	return s
}

// 读取指定数量的字节
func (cr *ClassReader) readBytes(n uint32) []byte {
	data := cr.data[:n]
	cr.data = cr.data[n:]
	return data
}
