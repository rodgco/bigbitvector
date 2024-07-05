// Copyright 2024 Rodrigo Garcia Couto. All rights reserved.

// Package bigbitvector implements a bit vector with a large number of bits.

// The inspiration for this package comes from the Programming Pearls book by
// Jon Bentley. On Column 1 he describes a problem that would be efficiently
// solved by using a bit vector.

// The bit vector is implemented as a slice of uints, where each uint contains
// bits.UintSize bits. The number of bits in the bit vector is specified when
// the bit vector is created. The bit vector supports setting and clearing
// individual bits, testing the value of individual bits, and setting and
// clearing all bits in the bit vector. The bit vector also supports creating
// a copy of the bit vector and determining the number of uints allocated to
// store the bits in the bit vector.

package bigbitvector

import (
	"errors"
	"math"
	"math/bits"
)

var uintSize = bits.UintSize

type BigBitVector struct {
	len int
	vec []uint
}

// The [New] function creates a new BigBitVector with the specified number of bits.
func New(n int) *BigBitVector {
	alloc := uint(math.Ceil(float64(n) / float64(uintSize)))
	
	return &BigBitVector{n, make([]uint, alloc)}
}

// The [SetBit] function sets the bit at the specified index in the bit vector.
func (b *BigBitVector) SetBit(n int) (err error) {
	if n < 0 || n >= b.len {
		return errors.New("index out of range")
	}
	b.vec[n/uintSize] |= 1 << (n % uintSize)
	return nil
}

// The [ClearBit] function clears the bit at the specified index in the bit vector.
func (b *BigBitVector) ClearBit(n int) (err error) {
	if n < 0 || n >= b.len {
		return errors.New("index out of range")
	}
	b.vec[n/uintSize] &= ^(1 << (n % uintSize))
	return nil
}

// The [IsBitSet] function returns true if the bit at the specified index in the
// bit vector is set. Otherwise, it returns false.
func (b *BigBitVector) IsBitSet(n int) (bool, error) {
	if n < 0 || n >= b.len {
		return false, errors.New("index out of range")
	}
	return (b.vec[n/uintSize] & (1 << (n % uintSize))) != 0, nil
}

// The [GetBit] function returns the value of the bit at the specified index in the bit vector.
func (b *BigBitVector) GetBit(n int) uint {
	return b.vec[n/uintSize] & (1 << (n % uintSize))
}

// The [Size] function returns the number of bits in the bit vector.
func (b *BigBitVector) Size() int {
	return b.len
}

// The [ClearAll] function clears all bits in the bit vector.
func (b *BigBitVector) ClearAll() {
	for i := range b.vec {
		b.vec[i] = 0
	}
}

// The [SetAll] function sets all bits in the bit vector.
func (b *BigBitVector) SetAll() {
	for i := range b.vec {
		b.vec[i] = math.MaxUint
	}
}

// The [Copy] function creates a copy of the bit vector.
func (b *BigBitVector) Copy() *BigBitVector {
	c := New(b.len)
	copy(c.vec, b.vec)
	return c
}

