/**
 *@Author luojunying
 *@Date 2021-11-18 21:58
 */
package main

import "fmt"

/**
 *  题目1:
 *   一个数组 只有一种数出现了奇数次,其他的所有数出现了偶数次, 找出这个出现奇数次的 数字
 */
func index(arr []uint) uint {
	if len(arr) < 3 {
		return 0
	}
	temp := uint(0)
	for i := 0; i < len(arr); i++ {
		temp = temp ^ arr[i]
	}
	return temp
}

/**
 *  题目2:
 *   一个数组 两种数出现了奇数次,其他的所有数出现了偶数次, 找出这两个出现奇数次的 数字
 */
func index1(arr []int) (int, int) {
	if len(arr) < 3 {
		return 0, 0
	}
	temp := 0
	//结果是 a ^ b 这两个奇数
	for i := 0; i < len(arr); i++ {
		temp = temp ^ arr[i]
	}
	//结果是 a ^ b
	//思路: 找出他们最右侧 位数为 1 的位置, temp &（^temp + 1） = 找出最右侧为 1 的
	eor := temp & (^temp + 1)
	temp1 := 0
	for i := 0; i <len(arr); i++ {
		if arr[i] & eor == 0 {
			temp1 ^= arr[i]
		}
	}
	return temp1, temp ^ temp1
}


func main() {
	fmt.Println("---------------------------------1.")
	arr := []uint{2,3,2,4,5,5,4,7,8,6,7,8,6}
	res := index(arr)
	fmt.Print(res)
	fmt.Println()
	fmt.Println("---------------------------------2.")
	arr1 := []int{2,3,2,4,5,5,4,7,8,6,7,8,6,10}
	res1, res2 := index1(arr1)
	fmt.Print(res1, res2)
}