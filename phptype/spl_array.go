package phptype

type PhpSplArray struct {
	Flags      int
	Array      Value
	Properties Value
}

func NewPhpSplArray(array, properties Value) *PhpSplArray {
	if array == nil {
		array = make(Array)
	}

	if properties == nil {
		properties = make(Array)
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
// func (self *PhpSplArray) GetArray() phptype.Value {
// 	return self.array
// }
//
// func (self *PhpSplArray) SetArray(value phptype.Value) {
// 	self.array = value
// }
//
// func (self *PhpSplArray) GetProperties() phptype.Value {
// 	return self.properties
// }
//
// func (self *PhpSplArray) SetProperties(value phptype.Value) {
// 	self.properties = value
// }
