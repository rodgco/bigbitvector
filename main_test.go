package bigbitvector

import (
	"math/bits"
	"testing"
)

func TestBitVector(t *testing.T) {
	var size int = 10

	bv := New(size)

	if bv.Size() != size {
		t.Error("Expected vector size to be the size assigned")
	}

	err := bv.Set(0)
	if err != nil {
		t.Error("Expected no error")
	}

	err = bv.Set(size)
	if err == nil {
		t.Error("Expected \"out of range\" error")
	}

	err = bv.Unset(-1)
	if err == nil {
		t.Error("Expected \"out of range\" error")
	}

	_, err = bv.Get(size)
	if err == nil {
		t.Error("Expected \"out of range\" error")
	}

	_, err = bv.IsSet(-1)
	if err == nil {
		t.Error("Expected \"out of range\" error")
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
