package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	if len(memory) == 0 {
		return
	}
	writeIdx := 0
	adressMapping := make(map[unsafe.Pointer]unsafe.Pointer)
	for idx, v := range memory {
		if v == 0 {
			writeIdx = idx
			break
		}
	}
	if writeIdx == len(memory)-1 {
		return
	}
	for idx := writeIdx + 1; idx < len(memory); idx++ {
		if memory[idx] != 0 {
			memory[writeIdx] = memory[idx]
			adressMapping[unsafe.Pointer(&memory[idx])] = unsafe.Pointer(&memory[writeIdx])
			memory[idx] = 0
			writeIdx++
		}
	}
	for idx := range pointers {
		if newAddr, ok := adressMapping[unsafe.Pointer(pointers[idx])]; ok {
			pointers[idx] = newAddr
		}
	}
}

func TestDefragmentation(t *testing.T) {
	var fragmentedMemory = []byte{
		0xFF, 0x00, 0x00, 0x00,
		0x00, 0xFF, 0x00, 0x00,
		0x00, 0x00, 0xFF, 0x00,
		0x00, 0x00, 0x00, 0xFF,
	}

	var fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&fragmentedMemory[0]),
		unsafe.Pointer(&fragmentedMemory[5]),
		unsafe.Pointer(&fragmentedMemory[10]),
		unsafe.Pointer(&fragmentedMemory[15]),
	}

	var defragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&fragmentedMemory[0]),
		unsafe.Pointer(&fragmentedMemory[1]),
		unsafe.Pointer(&fragmentedMemory[2]),
		unsafe.Pointer(&fragmentedMemory[3]),
	}

	var defragmentedMemory = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	Defragment(fragmentedMemory, fragmentedPointers)
	assert.True(t, reflect.DeepEqual(defragmentedMemory, fragmentedMemory))
	fmt.Println(defragmentedPointers, fragmentedPointers)
	assert.True(t, reflect.DeepEqual(defragmentedPointers, fragmentedPointers))
}
