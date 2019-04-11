package phpserialize

import (
	"encoding/json"
	"testing"

	"github.com/imantung/phpsessgo/phptype"
)

func TestDecodeNil(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("N;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding nil value: %v\n", err)
	} else {
		if val != nil {
			t.Errorf("Nil value decoded incorrectly, have got %v\n", val)
		}
	}
}

func TestDecodeBoolTrue(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("b:1;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding bool (true) value: %v\n", err)
	} else {
		if boolVal, ok := val.(bool); !ok {
			t.Errorf("Unable to convert %v to bool\n", val)
		} else if boolVal != true {
			t.Errorf("Bool (true) value decoded incorrectly, expected: %v, have got: %v\n", true, boolVal)
		}
	}
}

func TestDecodeBoolFalse(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("b:0;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding bool (false) value: %v\n", err)
	} else {
		if boolVal, ok := val.(bool); !ok {
			t.Errorf("Unable to convert %v to bool\n", val)
		} else if boolVal != false {
			t.Errorf("Bool (false) value decoded incorrectly, expected: %v, have got: %v\n", false, boolVal)
		}
	}
}

func TestDecodeInt(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("i:42;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding int value: %v\n", err)
	} else {
		if intVal, ok := val.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", val)
		} else if intVal != 42 {
			t.Errorf("Int value decoded incorrectly, expected: %v, have got: %v\n", 42, intVal)
		}
	}
}

func TestDecodeIntMinus(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("i:-42;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding int value: %v\n", err)
	} else {
		if intVal, ok := val.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", val)
		} else if intVal != -42 {
			t.Errorf("Int value decoded incorrectly, expected: %v, have got: %v\n", -42, intVal)
		}
	}
}

func TestDecodeFloat64(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("d:42.378900000000002;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding float4 value: %v\n", err)
	} else {
		if floatVal, ok := val.(float64); !ok {
			t.Errorf("Unable to convert %v to float\n", val)
		} else if floatVal != 42.378900000000002 {
			t.Errorf("Float64 value decoded incorrectly, expected: %v, have got: %v\n", 42.378900000000002, floatVal)
		}
	}
}

func TestDecodeFloat64Minus(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("d:-42.378900000000002;")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding float4 value: %v\n", err)
	} else {
		if floatVal, ok := val.(float64); !ok {
			t.Errorf("Unable to convert %v to float\n", val)
		} else if floatVal != -42.378900000000002 {
			t.Errorf("Float64 value decoded incorrectly, expected: %v, have got: %v\n", -42.378900000000002, floatVal)
		}
	}
}

func TestDecodeString(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("s:6:\"foobar\";")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding string value: %v\n", err)
	} else {
		if strVal, ok := val.(string); !ok {
			t.Errorf("Unable to convert %v to string\n", val)
		} else if strVal != "foobar" {
			t.Errorf("String value decoded incorrectly, expected: %v, have got: %v\n", "foobar", strVal)
		}
	}
}

func TestDecodeArray(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("a:3:{i:0;i:10;i:1;i:11;i:2;i:12;}")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding array value: %v\n", err)
	} else {
		if arrVal, ok := val.(phptype.Array); !ok {
			t.Errorf("Unable to convert %v to phptype.Array\n", val)
		} else if v1, ok1 := arrVal[phptype.Value(0)]; !ok1 {
			t.Errorf("Array value decoded incorrectly, key `0` doest not exists\n")
		} else if intV1, ok1 := v1.(int); !ok1 {
			t.Errorf("Unable to convert %v to int\n", v1)
		} else if intV1 != 10 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 10, intV1)
		} else if v2, ok2 := arrVal[phptype.Value(1)]; !ok2 {
			t.Errorf("Array value decoded incorrectly, key `1` doest not exists\n")
		} else if intV2, ok2 := v2.(int); !ok2 {
			t.Errorf("Unable to convert %v to int\n", v2)
		} else if intV2 != 11 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 11, intV2)
		} else if v3, ok3 := arrVal[phptype.Value(2)]; !ok3 {
			t.Errorf("Array value decoded incorrectly, key `2` doest not exists\n")
		} else if intV3, ok3 := v3.(int); !ok3 {
			t.Errorf("Unable to convert %v to int\n", v3)
		} else if intV3 != 12 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 12, intV3)
		}
	}
}

