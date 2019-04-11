package phptype

import (
	"testing"
)

func TestValueString(t *testing.T) {
	var (
		val      Value = "string"
		expected string   = "string"
	)
	if newVal := ValueString(val); newVal != expected {
		t.Errorf("Expected %q but got %q", expected, newVal)
	}
}

func TestValueBool(t *testing.T) {
	var (
		val      Value = true
		expected bool     = true
	)
	if newVal := ValueBool(val); newVal != expected {
		t.Errorf("Expected %t but got %t", expected, newVal)
	}
}

func TestValueInt(t *testing.T) {
	var (
		val      Value = 10
		expected int      = 10
	)
	if newVal := ValueInt(val); newVal != expected {
		t.Errorf("Expected %d but got %d", expected, newVal)
	}
}

func TestValueInt64(t *testing.T) {
	var (
		val      Value = int64(10)
		expected int64    = 10
	)
	if newVal := ValueInt64(val); newVal != expected {
		t.Errorf("Expected %d but got %d", expected, newVal)
	}
}

func TestValueUInt(t *testing.T) {
	var (
		val      Value = uint(10)
		expected uint     = 10
	)
	if newVal := ValueUInt(val); newVal != expected {
		t.Errorf("Expected %d but got %d", expected, newVal)
	}
}

func TestValueUInt64(t *testing.T) {
	var (
		val      Value = uint64(10)
		expected uint64   = 10
	)
	if newVal := ValueUInt64(val); newVal != expected {
		t.Errorf("Expected %d but got %d", expected, newVal)
	}
}

func TestValueFloat64(t *testing.T) {
	var (
		val      Value = float64(10.0)
		expected float64  = 10.0
	)
	if newVal := ValueFloat64(val); newVal != expected {
		t.Errorf("Expected %v but got %v", expected, newVal)
	}
}
