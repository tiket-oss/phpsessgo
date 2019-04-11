package phptype

import (
	"strconv"
)

func ValueString(p Value) (res string) {
	res, _ = p.(string)
	return
}

func ValueBool(p Value) (res bool) {
	switch p.(type) {
	case bool:
		res, _ = p.(bool)
	case string:
		str, _ := p.(string)
		res, _ = strconv.ParseBool(str)
	}
	return
}

func ValueInt(p Value) (res int) {
	switch p.(type) {
	case int:
		res, _ = p.(int)
	case int8:
		intVal, _ := p.(int8)
		res = int(intVal)
	case int16:
		intVal, _ := p.(int16)
		res = int(intVal)
	case int32:
		intVal, _ := p.(int32)
		res = int(intVal)
	case int64:
		intVal, _ := p.(int64)
		res = int(intVal)
	case uint:
		intVal, _ := p.(uint)
		res = int(intVal)
	case uint8:
		intVal, _ := p.(uint8)
		res = int(intVal)
	case uint16:
		intVal, _ := p.(uint16)
		res = int(intVal)
	case uint32:
		intVal, _ := p.(uint32)
		res = int(intVal)
	case uint64:
		intVal, _ := p.(uint64)
		res = int(intVal)
	case string:
		str, _ := p.(string)
		res, _ = strconv.Atoi(str)
	}
	return
}

func ValueInt64(p Value) (res int64) {
	switch p.(type) {
	case int64:
		res = p.(int64)
	default:
		res = int64(ValueInt(p))
	}
	return
}

func ValueUInt(p Value) (res uint) {
	switch p.(type) {
	case uint:
		res = p.(uint)
	default:
		res = uint(ValueInt(p))
	}
	return
}

func ValueUInt64(p Value) (res uint64) {
	switch p.(type) {
	case uint64:
		res = p.(uint64)
	default:
		res = uint64(ValueInt(p))
	}
	return
}

func ValueFloat64(p Value) (res float64) {
	switch p.(type) {
	case float64:
		res, _ = p.(float64)
	case string:
		str, _ := p.(string)
		res, _ = strconv.ParseFloat(str, 64)
	default:
		return float64(ValueInt(p))
	}
	return
}