func TestDecodeArrayMap(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("a:2:{s:3:\"foo\";i:4;s:3:\"bar\";i:2;}")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding array value: %v\n", err)
	} else {
		if arrVal, ok := val.(phptype.Array); !ok {
			t.Errorf("Unable to convert %v to phptype.Array\n", val)
		} else if v1, ok1 := arrVal["foo"]; !ok1 {
			t.Errorf("Array value decoded incorrectly, key `foo` doest not exists\n")
		} else if intV1, ok1 := v1.(int); !ok1 {
			t.Errorf("Unable to convert %v to int\n", v1)
		} else if intV1 != 4 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 4, intV1)
		} else if v2, ok2 := arrVal["bar"]; !ok2 {
			t.Errorf("Array value decoded incorrectly, key `bar` doest not exists\n")
		} else if intV2, ok2 := v2.(int); !ok2 {
			t.Errorf("Unable to convert %v to int\n", v2)
		} else if intV2 != 2 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 2, intV2)
		}
	}
}

func TestDecodeArrayArray(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("a:2:{s:3:\"foo\";a:3:{i:0;i:10;i:1;i:11;i:2;i:12;}s:3:\"bar\";i:2;}")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding array value: %v\n", err)
	} else {
		if arrVal, ok := val.(phptype.Array); !ok {
			t.Errorf("Unable to convert %v to phptype.Array\n", val)
		} else if v1, ok1 := arrVal["foo"]; !ok1 {
			t.Errorf("Array value decoded incorrectly, key `foo` doest not exists\n")
		} else if innerArr, ok1 := v1.(phptype.Array); !ok1 {
			t.Errorf("Unable to convert %v to inner phptype.Array\n", v1)
		} else if inv1, inOk1 := innerArr[phptype.Value(0)]; !inOk1 {
			t.Errorf("Array value decoded incorrectly, key `0` doest not exists\n")
		} else if inIntV1, inOk1 := inv1.(int); !inOk1 {
			t.Errorf("Unable to convert %v to int\n", inv1)
		} else if inIntV1 != 10 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 10, inIntV1)
		} else if inv2, inOk2 := innerArr[phptype.Value(1)]; !inOk2 {
			t.Errorf("Array value decoded incorrectly, key `1` doest not exists\n")
		} else if inIntV2, inOk2 := inv2.(int); !inOk2 {
			t.Errorf("Unable to convert %v to int\n", inv2)
		} else if inIntV2 != 11 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 11, inIntV2)
		} else if inv3, inOk3 := innerArr[phptype.Value(2)]; !inOk3 {
			t.Errorf("Array value decoded incorrectly, key `2` doest not exists\n")
		} else if inIntV3, inOk3 := inv3.(int); !inOk3 {
			t.Errorf("Unable to convert %v to int\n", inv3)
		} else if inIntV3 != 12 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 12, inIntV3)
		} else if v2, ok2 := arrVal["bar"]; !ok2 {
			t.Errorf("Array value decoded incorrectly, key `bar` doest not exists\n")
		} else if intV2, ok2 := v2.(int); !ok2 {
			t.Errorf("Unable to convert %v to int\n", v2)
		} else if intV2 != 2 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 2, intV2)
		}
	}
}

