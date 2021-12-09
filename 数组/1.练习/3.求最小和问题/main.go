package main

import "fmt"

func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)>>1
	return process(arr, l, mid) + process(arr, mid+1, r) + merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) int {
	temp := make([]int, r-l + 1)
	res := 0
	i := 0
	t1 := l
	t2 := m + 1
	for t1 <= m && t2 <= r {
		if arr[t1] <= arr[t2] {
			//计算有多少个比他小的
			res += arr[t1] * (r - t2 + 1)
			temp[i] = arr[t1]
			i++
			t2++
		} else {
			temp[i] = arr[t2]
			i++
			t1++
		}
	}
	for t1 <= m {
		temp[i] = arr[t1]
		i++
		t1++
	}
	for t2 <= r {
		temp[i] = arr[t2]
		i++
		t2++
	}
	for k, v := range temp {
		arr[k+l] = v
	}
	fmt.Println(arr)
	return res
}

func main() {
	var arr = []int{1,3,4,66,5,2,11}
	out := process(arr, 0, len(arr) - 1)
	fmt.Println(arr)
	fmt.Println(out)
}