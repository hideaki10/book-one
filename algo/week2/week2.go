package main

import "fmt"

// 选择排序
// 总是拿最小的元素
// 可否原地排序 不需要另外的内存空间
//
// 插入排序

var list = []int{10, 8, 20, 5, 3, 6, 7, 9, 1, 2}

func swap(list []int, i, minIndex int) {
	temp := list[i]
	list[i] = list[minIndex]
	list[minIndex] = temp
}

func main() {
	for i := 0; i < len(list); i++ {
		minIndex := i
		for j := i; j < len(list); j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}
		swap(list, i, minIndex)
	}

	for i, v := range list {
		fmt.Println(i, v)
	}
}
