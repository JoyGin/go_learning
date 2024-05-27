package series

import "fmt"

func GetFibonacciSeries(n int) []int {
	fiboList := []int{1, 1}
	for i := 2; i < n; i++ {
		fiboList = append(fiboList, fiboList[i-2]+fiboList[i-1])
	}
	return fiboList
}

func square(n int) int {
	return n * n
}

// 在 main 被执行之前，所有依赖的 package 的 init 方法都会被执行
// 每个包可以有多个 init
// 不同包的 init 函数按照包导入的依赖关系决定执行顺序
// 包的每个源文件可以有多个 init 函数
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
