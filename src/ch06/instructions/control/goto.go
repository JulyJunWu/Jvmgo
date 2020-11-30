package control

import "jvmgo/src/ch06/instructions/base"
import "jvmgo/src/ch06/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
