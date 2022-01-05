/**
 *@Author luojunying
 *@Date 2021-12-13 21:25
 *@desc 大堆
 */
package main

import "fmt"


func BigHeapInsert(arr []int, index int) {
	point := (index - 1) / 2
	for arr[index] > arr[point] {
		//交换
		arr[index], arr[point] = arr[point], arr[index]
		index = point
		point = (index - 1) / 2
	}
}

func BigHeapFy(arr []int, index, heapSize int) {
	//先查看是否有左孩子
	leftChild := 2 * index + 1
	for leftChild < heapSize {
		//找出左右孩子中最大的那个
		largest := leftChild
		if leftChild + 1 < heapSize && arr[leftChild + 1] > arr[leftChild] {
			largest = leftChild + 1
		}
		//找出孩子与父节点谁大
		if arr[largest] < arr[index] {
			largest = index
		}
		//如果父节点的大 那么没必要比下去了
		if largest == index {
			return
		}
		//否则继续比较
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		leftChild = 2 * index + 1
	}
}


func BigHeapInsert2(arr []int, index, heapSize int) {
	//先找到左孩子
	leftChild := 2 * index + 1
	//heapFy过程
	for leftChild < heapSize {
		largest := leftChild
		if leftChild + 1 < heapSize && arr[leftChild + 1] > arr[leftChild] {
			largest = leftChild + 1
		}
		if arr[largest] < arr[index] {
			largest = index
		}
		if arr[largest] == arr[index] {
			return
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		leftChild = 2 * index + 1
	}
}

//堆排序
func BigHeapSort(arr []int, index, heapSize int) {
	//先形成堆
	for i := len(arr) - 1; i >= 0; i-- {
		BigHeapInsert2(arr, i, len(arr))
	}
	//然后在取出
	for heapSize > 0 {
		arr[index], arr[heapSize - 1] = arr[heapSize - 1], arr[index]
		//heapSize--
		heapSize--
		//然后在heapFy
		BigHeapFy(arr, index, heapSize)
	}
}

//小根堆heapInsert
func SmallHeapInsert(arr []int, index, heapSize int){
	//找出父节点
	point := (index - 1) / 2
	//
	for arr[index] < arr[point] {
		arr[index], arr[point] = arr[point], arr[index]
		index = point
		point = (index - 1) / 2
	}
}

//小跟堆heapFy
func SmallHeapFy(arr []int, index, heapSize int) {
	//先找左孩子
	leftChild := index * 2 + 1
	//最小
	for leftChild < heapSize {
		smallest := leftChild
		if leftChild + 1 < heapSize && arr[leftChild + 1] < arr[leftChild] {
			smallest = leftChild + 1
		}
		if arr[smallest] > arr[index] {
			smallest = index
		}
		if arr[smallest] == arr[index] {
			return
		}
		//否则继续向下寻找
		arr[smallest], arr[index] = arr[index], arr[smallest]
		index = smallest
		leftChild = index * 2 + 1
	}
}

  //小跟堆排序
func SmallHeapSort(arr []int, index, heapSize int) {
	//先形成小跟堆
	for i := len(arr) - 1; i >= 0; i-- {
		SmallHeapFy(arr, i, heapSize)
	}
	fmt.Println(arr)
	//heapFy
	for heapSize > 0 {
		arr[index], arr[heapSize - 1] = arr[heapSize - 1], arr[index]
		heapSize--
		SmallHeapFy(arr, index, heapSize)
	}
}


func main() {
	var arr = []int{1,4,5,2,3,5,5,1,2,10}
	/*for i := 0; i < len(arr); i++ {
		SmallHeapInsert(arr, i, len(arr))
	}
	BigHeapSort(arr, 0, len(arr))
	fmt.Println(arr)
	return*/
	SmallHeapSort(arr, 0, len(arr))
	fmt.Println(arr)
}
