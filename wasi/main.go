package main

/*
#cgo LDFLAGS: add.o sub.o
*/
import "C"

//go:export add
func add(a, b int32) int32

//go:export sub
func sub(a, b int32) int32

func main() {
	println(add(2, 3))
	println(sub(3, 2))
	println(mul(3, 2))
}
