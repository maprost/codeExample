package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var emptySlice []byte
	sliceOfInts := []int{0, 1, 2, 3}
	fmt.Println(unsafe.Sizeof(emptySlice))  // 24byte
	fmt.Println(unsafe.Sizeof(sliceOfInts)) // 24byte

	/*
		type slice struct {
		 array unsafe.Pointer 	// -> 8byte
		 len   int				// -> 8byte
		 cap   int				// -> 8byte
		}
	*/

	sizeOfSliceOfInts := uintptr(len(sliceOfInts)) * reflect.TypeOf(sliceOfInts).Elem().Size()
	fmt.Println(sizeOfSliceOfInts) // 32byte -> 4 * 8byte (every int has 8byte)

	sizeOfEmptySlice := uintptr(len(emptySlice)) * reflect.TypeOf(emptySlice).Elem().Size()
	fmt.Println(sizeOfEmptySlice)

}
