package main

import "fmt"

func main() {
	var nums1 []interface{}
	nums2 := []int{1, 2, 3}
	nums3 := append(nums1,nums2)
	nums4:= []string{"1"}
	nums3 = append(nums3,nums4)
	fmt.Println(len(nums3))
	fmt.Println(111)

	a := []int{1, 2, 3}
	a = append(a, 1)
	fmt.Println(len(a))
}
