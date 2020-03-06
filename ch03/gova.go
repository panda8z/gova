package main

import (
	"fmt"
	"github.com/gova-jvm/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println(`
version 0.0.1
base go: go version go1.13.1
base java:java version "1.8.0_221"
Java(TM) SE Runtime Environment (build 1.8.0_221-b11)
Java HotSpot(TM) 64-Bit Server VM (build 25.221-b11, mixed mode)
`)
	}
	if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("\njrePath: %s \nclasspath: %s  \nclass: %s args: %v\n", cmd.XjreOption, cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not load or found main class from %s\n", cmd.class)
		return
	}

	fmt.Printf("classData: %v\n", classData)

}
