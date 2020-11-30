package main

import "fmt"
import "jvmgo/src/ch06/instructions"
import "jvmgo/src/ch06/instructions/base"
import "jvmgo/src/ch06/rtda"
import "jvmgo/src/ch06/rtda/heap"

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		//fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		//panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		// 指令 (10进制,需要转成对应的16进制,然后从 https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.dup 查找)
		// 每次读取一个字节的指令
		opcode := reader.ReadUint8()
		fmt.Printf("字节码指令:%x \n", opcode)
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
