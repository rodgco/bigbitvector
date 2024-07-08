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
	count int
}

// The [New] function creates a new BigBitVector with the specified number of bits.
func New(n int) *BigBitVector {
	alloc := uint(math.Ceil(float64(n) / float64(uintSize)))

	return &BigBitVector{n, make([]uint, alloc), 0}
}

// The [rangeCheck] function checks if the specified index is within the range of the bit vector.
func (b *BigBitVector) rangeCheck(n int) (err error) {
	if n < 0 || n >= b.len {
		err = errors.New("index out of range")
	}
	return
}

// The [Set] function sets the bit at the specified index in the bit vector.
func (b *BigBitVector) Set(n int) (err error) {
	var isSet bool

	isSet, err = b.IsSet(n)

	if isSet || err != nil {
		return 
	}

	b.count++
	b.vec[n/uintSize] |= 1 << (n % uintSize)

	return
}

// The [Unset] function clears the bit at the specified index in the bit vector.
func (b *BigBitVector) Unset(n int) (err error) {
	var isSet bool

	isSet, err = b.IsSet(n)

	if !isSet || err != nil {
		return
	}

	b.count--
	b.vec[n/uintSize] &= ^(1 << (n % uintSize))

	return
}


// The [Toggle] function toggles the bit at the specified index in the bit vector.
func (b *BigBitVector) Toggle(n int) (err error) {
	var isSet bool

	isSet, err = b.IsSet(n)
	if err != nil {
		return
	}

	if isSet {
		b.Unset(n)
	} else {
		b.Set(n)
	}

	return
}

// The [IsSet] function returns true if the bit at the specified index in the
// bit vector is set. Otherwise, it returns false.
func (b *BigBitVector) IsSet(n int) (isSet bool, err error) {
	err = b.rangeCheck(n)
	if err != nil {
		return 
	}

	isSet = (b.vec[n/uintSize] & (1 << (n % uintSize))) != 0

	return
}

// The [Get] function returns the value of the bit at the specified index in the bit vector.
func (b *BigBitVector) Get(n int) (value uint, err error) {
	var isSet bool

	isSet, err = b.IsSet(n)
	if err != nil {
		return
	}

	if isSet {
		value = 1
	} else {
		value = 0
	}

	return
}

// The [Size] function returns the number of bits in the bit vector.
func (b *BigBitVector) Size() (size int) {
	size = b.len

	return
}

// The [UnsetAll] function clears all bits in the bit vector.
func (b *BigBitVector) UnsetAll() {
	for i := range b.vec {
		b.vec[i] = 0
	}
	b.count = 0
}

// The [SetAll] function sets all bits in the bit vector.
func (b *BigBitVector) SetAll() {
	for i := range b.vec {
		b.vec[i] = math.MaxUint
	}
	b.count = b.len
}

// The [Copy] function creates a copy of the bit vector.
func (b *BigBitVector) Copy() *BigBitVector {
	c := New(b.len)
	copy(c.vec, b.vec)
	return c
}

// The [Count] function returns the number of bits set in the bit vector.
func (b *BigBitVector) Count() (count int) {
	count = b.count

	return
}

// The [FindFirstSet] function returns the index of the first bit set in the bit vector.
func (b *BigBitVector) FindFirstSet() (index int, err error) {
	for i := range b.vec {
		if b.vec[i] != 0 {
			for j := 0; j < uintSize; j++ {
				if b.vec[i] & (1 << j) != 0 {
					index = i * uintSize + j
					return
				}
			}
		}
	}

	index = -1
	err = errors.New("no bits set")

	return
}

// The [FindFirstUnset] function returns the index of the first bit unset in the bit vector.
func (b *BigBitVector) FindFirstUnset() (index int, err error) {
	for i := range b.vec {
		if b.vec[i] != math.MaxUint {
			for j := 0; j < uintSize; j++ {
				if b.vec[i] & (1 << j) == 0 {
					index = i * uintSize + j
					return
				}
			}
		}
	}

	index = -1
	err = errors.New("no bits unset")

	return
}

// The [FindNextSet] function returns the index of the next bit set in the bit vector after the specified index.
func (b *BigBitVector) FindNextSet(n int) (index int, err error) {
	err = b.rangeCheck(n)
	if err != nil {
		return
	}

	for i := n/uintSize; i < len(b.vec); i++ {
		if b.vec[i] != 0 {
			for j := n % uintSize; j < uintSize; j++ {
				if b.vec[i] & (1 << j) != 0 {
					index = i * uintSize + j
					return
				}
			}
		}
	}

	index = -1
	err = errors.New("no bits set")

	return
}

// The [FindNextUnset] function returns the index of the next bit unset in the bit vector after the specified index.
func (b *BigBitVector) FindNextUnset(n int) (index int, err error) {
	err = b.rangeCheck(n)
	if err != nil {
		return
	}

	for i := n/uintSize; i < len(b.vec); i++ {
		if b.vec[i] != math.MaxUint {
			for j := n % uintSize; j < uintSize; j++ {
				if b.vec[i] & (1 << j) == 0 {
					index = i * uintSize + j
					return
				}
			}
		}
	}

	index = -1
	err = errors.New("no bits unset")

	return
}

// The [CountRange] function returns the number of bits set in the bit vector from the specified start index to the specified end index.
func (b *BigBitVector) CountRange(start, end int) (count int, err error) {
	err = b.rangeCheck(start)
	if err != nil {
		return
	}

	err = b.rangeCheck(end)
	if err != nil {
		return
	}

	if start > end {
		err = errors.New("start index greater than end index")
	}

	// TODO: Evaluate if bitwise operations are faster than looping through the range
	for i := start; i <= end; i++ {
		isSet, _ := b.IsSet(i)
		if isSet {
			count++
		}
	}

	return
}

