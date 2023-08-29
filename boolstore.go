package main

import (
	"fmt"
	"math"
	"strings"
)

type StoreUnit uint64

const STORE_UNIT_SIZE = 64

// A BoolStore is a memory-efficient way to store a lot of bools
type BoolStore struct {
	store []StoreUnit
	size  uint64
}

// size is the number of bools you want to store
func NewBoolStore(size uint64) BoolStore {
	nm := BoolStore{
		size: size,
	}
	arraySize := uint64(math.Ceil(float64(size) / STORE_UNIT_SIZE))
	store := make([]StoreUnit, arraySize)
	nm.store = store
	return nm
}

// sets the nth bool to true
func (bs *BoolStore) Set(n uint64) {
	if n > bs.size {
		panic("Out of range")
	}
	offset, bit := uintToBitAndOffset(n, STORE_UNIT_SIZE)
	bs.store[offset] |= StoreUnit(bit)
}

// returns the bool value of the nth bool
func (bs *BoolStore) Get(n uint64) bool {
	if n > bs.size {
		panic("Out of range")
	}
	offset, bit := uintToBitAndOffset(n, STORE_UNIT_SIZE)
	return (StoreUnit(bit) & bs.store[offset]) != 0
}

func (bs BoolStore) String() string {
	o := strings.Builder{}
	for i := 0; i < len(bs.store); i++ {
		o.WriteString(fmt.Sprintf("[%08d] ", i))
		for j := 0; j < 64; j += 8 {
			byte := byte(bs.store[i] >> j)
			o.WriteString(fmt.Sprintf("%02x ", byte))
		}
		o.WriteString("\n")
	}
	return o.String()
}

func uintToBitAndOffset(i, size uint64) (uint64, uint64) {
	offset := i / size
	ordinal := i % size
	bit := uint64(1) << ordinal
	return offset, bit
}
