package classpath

import (
	"os"
	"path/filepath"
)

/*

针对 添加结构体方法的时候使用了self或 this的问题，代码检查报警告。

receiver name should be a reflection of its identity;
don't use generic names such as "this" or "self"

*/

// Classpath 定义类路径种类
type Classpath struct {
	bootClasspath   Entry
	extendClasspath Entry
	userClasspath   Entry
}

// Parse 解析所有类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}

	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// parseBootAndExtClasspath 解析 启动类路径 和 扩展类路径
func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibDir := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildcardEntry(jreLibDir)

	// jre/lib/ext/*
	jreExtDir := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extendClasspath = newWildcardEntry(jreExtDir)
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

// ReadClass 根据 给定的 类名 在所有已知的类路径下搜索匹配并读取该类文件。
func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	classPath := className + ".class"

	if data, entry, err := cp.bootClasspath.readClass(classPath); err == nil {
		return data, entry, err
	}

	if data, entry, err := cp.extendClasspath.readClass(classPath); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(classPath)
}

// String 打印类路径
func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not file jre folder! \n")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
