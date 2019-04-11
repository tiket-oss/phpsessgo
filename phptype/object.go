package phptype

type PhpObject struct {
	ClassName string
	Members   PhpArray
}

func NewPhpObject(className string) *PhpObject {
	return &PhpObject{
		ClassName: className,
		Members:   PhpArray{},
	}
}

//
// func (self *PhpObject) GetClassName() string {
// 	return self.ClassName
// }
//
// func (self *PhpObject) SetClassName(name string) *PhpObject {
// 	self.ClassName = name
// 	return self
// }
//
// func (self *PhpObject) GetMembers() PhpArray {
// 	return self.Members
// }
//
// func (self *PhpObject) SetMembers(members PhpArray) *PhpObject {
// 	self.Members = members
// 	return self
// }

func (self *PhpObject) GetPrivate(name string) (v PhpValue, ok bool) {
	v, ok = self.Members["\x00"+self.ClassName+"\x00"+name]
	return
}

func (self *PhpObject) SetPrivate(name string, value PhpValue) *PhpObject {
	self.Members["\x00"+self.ClassName+"\x00"+name] = value
	return self
}

func (self *PhpObject) GetProtected(name string) (v PhpValue, ok bool) {
	v, ok = self.Members["\x00*\x00"+name]
	return
}

func (self *PhpObject) SetProtected(name string, value PhpValue) *PhpObject {
	self.Members["\x00*\x00"+name] = value
	return self
}

func (self *PhpObject) GetPublic(name string) (v PhpValue, ok bool) {
	v, ok = self.Members[name]
	return
}

func (self *PhpObject) SetPublic(name string, value PhpValue) *PhpObject {
	self.Members[name] = value
	return self
}
