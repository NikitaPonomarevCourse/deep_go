package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Number interface {
	uint16 | uint32 | uint64
}

func ToLittleEndian[T Number](number T) T {
	var littleEndianNumber T = 0
	size := uint64((unsafe.Sizeof(number)))
	var i uint64

	fmt.Println(size)
	for i = 0; i < size; i++ {
		littleEndianNumber |= (((number & (255 << (i * 8))) >> (i * 8)) << (8*(size-1) - (i * 8)))
	}
	return littleEndianNumber
}

func TestConversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion_Uint32(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"zero":                {0x00000000, 0x00000000},
		"all ones":            {0xFFFFFFFF, 0xFFFFFFFF},
		"alternating bytes":   {0x00FF00FF, 0xFF00FF00},
		"half ones":           {0x0000FFFF, 0xFFFF0000},
		"sequential bytes":    {0x01020304, 0x04030201},
		"reversed sequential": {0x04030201, 0x01020304},
		"high byte set":       {0xFF000000, 0x000000FF},
		"low byte set":        {0x000000FF, 0xFF000000},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion_Uint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"zero":                {0x0000000000000000, 0x0000000000000000},
		"all ones":            {0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
		"alternating words":   {0x0000FFFF0000FFFF, 0xFFFF0000FFFF0000},
		"sequential bytes":    {0x0102030405060708, 0x0807060504030201},
		"reversed sequential": {0x0807060504030201, 0x0102030405060708},
		"high qword set":      {0xFF00000000000000, 0x00000000000000FF},
		"low qword set":       {0x00000000000000FF, 0xFF00000000000000},
		"mixed pattern":       {0x1122334455667788, 0x8877665544332211},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion_Uint16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"zero":          {0x0000, 0x0000},
		"all ones":      {0xFFFF, 0xFFFF},
		"sequential":    {0x0102, 0x0201},
		"reversed":      {0x0201, 0x0102},
		"high byte set": {0xFF00, 0x00FF},
		"low byte set":  {0x00FF, 0xFF00},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
