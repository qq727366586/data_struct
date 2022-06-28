/**
 *@Author luojunying
 *@Date 2021-11-28 20:26
 */
package main

import "fmt"

//前提 有序数组
func getLeftNumber(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	k := -1
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid] >= key {
			k = mid
			r = mid

		} else {
			l = mid + 1
		}
	}
	return k
}

// 二分查找 查找最右侧最小值下标
func getRightNumber(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	index := -1
	for l <= r {
		mid := l + (r-l)>>1
		if arr[mid] > key {
			r = mid - 1
		}
		if arr[mid] < key {
			l = mid + 1
		}
		// 说明右侧还有可能存在
		if arr[mid] == key {
			index = mid
			l = mid + 1
		}
	}
	return index
}

func main() {
	out := getRightNumber([]int{110, 111, 111, 111, 111, 111, 111}, 111)
	fmt.Print(out)
}