func TestDecodeObject(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("O:4:\"Test\":3:{s:6:\"public\";i:1;s:12:\"\x00*\x00protected\";i:2;s:13:\"\x00Test\x00private\";i:3;}")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding object value: %v\n", err)
	} else {
		if obj, ok := val.(*phptype.Object); !ok {
			t.Errorf("Unable to convert %v to *phptype.Object\n", val)
		} else if obj.ClassName != "Test" {
			t.Errorf("Object class name decoded incorrectly, expected: %s, have got: %s\n", "Test", obj.ClassName)
		} else if value1, ok := obj.GetPublic("public"); !ok {
			t.Errorf("Public member of object was decoded incorrectly, key `publice` doest not exists\n")
		} else if intV1, ok := value1.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value1)
		} else if intV1 != 1 {
			t.Errorf("Public member of object was decoded incorrectly, expected: %v, have got: %v\n", 1, intV1)
		} else if value2, ok := obj.GetProtected("protected"); !ok {
			t.Errorf("Protected member of object was decoded incorrectly, key `protected` doest not exists\n")
		} else if intV2, ok := value2.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value2)
		} else if intV2 != 2 {
			t.Errorf("Protected member of object was decoded incorrectly, expected: %v, have got: %v\n", 2, intV2)
		} else if value3, ok := obj.GetPrivate("private"); !ok {
			t.Errorf("Private member of object was decoded incorrectly, key `private` doest not exists\n")
		} else if intV3, ok := value3.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value3)
		} else if intV3 != 3 {
			t.Errorf("Private member of object was decoded incorrectly, expected: %v, have got: %v\n", 3, intV3)
		}
	}
}

func TestDecodeArrayOfObjects(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("a:2:{i:0;O:5:\"Test1\":3:{s:6:\"public\";i:11;s:12:\"\x00*\x00protected\";i:12;s:14:\"\x00Test1\x00private\";i:13;}i:1;O:5:\"Test2\":3:{s:6:\"public\";i:21;s:12:\"\x00*\x00protected\";i:22;s:14:\"\x00Test2\x00private\";i:23;}}")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding array of objects value: %v\n", err)
	} else {
		if arrVal, ok := val.(phptype.Array); !ok {
			t.Errorf("Unable to convert %v to phptype.Array\n", val)
		} else if v1, ok1 := arrVal[phptype.Value(0)]; !ok1 {
			t.Errorf("Array value decoded incorrectly, key `0` doest not exists\n")
		} else if obj1, ok1 := v1.(*phptype.Object); !ok1 {
			t.Errorf("Unable to convert %v to *phptype.Object\n", v1)
		} else if obj1.ClassName != "Test1" {
			t.Errorf("Object class name decoded incorrectly, expected: %s, have got: %s\n", "Test1", obj1.ClassName)
		} else if value1, ok := obj1.GetPublic("public"); !ok {
			t.Errorf("Public member of object was decoded incorrectly, key `publice` doest not exists\n")
		} else if intV1, ok := value1.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value1)
		} else if intV1 != 11 {
			t.Errorf("Public member of object was decoded incorrectly, expected: %v, have got: %v\n", 11, intV1)
		} else if value2, ok := obj1.GetProtected("protected"); !ok {
			t.Errorf("Protected member of object was decoded incorrectly, key `protected` doest not exists\n")
		} else if intV2, ok := value2.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value2)
		} else if intV2 != 12 {
			t.Errorf("Protected member of object was decoded incorrectly, expected: %v, have got: %v\n", 12, intV2)
		} else if value3, ok := obj1.GetPrivate("private"); !ok {
			t.Errorf("Private member of object was decoded incorrectly, key `private` doest not exists\n")
		} else if intV3, ok := value3.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value3)
		} else if intV3 != 13 {
			t.Errorf("Private member of object was decoded incorrectly, expected: %v, have got: %v\n", 13, intV3)
		} else if v2, ok2 := arrVal[phptype.Value(1)]; !ok2 {
			t.Errorf("Array value decoded incorrectly, key `1` doest not exists\n")
		} else if obj2, ok2 := v2.(*phptype.Object); !ok2 {
			t.Errorf("Unable to convert %v to *phptype.Object\n", v2)
		} else if obj2.ClassName != "Test2" {
			t.Errorf("Object class name decoded incorrectly, expected: %s, have got: %s\n", "Test2", obj2.ClassName)
		} else if value1, ok := obj2.GetPublic("public"); !ok {
			t.Errorf("Public member of object was decoded incorrectly, key `publice` doest not exists\n")
		} else if intV1, ok := value1.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value1)
		} else if intV1 != 21 {
			t.Errorf("Public member of object was decoded incorrectly, expected: %v, have got: %v\n", 21, intV1)
		} else if value2, ok := obj2.GetProtected("protected"); !ok {
			t.Errorf("Protected member of object was decoded incorrectly, key `protected` doest not exists\n")
		} else if intV2, ok := value2.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value2)
		} else if intV2 != 22 {
			t.Errorf("Protected member of object was decoded incorrectly, expected: %v, have got: %v\n", 22, intV2)
		} else if value3, ok := obj2.GetPrivate("private"); !ok {
			t.Errorf("Private member of object was decoded incorrectly, key `private` doest not exists\n")
		} else if intV3, ok := value3.(int); !ok {
			t.Errorf("Unable to convert %v to int\n", value3)
		} else if intV3 != 23 {
			t.Errorf("Private member of object was decoded incorrectly, expected: %v, have got: %v\n", 23, intV3)
		}
	}
}

