package phptype

type PhpObjectSerialized struct {
	ClassName string
	Data      string
	Value     PhpValue
}

func NewPhpObjectSerialized(className string) *PhpObjectSerialized {
	return &PhpObjectSerialized{
		ClassName: className,
	}
}

// func (self *PhpObjectSerialized) GetClassName() string {
// 	return self.className
// }
//
// func (self *PhpObjectSerialized) SetClassName(name string) *PhpObjectSerialized {
// 	self.className = name
// 	return self
// }
//
// func (self *PhpObjectSerialized) GetData() string {
// 	return self.data
// }
//
// func (self *PhpObjectSerialized) SetData(data string) *PhpObjectSerialized {
// 	self.data = data
// 	return self
// }
//
// func (self *PhpObjectSerialized) GetValue() phptype.PhpValue {
// 	return self.value
// }
//
// func (self *PhpObjectSerialized) SetValue(value phptype.PhpValue) *PhpObjectSerialized {
// 	self.value = value
// 	return self
// }
