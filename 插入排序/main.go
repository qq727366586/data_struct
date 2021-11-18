/**
 *@Author luojunying
 *@Date 2021-11-18 23:55
 */
package main

import "fmt"

func index(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j + 1] {
				swap(arr, j, j + 1)
			}
		}
	}
	list(arr)
}

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func list(arr []int) {
	for k, v := range arr {
		if k == len(arr) - 1 {
			fmt.Print(v)
		} else {
			fmt.Print(fmt.Sprintf("%d,", v))
		}
	}
}

func main() {
	arr := []int{2,3,4,5,3,4,12,534,53,2,412,16,7,8,86,4,42,4,5,1}
	index(arr)
}