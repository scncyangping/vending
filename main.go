package main

func main() {
	println("world")
}

// 定义约束
type Integer interface {
	~int | ~int32 | ~int64
}
