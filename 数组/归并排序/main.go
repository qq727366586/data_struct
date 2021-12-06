package main

import "fmt"

func MergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	MergeSort(arr, l, mid)
	MergeSort(arr, mid + 1, r)
	Merge(arr, l ,mid, r)
}

func Merge(arr []int, l, mid, r int) {
	temp := make([]int, r - l + 1)
	i := 0
	t1 := l
	t2 := mid + 1
	for t1 <= mid && t2 <= r {
		if arr[t1] <= arr[t2] {
			temp[i] = arr[t1]
			t1++
			i++
		} else {
			temp[i] = arr[t2]
			t2++
			i++
		}
	}
	for t1 <= mid {
		temp[i] = arr[t1]
		t1++
		i++
	}

	for t2 <= r {
		temp[i] = arr[t2]
		t2++
		i++
	}
	for k, v := range temp {
		arr[k + l] = v
	}
	fmt.Println(arr)
}

func main() {
	var arr = []int{2,1,123,4,4,12,312,4,5,3,4,34,23,4}
	MergeSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}
