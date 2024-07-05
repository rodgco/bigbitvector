# Big Bit Vector
**A Big Bit Vector implementation in Go lang**

[![GoDoc](https://godoc.org/github.com/rodgco/bigbitvector?status.svg)](https://godoc.org/github.com/rodgco/bigbitvector)
[![Go Report Card](https://goreportcard.com/badge/github.com/rodgco/bigbitvector)](https://goreportcard.com/report/github.com/rodgco/bigbitvector)
[![Coverage Status](https://coveralls.io/repos/github/rodgco/bigbitvector/badge.svg?branch=main)](https://coveralls.io/github/rodgco/bigbitvector?branch=main)

Package bigbitvector implements a bit vector with a large number of bits.

The inspiration for this package comes from the [Programming Pearls](https://amzn.to/45SEgwA) book by Jon Bentley. On Column 1 he describes a problem that would be efficiently solved by using a bit vector.

The bit vector is implemented as a slice of uints, where each uint contains bits.UintSize bits. The number of bits in the bit vector is specified when the bit vector is created. The bit vector supports setting and clearing individual bits, testing the value of individual bits, and setting and clearing all bits in the bit vector. The bit vector also supports creating a copy of the bit vector and determining the number of uints allocated to store the bits in the bit vector.
