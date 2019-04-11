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
