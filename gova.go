package main

import (
	"fmt"
	"github.com/panda8z/gova-jvm/cmd"

	_ "github.com/panda8z/gova-jvm/classfile"
	_ "github.com/panda8z/gova-jvm/classpath"
	"github.com/panda8z/gova-jvm/rtda"
)

func main() {
	c := cmd.ParseCmd()
	if c.VersionFlag {
		fmt.Println(`
version 0.0.1
base go: go version go1.13.1
base java:java version "1.8.0_221"
Java(TM) SE Runtime Environment (build 1.8.0_221-b11)
Java HotSpot(TM) 64-Bit Server VM (build 25.221-b11, mixed mode)
`)
	}
	if c.HelpFlag || c.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(c)
	}
}

func startJVM(cmd *cmd.Cmd) {
	// cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// fmt.Printf("\njrePath: %s \nclasspath: %s  \nclass: %s args: %v\n", cmd.XjreOption, cp, cmd.class, cmd.args)

	// className := strings.Replace(cmd.class, ".", "/", -1)
	// cf := loadClass(className, cp)
	// fmt.Println(cmd.class)
	// printClassInfo(cf)

	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

// func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
// 	classData, _, err := cp.ReadClass(className)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cf, err := classfile.Parse(classData)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return cf
// }

// func printClassInfo(cf *classfile.ClassFile) {
// 	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
// 	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
// 	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
// 	fmt.Printf("this class: %v\n", cf.ClassName())
// 	fmt.Printf("super class: %v\n", cf.SuperClassName())
// 	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
// 	fmt.Printf("fields count: %v\n", len(cf.Fields()))
// 	for _, f := range cf.Fields() {
// 		fmt.Printf(" %s\n", f.Name())
// 	}
// 	fmt.Printf("methods count: %v\n", len(cf.Methods()))
// 	for _, m := range cf.Methods() {
// 		fmt.Printf(" %s\n", m.Name())
// 	}
// }
