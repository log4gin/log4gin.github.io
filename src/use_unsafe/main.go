package main

import (
	"fmt"
	"unsafe"
	"use_unsafe/module"
)

func main() {
	// unsafe 修改没有被导出的字段
	p := module.NewSturct()
	name := (*string)(unsafe.Pointer(p))
	*name = "新名字"
	// status := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))

	status := (*string)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0))+unsafe.Sizeof(string(""))))

	*status = "已经被unsafe修改"
	fmt.Println(*p)

	fmt.Println(bytesToString(stringTobytes("hi")))

}

func stringTobytes(s string) []byte {

	// rp := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	// rp.Len = 1
	// p := (*[]byte)(unsafe.Pointer(rp))
	// return *p

	b := unsafe.Slice(unsafe.StringData(s), len(s))
	return b
}

func bytesToString(bs []byte) string {
	p := (*string)(unsafe.Pointer(&bs))
	return *p
}
