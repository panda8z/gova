package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// Entry classpath的入口协议定义。
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// newEntry 包内方法，仅供本包内部使用。
func newEntry(path string) Entry {

	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||strings.HasSuffix(path, ".zip") ||strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
