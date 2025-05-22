package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// dfs
func Trace(stacks [][]uintptr) []uintptr {
	visited := make(map[uintptr]struct{})
	result := []uintptr{}
	for _, stackValues := range stacks {
		for _, value := range stackValues {
			trace(value, &visited, &result)
		}
	}
	return result
}

func trace(value uintptr, visited *map[uintptr]struct{}, result *[]uintptr) {
	if value == 0 {
		return
	}
	if _, ok := (*visited)[value]; ok {
		return
	}
	(*visited)[value] = struct{}{}
	*result = append(*result, value)
	trace(*(*uintptr)(unsafe.Pointer(value)), visited, result)
}

func TestTrace(t *testing.T) {
	var heapObjects = []int{
		0x00, 0x00, 0x00, 0x00, 0x00,
	}

	var heapPointer1 *int = &heapObjects[1]
	var heapPointer2 *int = &heapObjects[2]
	var heapPointer3 *int = nil
	var heapPointer4 **int = &heapPointer3

	var stacks = [][]uintptr{
		{
			uintptr(unsafe.Pointer(&heapPointer1)), 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[0])),
			0x00, 0x00, 0x00, 0x00,
		},
		{
			uintptr(unsafe.Pointer(&heapPointer2)), 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[1])),
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[2])),
			uintptr(unsafe.Pointer(&heapPointer4)), 0x00, 0x00, 0x00,
		},
		{
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[3])),
		},
	}

	pointers := Trace(stacks)
	expectedPointers := []uintptr{
		uintptr(unsafe.Pointer(&heapPointer1)),
		uintptr(unsafe.Pointer(&heapObjects[0])),
		uintptr(unsafe.Pointer(&heapPointer2)),
		uintptr(unsafe.Pointer(&heapObjects[1])),
		uintptr(unsafe.Pointer(&heapObjects[2])),
		uintptr(unsafe.Pointer(&heapPointer4)),
		uintptr(unsafe.Pointer(&heapPointer3)),
		uintptr(unsafe.Pointer(&heapObjects[3])),
	}
	expectedSet := make(map[uintptr]struct{})
	for _, ptr := range expectedPointers {
		expectedSet[ptr] = struct{}{}
	}
	pointersSet := make(map[uintptr]struct{})
	for _, ptr := range pointers {
		pointersSet[ptr] = struct{}{}
	}

	assert.Equal(t, len(expectedSet), len(pointersSet), "Arrays have different lengths")
	for ptr := range expectedSet {
		assert.Contains(t, pointersSet, ptr, "Expected pointer %v not found in result", ptr)
	}
}
