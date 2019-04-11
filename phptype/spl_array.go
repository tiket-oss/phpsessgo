package phptype

type PhpSplArray struct {
	Flags      int
	Array      PhpValue
	Properties PhpValue
}

func NewPhpSplArray(array, properties PhpValue) *PhpSplArray {
	if array == nil {
		array = make(PhpArray)
	}

	if properties == nil {
		properties = make(PhpArray)
	}

	return &PhpSplArray{
		Array:      array,
		Properties: properties,
	}
}

//
// func (self *PhpSplArray) GetFlags() int {
// 	return self.flags
// }
//
// func (self *PhpSplArray) SetFlags(value int) {
// 	self.flags = value
// }
//
// func (self *PhpSplArray) GetArray() phptype.PhpValue {
// 	return self.array
// }
//
// func (self *PhpSplArray) SetArray(value phptype.PhpValue) {
// 	self.array = value
// }
//
// func (self *PhpSplArray) GetProperties() phptype.PhpValue {
// 	return self.properties
// }
//
// func (self *PhpSplArray) SetProperties(value phptype.PhpValue) {
// 	self.properties = value
// }
