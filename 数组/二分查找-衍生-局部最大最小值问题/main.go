/**
 *@Author luojunying
 *@Date 2021-11-28 19:26
 */
package main

import (
	"fmt"
)

//局部最小值问题。数组无序，相邻元素不等
//此题求一个 一个 局部最小值 返回下标
func getPartMaxFromArr(arr []int) int {
	l := 0
	r := len(arr) - 1
	if arr[0] < arr[1] {
		return 0
	}
	if arr[r] < arr[r - 1] {
		return r
	}
	for l < r {
		mid := l + (r - l) >> 1
		if arr[mid] < arr[mid + 1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 || r == len(arr) - 1 {
		return -1
	}
	return l
}

func main()  {
	out := getPartMaxFromArr([]int{11,10,9,8,7,6,11})
	fmt.Print(out)
}