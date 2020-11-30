package constants

import "jvmgo/src/ch05/instructions/base"
import "jvmgo/src/ch05/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
