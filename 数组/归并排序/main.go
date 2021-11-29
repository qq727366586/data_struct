package main

func MergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r -l) >> 1
	MergeSort(arr, l, arr[mid])
	MergeSort(arr, arr[mid] + 1, r)
	Merge(arr, mid, l , r)
}

func Merge(arr []int, mid, l, r int) {
	var temp []uint
	i := l
	y := r

}

func main() {

}
