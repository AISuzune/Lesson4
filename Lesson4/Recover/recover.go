package main

import (
	"fmt"
	"strings"
)

//type Node struct {
//	Next  *Node
//	Value interface{}
//}
//
//var first *Node
//
//func main() {
//	visited := make(map[*Node]bool)
//	for n := first; n != nil; n = n.Next {
//		if visited[n] {
//			fmt.Println("cycle detected")
//			break
//		}
//		visited[n] = true
//		fmt.Println(n.Value)
//	}
//}

//func main() {
//	var s []string
//	for i := 0; i < 9999; i++ {
//		go func() {
//			s = append(s, "袁神又在玩原了！")
//		}()
//	}
//
//	fmt.Printf("玩 %d 次原神", len(s))
//}

//var counter = struct {
//	sync.RWMutex
//	m map[string]int
//}{m: make(map[string]int)}
//
//func main() {
//	s := make(map[string]string)
//	for i := 0; i < 99; i++ {
//		go func() {
//			s["袁神"] = "原"
//		}()
//	}
//
//	fmt.Printf("玩了 %d 次原神", len(s))
//}

//var m sync.Map
//
//func main() {
//	//写入
//	data := []string{"jjz", "hhz", "rrz"}
//	for i := 0; i < 4; i++ {
//		go func(i int) {
//			m.Store(i, data[i])
//		}(i)
//	}
//	time.Sleep(time.Second)
//
//	//读取
//	v, ok := m.Load(0)
//	fmt.Printf("Load: %v, %v\n", v, ok)
//
//	//删除
//	m.Delete(1)
//
//	//读或写
//	v, ok = m.LoadOrStore(1, "xxz")
//	fmt.Printf("LoadOrStore: %v, %v\n", v, ok)
//
//	//遍历
//	m.Range(func(key, value interface{}) bool {
//		fmt.Printf("Range: %v, %v\n", key, value)
//		return true
//	})
//}

//	func alignSize(nums []int) int {
//		size := 0
//		for _, n := range nums {
//			if s := int(math.Log10(float64(n))) + 1; s > size {
//				size = s
//			}
//		}
//
//		return size
//	}
//
//	func main() {
//		nums := []int{12, 237, 3878, 3}
//		size := alignSize(nums)
//		for i, n := range nums {
//			fmt.Printf("%02d %*d\n", i, size, n)
//		}
//	}
//func main() {
//	timer := time.NewTimer(10 * time.Second)
//
//	ticker := time.NewTicker(1 * time.Second)
//	defer ticker.Stop()
//	for {
//		select {
//		case <-timer.C:
//			fmt.Println("Done")
//			return
//		case t := <-ticker.C:
//			fmt.Println("ticker at: ", t)
//		}
//	}
//}

//func main() {
//	str := "我是袁鑫浩"
//	new := "袁神"
//	old := "袁鑫浩"
//	n := 1
//	println(strings.Replace(str, old, new, n))
//}

//func main() {
//
//	str := "Golang is cool, right?"
//	str1 := "一起玩原神吗铁汁"
//	var manyG string = "o"
//
//	fmt.Printf("%d\n", strings.Count(str, manyG))
//	fmt.Printf("%d\n", strings.Count(str, "oo"))
//	fmt.Printf("%d\n", strings.Count(str1, "袁神"))
//}

//func main() {
//	str := "一起玩嘛？袁神"
//	fmt.Printf("%d\n", len([]rune(str)))
//	fmt.Println(utf8.RuneCountInString(str))
//}

//func main() {
//	str := "Hello World!\n"
//	fmt.Printf("%s", str)
//	fmt.Printf(strings.ToLower(str))
//	fmt.Printf(strings.ToUpper(str))
//}

func main() {
	str := "袁神 Hello Go 袁神"
	str1 := "          请使用TrimSpace()修剪空白字符       "
	fmt.Println(str)
	fmt.Printf("%q\n", strings.Trim(str, "袁神"))
	fmt.Printf("%q\n", strings.Trim(str, "袁神 "))

	fmt.Printf("%q\n", strings.TrimLeft(str, "袁神 "))
	fmt.Printf("%q\n", strings.TrimRight(str, " 袁神"))
	fmt.Printf("%q\n", strings.TrimSpace(str1))
}

//func main() {
//	str := "This is Golang Project"
//	fmt.Println(str)
//	fmt.Printf("%q\n", strings.Split(str, "Project")) //分割指定字符
//	fmt.Printf("%q\n", strings.Split(str, " "))       //分割空白字符
//}
