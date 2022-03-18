package main

import "fmt"

func main() {
	print(1, "2")
	print(int64(2), "3")
}

// 定义约束
type Integer interface {
	~int | ~int32 | ~int64
}

func print[I Integer, V string](i I, v V) {
	fmt.Println(i, v)
}
