package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry 定义型似 .jar  .zip 之类的文件路径
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (cp *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(cp.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, cp, err
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

// String 实现 Entry 接口的打印路径方法。
func (cp *ZipEntry) String() string {
	return cp.absPath
}
