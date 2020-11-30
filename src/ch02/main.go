package main

import (
	"fmt"
)
import "strings"
import "jvmgo/src/ch02/classpath"

func main() {
	cmd := &Cmd{}
	cmd.XjreOption = "D:\\jdk_1.8\\jre\\"
	cmd.cpOption = "E:\\workspace\\remedy\\target\\classes"
	cmd.class = "com.qdream.Test"
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}
