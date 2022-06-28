package main

import "fmt"

//定某个数 小于等于的放左边 大于的放右边
func Sort(arr []int, num int) {
	point := -1
	for k, _ := range arr {
		if arr[k] < num {
			arr[point + 1], arr[k] = arr[k], arr[point + 1]
			point++
		}
		if arr[k] == num {
			continue
		}
		if arr[k] > num {
			arr[point + 1], arr[k] = arr[k], arr[point + 1]
		}
	}
}

//定某个数 小于它的数放左边 等于它的数放中间 大于他的数放右边
func Sort2(arr []int, num int) {
	left := -1
	right := len(arr)
	for k, _ := range arr {
		if k == right {
			return
		}
		if arr[k] < num {
			arr[left + 1], arr[k] = arr[k], arr[left + 1]
			left++
		}
		if arr[k] == num {
			continue
		}
		if arr[k] > num {
			arr[right - 1], arr[k] = arr[k], arr[right - 1]
		}
 	}
}


func main() {
	var arr = []int{1,2,4,4,6,5,7,8,10}
	Sort(arr, 6)
	fmt.Println(arr)
	Sort2(arr, 2)
	fmt.Println(arr)
}
