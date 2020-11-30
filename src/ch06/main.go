package main

import "fmt"
import "strings"
import "jvmgo/src/ch06/classpath"
import "jvmgo/src/ch06/rtda/heap"

/**
	这个里面做的事:
      1. 加载类
	  2. 链接(验证/准备/解析)
      3. 对JVM字节码进行逐步转换成对应的操作(压站出栈等等操作)
*/
func main() {
	/*	cmd := parseCmd()

		if cmd.versionFlag {
			fmt.Println("version 0.0.1")
		} else if cmd.helpFlag || cmd.class == "" {
			printUsage()
		} else {
			startJVM(cmd)
		}*/
	cmd := &Cmd{}
	cmd.XjreOption = "D:\\jdk_1.8\\jre\\"
	cmd.cpOption = "E:\\workspace\\remedy\\target\\classes"
	cmd.class = "com.qdream.Test"
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
