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
		mid := l + (r - l) >> 1
		if arr[mid] >= key {
			k = mid
			r = mid

		} else {
			l = mid + 1
		}
	}
	return k
}


func main()  {
	out := getLeftNumber([]int{110,111,111,111,110,110,110}, 111)
	fmt.Print(out)
}
