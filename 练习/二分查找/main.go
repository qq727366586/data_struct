/**
 *@Author luojunying
 *@Date 2022-03-31 22:20
 */
package main

/**
请实现无重复数字的升序数组的二分查找
给定一个 元素升序的、无重复数字的整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标（下标从 0 开始），否则返回 -1
数据范围： ， 数组中任意值满足
进阶：时间复杂度  ，空间复杂度
*/
func FindIndex(arr []int, target, start, end int) int {
	for start <= end {
		mid := start + (end - start) >> 1
		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target{
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

/**
 求旋转数组最小值
 */
func getPartMaxFromArr(arr []int) int {
	l := 0
	r := len(arr) - 1

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
