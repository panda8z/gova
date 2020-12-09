package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 指向 .class 文件的类路径
type DirEntry struct {
	absDir string
}

// newDirEntry 包内方法，仅供 classpath 包内使用
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

// readClass： 给 DirEntry 添加一个 readClass 方法。 根据函数名，参数列表，返回值列表三个元素我们知道 这个方法实现了 Entry 接口的方法
func (cp *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(cp.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, cp, err
}

// String: 给 DirEntry 添加一个 String 方法。 根据函数名，参数列表，返回值列表三个元素我们知道 这个方法实现了 Entry 接口的方法
// 至此，DirEntry 完整实现了 Entry 接口
func (cp *DirEntry) String() string {
	return cp.absDir
}
