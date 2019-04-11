package phpserialize

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/imantung/phpsessgo/phptype"
)

func Serialize(v phptype.Value) (string, error) {
	encoder := NewSerializer()
	encoder.SetSerializedEncodeFunc(SerializedEncodeFunc(Serialize))
	return encoder.Encode(v)
}

type Serializer struct {
	lastErr    error
	encodeFunc SerializedEncodeFunc
}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (self *Serializer) SetSerializedEncodeFunc(f SerializedEncodeFunc) {
	self.encodeFunc = f
}

func (self *Serializer) Encode(v phptype.Value) (string, error) {
	var value bytes.Buffer

	switch t := v.(type) {
	default:
		self.saveError(fmt.Errorf("phpserialize: Unknown type %T with value %#v", t, v))
	case nil:
		value = self.encodeNull()
	case bool:
		value = self.encodeBool(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		value = self.encodeNumber(v)
	case string:
		value = self.encodeString(v, DELIMITER_STRING_LEFT, DELIMITER_STRING_RIGHT, true)
	case phptype.Array, map[phptype.Value]phptype.Value, phptype.Slice:
		value = self.encodeArray(v, true)
	case *phptype.Object:
		value = self.encodeObject(v)
	case *phptype.ObjectSerialized:
		value = self.encodeSerialized(v)
	case *phptype.PhpSplArray:
		value = self.encodeSplArray(v)
	}

	return value.String(), self.lastErr
}

func (self *Serializer) encodeNull() (buffer bytes.Buffer) {
	buffer.WriteRune(TOKEN_NULL)
	buffer.WriteRune(SEPARATOR_VALUES)
	return
}

func (self *Serializer) encodeBool(v phptype.Value) (buffer bytes.Buffer) {
	buffer.WriteRune(TOKEN_BOOL)
	buffer.WriteRune(SEPARATOR_VALUE_TYPE)

	if bVal, ok := v.(bool); ok && bVal == true {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}

	buffer.WriteRune(SEPARATOR_VALUES)
	return
}

func (self *Serializer) encodeNumber(v phptype.Value) (buffer bytes.Buffer) {
	var val string

	isFloat := false

	switch v.(type) {
	default:
		val = "0"
	case int:
		intVal, _ := v.(int)
		val = strconv.FormatInt(int64(intVal), 10)
	case int8:
		intVal, _ := v.(int8)
		val = strconv.FormatInt(int64(intVal), 10)
	case int16:
		intVal, _ := v.(int16)
		val = strconv.FormatInt(int64(intVal), 10)
	case int32:
		intVal, _ := v.(int32)
		val = strconv.FormatInt(int64(intVal), 10)
	case int64:
		intVal, _ := v.(int64)
		val = strconv.FormatInt(int64(intVal), 10)
	case uint:
		intVal, _ := v.(uint)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint8:
		intVal, _ := v.(uint8)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint16:
		intVal, _ := v.(uint16)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint32:
		intVal, _ := v.(uint32)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint64:
		intVal, _ := v.(uint64)
		val = strconv.FormatUint(uint64(intVal), 10)
	// PHP has precision = 17 by default
	case float32:
		floatVal, _ := v.(float32)
		val = strconv.FormatFloat(float64(floatVal), FORMATTER_FLOAT, FORMATTER_PRECISION, 32)
		isFloat = true
	case float64:
		floatVal, _ := v.(float64)
		val = strconv.FormatFloat(float64(floatVal), FORMATTER_FLOAT, FORMATTER_PRECISION, 64)
		isFloat = true
	}

	if isFloat {
		buffer.WriteRune(TOKEN_FLOAT)
	} else {
		buffer.WriteRune(TOKEN_INT)
	}

	buffer.WriteRune(SEPARATOR_VALUE_TYPE)
	buffer.WriteString(val)
	buffer.WriteRune(SEPARATOR_VALUES)

	return
}

func (self *Serializer) encodeString(v phptype.Value, left, right rune, isFinal bool) (buffer bytes.Buffer) {
	val, _ := v.(string)

	if isFinal {
		buffer.WriteRune(TOKEN_STRING)
	}

	buffer.WriteString(self.prepareLen(len(val)))
	buffer.WriteRune(left)
	buffer.WriteString(val)
	buffer.WriteRune(right)

	if isFinal {
		buffer.WriteRune(SEPARATOR_VALUES)
	}

	return
}

func (self *Serializer) encodeArray(v phptype.Value, isFinal bool) (buffer bytes.Buffer) {
	var (
		arrLen int
		s      string
	)

	if isFinal {
		buffer.WriteRune(TOKEN_ARRAY)
	}

	switch v.(type) {
	case phptype.Array:
		arrVal, _ := v.(phptype.Array)
		arrLen = len(arrVal)

		buffer.WriteString(self.prepareLen(arrLen))
		buffer.WriteRune(DELIMITER_OBJECT_LEFT)

		for k, v := range arrVal {
			s, _ = self.Encode(k)
			buffer.WriteString(s)
			s, _ = self.Encode(v)
			buffer.WriteString(s)
		}

	case map[phptype.Value]phptype.Value:
		arrVal, _ := v.(map[phptype.Value]phptype.Value)
		arrLen = len(arrVal)

		buffer.WriteString(self.prepareLen(arrLen))
		buffer.WriteRune(DELIMITER_OBJECT_LEFT)

		for k, v := range arrVal {
			s, _ = self.Encode(k)
			buffer.WriteString(s)
			s, _ = self.Encode(v)
			buffer.WriteString(s)
		}
	case phptype.Slice:
		arrVal, _ := v.(phptype.Slice)
		arrLen = len(arrVal)

		buffer.WriteString(self.prepareLen(arrLen))
		buffer.WriteRune(DELIMITER_OBJECT_LEFT)

		for k, v := range arrVal {
			s, _ = self.Encode(k)
			buffer.WriteString(s)
			s, _ = self.Encode(v)
			buffer.WriteString(s)
		}
	}

	buffer.WriteRune(DELIMITER_OBJECT_RIGHT)

	return
}

func (self *Serializer) encodeObject(v phptype.Value) (buffer bytes.Buffer) {
	obj, _ := v.(*phptype.Object)
	buffer.WriteRune(TOKEN_OBJECT)
	buffer.WriteString(self.prepareClassName(obj.ClassName))
	encoded := self.encodeArray(obj.Members, false)
	buffer.WriteString(encoded.String())
	return
}

func (self *Serializer) encodeSerialized(v phptype.Value) (buffer bytes.Buffer) {
	var serialized string

	obj, _ := v.(*phptype.ObjectSerialized)
	buffer.WriteRune(TOKEN_OBJECT_SERIALIZED)
	buffer.WriteString(self.prepareClassName(obj.ClassName))

	if self.encodeFunc == nil {
		serialized = obj.Data
	} else {
		var err error
		if serialized, err = self.encodeFunc(obj.Value); err != nil {
			self.saveError(err)
		}
	}

	encoded := self.encodeString(serialized, DELIMITER_OBJECT_LEFT, DELIMITER_OBJECT_RIGHT, false)
	buffer.WriteString(encoded.String())
	return
}

func (self *Serializer) encodeSplArray(v phptype.Value) bytes.Buffer {
	var buffer bytes.Buffer
	obj, _ := v.(*phptype.PhpSplArray)

	buffer.WriteRune(TOKEN_SPL_ARRAY)
	buffer.WriteRune(SEPARATOR_VALUE_TYPE)

	encoded := self.encodeNumber(obj.Flags)
	buffer.WriteString(encoded.String())

	data, _ := self.Encode(obj.Array)
	buffer.WriteString(data)

	buffer.WriteRune(SEPARATOR_VALUES)
	buffer.WriteRune(TOKEN_SPL_ARRAY_MEMBERS)
	buffer.WriteRune(SEPARATOR_VALUE_TYPE)

	data, _ = self.Encode(obj.Properties)
	buffer.WriteString(data)

	return buffer
}

func (self *Serializer) prepareLen(l int) string {
	return string(SEPARATOR_VALUE_TYPE) + strconv.Itoa(l) + string(SEPARATOR_VALUE_TYPE)
}

func (self *Serializer) prepareClassName(name string) string {
	encoded := self.encodeString(name, DELIMITER_STRING_LEFT, DELIMITER_STRING_RIGHT, false)
	return encoded.String()
}

func (self *Serializer) saveError(err error) {
	if self.lastErr == nil {
		self.lastErr = err
	}
}

func wrapWithRune(s string, left, right rune) string {
	return string(left) + s + string(right)
}
