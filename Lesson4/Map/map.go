package main

import "fmt"

func main() {
	//利用map实现统计数字的功能
	//自然数和该数出现的次数
	//1.定义一个map[int]int
	var n, m = 0, 0
	fmt.Scanln(&n)
	var numCount = make(map[int]int, n)

	//2.将要输入的数字放入一个切片中
	var numbers []int
	for i := 0; i < n; i++ {
		fmt.Scanln(&m)
		numbers = append(numbers, m)
	}
	//fmt.Println(numbers)

	//3.遍历数字切片做统计
	for _, number := range numbers {
		v, ok := numCount[number]
		if ok {
			//map中有这个数字的统计记录,次数在原来的基础上加1
			numCount[number] = v + 1
		} else {
			//map中没有这个数字的统计记录，初始化为1
			numCount[number] = 1
		}
	}
	fmt.Println()
	for k, v := range numCount {
		fmt.Println(k, v)
	}

}
