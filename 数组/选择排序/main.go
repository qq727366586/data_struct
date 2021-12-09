/**
 *@Author luojunying
 *@Date 2021-11-18 21:44
 */
package main

import "fmt"

func index(arr []uint) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
	list(arr)
}

func swap(arr []uint, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func list(arr []uint) {
	for k, v := range arr {
		if k == len(arr) - 1 {
			fmt.Print(v)
		} else {
			fmt.Print(fmt.Sprintf("%d,", v))
		}
	}
}
func main() {
	arr := []uint{2,3,4,5,3,4,12,534,53,2,412,16,7,8,86,4,42,4,5}
	index(arr)
}