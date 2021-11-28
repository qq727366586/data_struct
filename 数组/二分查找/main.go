/**
 *@Author luojunying
 *@Date 2021-11-28 19:03
 */
package main

import "fmt"

//二分查找,数组必须有序
func binarySearch(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r - l) >> 1
		if key > arr[mid] {
			l = mid + 1
		}
		if key < arr[mid] {
			r = mid - 1
		}
		if key == arr[mid] {
			return mid
		}
 	}
 	return -1
}

func main() {
	out := binarySearch([]int{1,2,3,4,5,6,7,8,9}, 19)
	fmt.Print(out)
}
