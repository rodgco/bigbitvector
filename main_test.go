package bigbitvector

import (
	"math/bits"
	"testing"
)

func TestBitVector(t *testing.T) {
	var size int = 10


	// Test bigbitvector initialization
	bv := New(size)
	if bv.Size() != size {
		t.Error("Expected vector size to be the size assigned")
	}

	if bv.Count() != 0 {
		t.Error("Expected count to be 0")
	}

	// Test bigbitvector range checking
	_, err := bv.IsSet(0)
	if err != nil {
		t.Error("Expected no error, value within range")
	}

	_, err = bv.IsSet(size-1)
	if err != nil {
		t.Error("Expected no error, value within range")
	}

	_, err = bv.IsSet(-1)
	if err == nil {
		t.Error("Expected \"out of range\" error (lower bound)")
	}

	_, err = bv.IsSet(size)
	if err == nil {
		t.Error("Expected \"out of range\" error (upper bound)")
	}

	_, err = bv.Get(0)
	if err != nil {
		t.Error("Expected no error, value within range")
	}

	_, err = bv.Get(size-1)
	if err != nil {
		t.Error("Expected no error, value within range")
	}

	_, err = bv.Get(-1)
	if err == nil {
		t.Error("Expected \"out of range\" error (lower bound)")
	}

	_, err = bv.Get(size)
	if err == nil {
		t.Error("Expected \"out of range\" error (upper bound)")
	}

	// Test bigbitvector set/unset
	_ = bv.Set(0)
	isSet, _ := bv.IsSet(0)
	if !isSet {
		t.Error("Expected bit 0 to be set")
	}

	value, _ := bv.Get(0)
	if value != 1 {
		t.Error("Expected value to be 1")
	}

	if bv.Count() != 1 {
		t.Error("Expected count to be 1 after set")
	}

	_ = bv.Set(0)
	if bv.Count() != 1 {
		t.Error("Expected count to be 1 after setting a setted bit")
	}

	_ = bv.Unset(0)

	isSet, _ = bv.IsSet(0)
	if isSet {
		t.Error("Expected bit 0 to be unset")
	}

	value, _ = bv.Get(0)
	if value != 0 {
		t.Error("Expected value to be 0")
	}

	if bv.Count() != 0 {
		t.Error("Expected count to be 0 after unset")
	}

	_ = bv.Unset(0)
	if bv.Count() != 0 {
		t.Error("Expected count to be 0 after unsetting an unset bit")
	}
}

func TestBitVectorToggle(t *testing.T) {
	var size int = 10
	bv := New(size)

	// Test bigbitvector toggle range checking
	err := bv.Toggle(-1)
	if err == nil {
		t.Error("Expected \"out of range\" error (lower bound)")
	}

	err = bv.Toggle(size)
	if err == nil {
		t.Error("Expected \"out of range\" error (upper bound)")
	}

	// Test bigbitvector toggle
	_ = bv.Toggle(0)
	isSet, _ := bv.IsSet(0)
	if !isSet {
		t.Error("Expected bit 0 to be set after toggle")
	}

	value, _ := bv.Get(0)
	if value != 1 {
		t.Error("Expected value to be 1 after toggle")
	}

	if bv.Count() != 1 {
		t.Error("Expected count to be 1 after toggle")
	}

	_ = bv.Toggle(0)
	isSet, _ = bv.IsSet(0)
	if isSet {
		t.Error("Expected bit 0 to be unset after toggle")
	}

	value, _ = bv.Get(0)
	if value != 0 {
		t.Error("Expected value to be 0 after toggle")
	}

	if bv.Count() != 0 {
		t.Error("Expected count to be 0 after toggle")
	}
}

func TestBitVectorAllocation(t *testing.T) {
	bv1 := New(bits.UintSize)
	if len(bv1.vec) != 1 {
		t.Error("Expected 1 unit of uint allocation")
	}

	bv2 := New(bits.UintSize + 1)
	if len(bv2.vec) != 2 {
		t.Error("Expected 2 units of uint allocation")
	}
}

func TestSetBit(t *testing.T) {
	bv := New(10)

	test, _ := bv.IsSet(0)

	if test {
		t.Error("Expected bit 0 to be unset on init")
	}

	bv.Set(0)

	test, _ = bv.IsSet(0)

	if !test {
		t.Error("Expected true, got false")
	}
}

func TestClearBit(t *testing.T) {
	bv := New(10)

	bv.Set(0)

	test, _ := bv.IsSet(0)

	if !test {
		t.Error("Expected true, got false")
	}

	bv.Unset(0)

	test, _ = bv.IsSet(0)

	if test {
		t.Error("Expected false, got true")
	}
}

