package main

import (
	"log"
	"unsafe"
)

func accessSlice() {
	// init slice
	a1 := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		a1[i] = 100 + i
	}

	// access
	for i := 0; i < 10; i++ {
		v := *(*int)(unsafe.Pointer(&a1[i]))

		println(v)
	}
}

func accessStruct() {
	type Friends struct {
		Name string
		Age  int
	}

	name := "zy"
	age := 18
	f1 := Friends{Name: name, Age: age}

	log.Printf("sizeof(Friends): %d, sizeof(string)=%d, sizeof(int)=%d", unsafe.Sizeof(f1), unsafe.Sizeof(name), unsafe.Sizeof(age))

	log.Printf("before update f1: %+v", f1)

	// 指针指向f1对象
	f1Ptr := unsafe.Pointer(&f1)

	// 方法1： f1对象的首地址 + name的大小，等于Age元素
	agePtr := unsafe.Pointer((uintptr(f1Ptr) + unsafe.Sizeof(name)))
	*(*int)(agePtr) = 20
	log.Printf("after update f1: %+v", f1)

	// 方法2：使用unsafe.Offsetof 访问Age元素
	agePtr = unsafe.Pointer((uintptr(f1Ptr) + unsafe.Offsetof(f1.Age)))
	*(*int)(agePtr) = 30
	log.Printf("after update f1: %+v", f1)

}

func main() {
	accessSlice()

	accessStruct()
}
