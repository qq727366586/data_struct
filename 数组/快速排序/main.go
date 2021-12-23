package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int, left, right int){

	if left < right {
		//找出随机数
		rand.Seed(time.Now().Unix())
		randNumber := rand.Intn(right - left)

		//分区
		l, r := Partition(arr, left, right ,arr[randNumber])
		fmt.Println(arr, l , r, arr[randNumber])
		//左边
		quickSort(arr, left, l - 1)
		//右边
		quickSort(arr, r + 1, right)
	}

}

//定某个数 小于它的数放左边 等于它的数放中间 大于他的数放右边
func Partition(arr []int, left, right, num int) (int, int) {
	k := left - 1
	more := right
	for left < more {
		fmt.Println(left)
		if arr[left] < num {
			arr[left], arr[k + 1] = arr[k + 1], arr[left]
			left++
			k++
		}
		if arr[left] == num {
			k++
		}
		if arr[left] > num {
			arr[more - 1], arr[left] = arr[left], arr[more - 1]
			more--
			k++
		}
	}
	return k + 1, more
}

func main()  {
	var arr = []int{1,2,4,4,6,5,7,8,10}
	quickSort(arr, 0, len(arr) -1 )
	fmt.Println(arr)
}