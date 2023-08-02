package main

import "fmt"

func deleteAtIndex[T any](s []T, index int) []T {
	// 删除切片的指定索引元素
	s = append(s[:index], s[index+1:]...)

	// 如果切片的长度小于其容量的一半，那么进行缩容
	if len(s) < cap(s)/2 {
		newSlice := make([]T, len(s))
		copy(newSlice, s)
		s = newSlice
	}

	return s
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	str := []string{"a", "b", "c", "d"}

	nums = deleteAtIndex(nums, 2)
	str = deleteAtIndex(str, 1)

	fmt.Println(nums) // [1 2 4 5]
	fmt.Println(str)  // ["a", "c", "d"]
}
