package rtda

import "jvmgo/src/ch06/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