func TestSetAllBits(t *testing.T) {
	bv := New(10)

	bv.SetAll()

	for i := 0; i < 10; i++ {
		test, _ := bv.IsSet(i)
		if !test {
			t.Error("Expected true, got false")
		}
	}
}

func TestClearAll(t *testing.T) {
	bv := New(10)

	bv.SetAll()

	for i := 0; i < 10; i++ {
		test, _ := bv.IsSet(i)
		if !test {
			t.Error("Expected true, got false")
		}
	}

	bv.UnsetAll()

	for i := 0; i < 10; i++ {
		test, _ := bv.IsSet(i)
		if test {
			t.Error("Expected false, got true")
		}
	}
}

func TestCopy(t *testing.T) {
	bv1 := New(10)

	bv1.SetAll()

	bv2 := bv1.Copy()

	for i := 0; i < 10; i++ {
		test1, _ := bv1.Get(i)
		test2, _ := bv2.Get(i)
		if test1 != test2 {
			t.Error("Expected equality")
		}
	}
}

func TestFindFirstSet(t *testing.T) {
	bv := New(10)

	bv.Set(3)

	index, err := bv.FindFirstSet()
	if err != nil {
		t.Error("Expected no error")
	}
	if index != 3 {
		t.Error("Expected 3, got ", index)
	}

	bv.Unset(3)

	index, err = bv.FindFirstSet()
	if index != -1 {
		t.Error("Expected -1, got ", index)
	}
	if err == nil {
		t.Error("Expected error")
	}
}

func TestFindFirstUnset(t *testing.T) {
	bv := New(10)

	bv.SetAll()
	bv.Unset(3)

	index, err := bv.FindFirstUnset()
	if index != 3 {
		t.Error("Expected 3, got ", index)
	}
	if err != nil {
		t.Error("Expected no error")
	}

	bv.Set(3)

	index, err = bv.FindFirstUnset()
	if index != -1 {
		t.Error("Expected -1, got ", index)
	}
	if err == nil {
		t.Error("Expected error")
	}
}

func TestFindNextSet(t *testing.T) {
	bv := New(10)

	bv.Set(3)

	index, err := bv.FindNextSet(0)
	if err != nil {
		t.Error("Expected no error")
	}
	if index != 3 {
		t.Error("Expected 3, got ", index)
	}

	index, err = bv.FindNextSet(4)
	if index != -1 {
		t.Error("Expected -1, got ", index)
	}
	if err == nil {
		t.Error("Expected error")
	}

	_, err = bv.FindNextSet(-1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestFindNextUnset(t *testing.T) {
	bv := New(10)

	bv.SetAll()
	bv.Unset(3)

	index, err := bv.FindNextUnset(0)
	if err != nil {
		t.Error("Expected no error")
	}
	if index != 3 {
		t.Error("Expected 3, got ", index)
	}

	index, err = bv.FindNextUnset(4)
	if index != -1 {
		t.Error("Expected -1, got ", index)
	}
	if err == nil {
		t.Error("Expected error")
	}

	_, err = bv.FindNextUnset(-1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestFindNthSet(t *testing.T) {
	bv := New(10)

	bv.Set(3)
	bv.Set(5)
	bv.Set(7)

	index, err := bv.FindNthSet(1)
	if err != nil {
		t.Error("Expected no error")
	}
	if index != 3 {
		t.Error("Expected 3, got ", index)
	}

	index, err = bv.FindNthSet(3)
	if err != nil {
		t.Error("Expected no error")
	}
	if index != 7 {
		t.Error("Expected 7, got ", index)
	}

	_, err = bv.FindNthSet(4)
	if err == nil {
		t.Error("Expected error")
	}

	_, err = bv.FindNthSet(-1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestCountRange(t *testing.T) {
	bv := New(10)

	bv.Set(0)
	bv.Set(1)
	bv.Set(2)

	_, err := bv.CountRange(2, 0)
	if err == nil {
		t.Error("Expected error")
	}

	count, err := bv.CountRange(0, 2)
	if err != nil {
		t.Error("Expected no error")
	}
	if count != 3 {
		t.Error("Expected 3, got ", count)
	}

	count, err = bv.CountRange(0, 1)
	if count != 2 {
		t.Error("Expected 2, got ", count)
	}

	count, err = bv.CountRange(1, 2)
	if count != 2 {
		t.Error("Expected 2, got ", count)
	}

	count, err = bv.CountRange(1, 1)
	if count != 1 {
		t.Error("Expected 1, got ", count)
	}

	_, err = bv.CountRange(-1, 2)
	if err == nil {
		t.Error("Expected error")
	}

	_, err = bv.CountRange(0, 11)
	if err == nil {
		t.Error("Expected error")
	}
}

