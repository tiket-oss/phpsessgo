package phptype

type ObjectSerialized struct {
	ClassName string
	Data      string
	Value     Value
}

func NewObjectSerialized(className string) *ObjectSerialized {
	return &ObjectSerialized{
		ClassName: className,
	}
}

// func (self *ObjectSerialized) GetClassName() string {
// 	return self.className
// }
//
// func (self *ObjectSerialized) SetClassName(name string) *ObjectSerialized {
// 	self.className = name
// 	return self
// }
//
// func (self *ObjectSerialized) GetData() string {
// 	return self.data
// }
//
// func (self *ObjectSerialized) SetData(data string) *ObjectSerialized {
// 	self.data = data
// 	return self
// }
//
// func (self *ObjectSerialized) GetValue() phptype.Value {
// 	return self.value
// }
//
// func (self *ObjectSerialized) SetValue(value phptype.Value) *ObjectSerialized {
// 	self.value = value
// 	return self
// }
