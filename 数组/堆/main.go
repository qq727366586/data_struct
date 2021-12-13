/**
 *@Author luojunying
 *@Date 2021-12-13 21:25
 *@desc 大堆
 */
package main

import "fmt"


func heapInsert(arr []int, index int) {
	 //找父节点
	 i := (index - 1) / 2
	 //是否必父节点大
	 for arr[index] > arr[i] {
	 	//交换
	 	arr[i], arr[index] = arr[index], arr[i]
	 	index = i
	 	i = (index - 1) / 2
	 }
}

func heapInsert2(arr []int, index, heapSize int) {
	left := 2 * index + 1
	for left < heapSize {
		largest := left
		if left + 1 < heapSize && arr[left + 1] > arr[left] {
			largest = left + 1
		}
		// 父 与 子 比较
		if arr[largest] < arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2 * index + 1
	}
}

//某个数在index位置,能否往下移动
func heapFy(arr []int, index, heapSize int) {
	//先判断这个数是否有左孩子
	left := 2 * index + 1
	//找出左右两个孩子最大的值
	for left < heapSize {
		largest := left
		if left + 1 < heapSize && arr[left + 1] > arr[left] {
			largest = left + 1
		}
		// 父 与 子 比较
		if arr[largest] < arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2 * index + 1

	}
}

func main() {
	var arr = []int{1,4,5,2,3,5,5,1,2,10}
	/*for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}*/
	for i := len(arr) - 1; i >= 0; i-- {
		heapInsert2(arr, i, len(arr))
	}

	fmt.Println(arr)
}
