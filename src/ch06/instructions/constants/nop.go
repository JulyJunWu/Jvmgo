package constants

import "jvmgo/src/ch06/instructions/base"
import "jvmgo/src/ch06/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
