/**
 *@Author luojunying
 *@Date 2021-11-18 23:46
 */
package main

import "fmt"

func index(arr []uint) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
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