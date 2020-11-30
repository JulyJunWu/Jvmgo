package rtda

import "jvmgo/src/ch06/rtda/heap"

// stack frame
type Frame struct {
	lower *Frame // stack is implemented as linked list
	// 局部变量
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
	// 持有的线程
	thread *Thread
	// 这个栈帧对应的方法
	method *heap.Method
	nextPC int // the next instruction after the call
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
