package phptype

type Object struct {
	ClassName string
	Members   Array
}

func NewObject(className string) *Object {
	return &Object{
		ClassName: className,
		Members:   Array{},
	}
}

//
// func (self *Object) GetClassName() string {
// 	return self.ClassName
// }
//
// func (self *Object) SetClassName(name string) *Object {
// 	self.ClassName = name
// 	return self
// }
//
// func (self *Object) GetMembers() Array {
// 	return self.Members
// }
//
// func (self *Object) SetMembers(members Array) *Object {
// 	self.Members = members
// 	return self
// }

func (self *Object) GetPrivate(name string) (v Value, ok bool) {
	v, ok = self.Members["\x00"+self.ClassName+"\x00"+name]
	return
}

func (self *Object) SetPrivate(name string, value Value) *Object {
	self.Members["\x00"+self.ClassName+"\x00"+name] = value
	return self
}

func (self *Object) GetProtected(name string) (v Value, ok bool) {
	v, ok = self.Members["\x00*\x00"+name]
	return
}

func (self *Object) SetProtected(name string, value Value) *Object {
	self.Members["\x00*\x00"+name] = value
	return self
}

func (self *Object) GetPublic(name string) (v Value, ok bool) {
	v, ok = self.Members[name]
	return
}

func (self *Object) SetPublic(name string, value Value) *Object {
	self.Members[name] = value
	return self
}
