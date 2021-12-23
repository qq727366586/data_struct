package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int, left, right int){
	if left < right {
		//找出随机数
		time.Sleep(time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		randNumber := rand.Intn(right - left) + left

		arr[randNumber], arr[right] = arr[right], arr[randNumber]
		fmt.Println(arr)
		//分区
		l, r := Partition2(arr, left, right)
		quickSort(arr, left, l - 1)
		quickSort(arr, r + 1, right)
	}

}

//定某个数 小于它的数放左边 等于它的数放中间 大于他的数放右边
func Partition(arr []int, left, right int) (int, int) {
	point := left - 1
	more := right + 1
	num := arr[right]
	for left < more - 1 {
		if arr[left] < num {
			arr[point + 1], arr[left] = arr[left], arr[point + 1]
			point+=1
			left+=1
		} else if arr[left] == num {
			left+=1
		} else if arr[left] > num {
			arr[more - 1], arr[left] = arr[left], arr[more - 1]
			more-=1
		}
	}
	return point, more
}

//定某个数 小于它的数放左边 等于它的数放中间 大于他的数放右边
func Partition2(arr []int, left, right int) (int, int) {
	point := left - 1
	more := right
	for left < more {
		if arr[left] < arr[right]  {
			point++
			arr[point], arr[left] = arr[left], arr[point]
			left++
		}else if arr[left] > arr[right] {
			more--
			arr[more], arr[left] = arr[left], arr[more]
		}else  {
			left++
		}
	}
	fmt.Println(arr)
	return point + 1, more - 1
}

func main()  {
	var arr = []int{1,2,4,4,6,5,7,8,10}
	quickSort(arr, 0, len(arr) -1 )
	fmt.Println(arr)
}