func TestDecodeObjectSerializable(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("C:16:\"TestSerializable\":6:{foobar}")
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding object value: %v\n", err)
	} else {
		if obj, ok := val.(*phptype.ObjectSerialized); !ok {
			t.Errorf("Unable to convert %v to *phptype.ObjectSerialized\n", val)
		} else if obj.ClassName != "TestSerializable" {
			t.Errorf("Object class name decoded incorrectly, expected: %s, have got: %s\n", "TestSerializable", obj.ClassName)
		} else if obj.Data != "foobar" {
			t.Errorf("Object value decoded incorrectly, expected: %v, have got: %v\n", "foobar", obj.Data)
		}
	}
}

func TestDecodeObjectSerializableArray(t *testing.T) {
	var (
		val phptype.Value
		err error
	)

	decoder := NewUnSerializer("C:17:\"TestSerializable1\":34:{a:2:{s:3:\"foo\";i:4;s:3:\"bar\";i:2;}}")
	decoder.SetSerializedDecodeFunc(SerializedDecodeFunc(UnSerialize))
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding object value: %v\n", err)
	} else {
		if obj, ok := val.(*phptype.ObjectSerialized); !ok {
			t.Errorf("Unable to convert %v to *phptype.ObjectSerialized\n", val)
		} else if obj.ClassName != "TestSerializable1" {
			t.Errorf("Object class name decoded incorrectly, expected: %s, have got: %s\n", "TestSerializable1", obj.ClassName)
		} else if obj.Data != "a:2:{s:3:\"foo\";i:4;s:3:\"bar\";i:2;}" {
			t.Errorf("Object value decoded incorrectly, expected: %v, have got: %v\n", "a:2:{s:3:\"foo\";i:4;s:3:\"bar\";i:2;}", obj.Data)
		} else if vv := obj.Value; vv == nil {
			t.Errorf("Object value decoded incorrectly, expected value as phptype.Array, have got: %v\n", obj.Value)
		} else if arrVal, ok := vv.(phptype.Array); !ok {
			t.Errorf("Unable to convert %v to phptype.Array\n", vv)
		} else if v1, ok1 := arrVal["foo"]; !ok1 {
			t.Errorf("Array value decoded incorrectly, key `foo` doest not exists\n")
		} else if intV1, ok1 := v1.(int); !ok1 {
			t.Errorf("Unable to convert %v to int\n", v1)
		} else if intV1 != 4 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 4, intV1)
		} else if v2, ok2 := arrVal["bar"]; !ok2 {
			t.Errorf("Array value decoded incorrectly, key `bar` doest not exists\n")
		} else if intV2, ok2 := v2.(int); !ok2 {
			t.Errorf("Unable to convert %v to int\n", v2)
		} else if intV2 != 2 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 2, intV2)
		}
	}
}

