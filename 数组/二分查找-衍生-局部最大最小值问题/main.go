/**
 *@Author luojunying
 *@Date 2021-11-28 19:26
 */
package main

import (
	"fmt"
)

//在一个无序数组中,求峰值问题 也就是 局部最小值 和 局部最大值
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
		if arr[mid] <= arr[mid + 1] {
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
	out := getPartMaxFromArr([]int{110,110,110,110,110,110,110})
	fmt.Print(out)
}