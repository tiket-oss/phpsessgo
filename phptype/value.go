package phptype

type PhpValue interface{}

type PhpArray map[PhpValue]PhpValue

type PhpSlice []PhpValue
