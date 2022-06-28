/**
 *@Author luojunying
 *@Date 2021-11-28 17:43
 */
package main

import "fmt"

func maxArr(arr []int, l int, r int) int {
	//如果长度为1
	if len(arr) == 1 {
		return arr[0]
	}
	if l == r {
		return arr[l]
	}
	//取中间值
	mid := l + ((r - l) >> 1)
	leftMax := maxArr(arr, l, mid)
	rightMax := maxArr(arr, mid + 1, r)
	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}

func main() {
	arr := []int{1,6,5,7,8,2,1}
	res := maxArr(arr, 0, len(arr) - 1)
	fmt.Println(res)
}
