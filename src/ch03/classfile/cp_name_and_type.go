package classfile

import "fmt"

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) printInfo() {
	fmt.Print(" -> #nameIndex=", self.nameIndex)
	fmt.Println(" #descriptorIndex=", self.descriptorIndex)
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
	//fmt.Println("字段名或方法名:",self.nameIndex,",以及对应的类型(参数类型:返回类型):",self.descriptorIndex)
}
