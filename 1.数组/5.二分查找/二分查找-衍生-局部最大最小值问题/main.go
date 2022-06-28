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
	// 先判断数组长度
	index := -1
	l := 0
	r := len(arr) - 1
	if len(arr) == 0 {
		return index
	}
	if len(arr) == 1 {
		return 0
	}
	if len(arr) == 2 && arr[0] < arr[1] {
		return l
	}
	if arr[r] < arr[r-1] {
		return r - 1
	}
	for l < r {
		mid := l + (r-l)>>1
		// 处于 \/ 中间
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		}
		// 处于 / 中间
		if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			r = mid - 1
		}
		// 处于 \ 中间
		if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
			l = mid + 1
		}
		// 处于 /\中间
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			r = mid
		}
	}
	return -1
}

func getPartMaxFromArr2(arr []int) int {
	// 先判断数组长度
	index := -1
	l := 0
	r := len(arr) - 1
	mid := 0
	if len(arr) == 0 {
		return index
	}
	if len(arr) == 1 {
		return 0
	}
	if len(arr) == 2 || arr[0] < arr[1] {
		return l
	}
	if arr[r] < arr[r-1] {
		return r
	}
	for l < r {
		mid = l + (r-l)>>1
		if arr[mid] > arr[mid+1] {
			l = mid + 1
		} else if arr[mid] > arr[mid-1] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return l
}

func main() {
	//   \
	//	  \/\
	//	   	 \
	//		  \/
	//
	arr := []int{345, 1, 23}
	out := getPartMaxFromArr2(arr)
	fmt.Println(out)

}