func TestDecodeObjectSerializableJSON(t *testing.T) {
	var (
		val phptype.Value
		err error
		f   SerializedDecodeFunc
	)

	f = func(s string) (phptype.Value, error) {
		var (
			val map[string]int
			err error
		)
		err = json.Unmarshal([]byte(s), &val)
		return val, err
	}

	decoder := NewUnSerializer("C:17:\"TestSerializable2\":17:{{\"foo\":4,\"bar\":2}}")
	decoder.SetSerializedDecodeFunc(f)
	if val, err = decoder.Decode(); err != nil {
		t.Errorf("Error while decoding object value: %v\n", err)
	} else {
		if obj, ok := val.(*phptype.ObjectSerialized); !ok {
			t.Errorf("Unable to convert %v to *phptype.ObjectSerialized\n", val)
		} else if obj.ClassName != "TestSerializable2" {
			t.Errorf("Object class name decoded incorrectly, expected: %s, have got: %s\n", "TestSerializable2", obj.ClassName)
		} else if obj.Data != "{\"foo\":4,\"bar\":2}" {
			t.Errorf("Object value decoded incorrectly, expected: %v, have got: %v\n", "{\"foo\":4,\"bar\":2}", obj.Data)
		} else if vv := obj.Value; vv == nil {
			t.Errorf("Object value decoded incorrectly, expected value as phptype.Array, have got: %v\n", obj.Value)
		} else if arrVal, ok := vv.(map[string]int); !ok {
			t.Errorf("Unable to convert %v to phptype.Array\n", vv)
		} else if v1, ok1 := arrVal["foo"]; !ok1 {
			t.Errorf("Array value decoded incorrectly, key `foo` doest not exists\n")
		} else if v1 != 4 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 4, v1)
		} else if v2, ok2 := arrVal["bar"]; !ok2 {
			t.Errorf("Array value decoded incorrectly, key `bar` doest not exists\n")
		} else if v2 != 2 {
			t.Errorf("Array value decoded incorrectly, expected: %v, have got: %v\n", 2, v2)
		}
	}
}

func TestDecodeSplArray(t *testing.T) {
	val, err := UnSerialize("x:i:0;a:1:{s:3:\"foo\";s:3:\"bar\";};m:a:0:{}")
	if err != nil {
		t.Errorf("Can't decode array object: %v\n", err)
	}

	obj, ok := val.(*phptype.PhpSplArray)
	if !ok {
		t.Errorf("Unable to convert %v to *phptype.PhpSplArray", val)
	}

	array, ok := obj.Array.(phptype.Array)
	if !ok {
		t.Errorf("Can't convert %v to phptype.Array", obj.Array)
	}

	if len(array) != 1 || array["foo"] != "bar" {
		t.Errorf("Can't find 'foo' key in %v", array)
	}

	properties, ok := obj.Properties.(phptype.Array)
	if !ok {
		t.Errorf("Can't convert %v to phptype.Array", obj.Properties)
	}

	if len(properties) > 0 {
		t.Errorf("Expected empty phptype.Array, got %v", properties)
	}
}

func TestDecodeSplArraySerialized(t *testing.T) {
	objValue, err := UnSerialize("C:11:\"ArrayObject\":21:{x:i:0;a:0:{};m:a:0:{}}")
	if err != nil {
		t.Errorf("Error while decoding object value: %v\n", err)
	}

	obj, ok := objValue.(*phptype.ObjectSerialized)
	if !ok {
		t.Errorf("Unable to convert %v to *phptype.ObjectSerialized\n", objValue)
	}

	array, ok := obj.Value.(*phptype.PhpSplArray)
	if !ok {
		t.Errorf("Unable to convert %v to *phptype.PhpSplArray\n", obj.Value)
	}

	if array.Flags != 0 {
		t.Errorf("SplArray flags expected: 0, got %v\n", array.Flags)
	}

	arrayStorage, ok := array.Array.(phptype.Array)
	if !ok || arrayStorage == nil {
		t.Errorf("SplArray.Array expected: empty phptype.Array, got %v", array.Array)
	}

	arrayProperties, ok := array.Properties.(phptype.Array)
	if !ok || arrayProperties == nil {
		t.Errorf("SplArray.Properties expected: empty phptype.Array, got %v", array.Properties)
	}
}
