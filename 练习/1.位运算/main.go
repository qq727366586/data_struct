/**
 *@Author luojunying
 *@Date 2022-02-21 21:23
 */
package main

import "fmt"

//32位整形 打印出二进制
func num(num int) {
	for i := 31; i >= 0; i-- {
		temp := num & (1 << i)
		if temp == 0 {
			fmt.Print(0)
		} else {
			fmt.Print(1)
		}
	}
}

// 一个数组中 一个数出现了奇数次,其余的都是偶数次 请返回这个奇数次的数
func findOddNumber(arr []int) int {
	var temp int = 0
	for _, v := range arr {
		temp = temp ^ v
	}
	return temp
}

// 一个数组中 有两个数出现了奇数次, 其余出现了偶数次 请返回这两个数
func findTwoOddNumber(arr []int) (int, int) {
	var temp int = 0
	for _, v := range arr {
		temp ^= v
	}
	// 求出这两个数最右侧为1的二进制
	// 假设 temp = 1,  ^temp = ....1110
	// ^temp + 1= 1...0001
	// temp & (^temp + 1) = 0....0001 & 1...0001
	nextTemp := temp & (^temp + 1)
	oneOfOdd := 0
	for _, v := range arr {
		if v&nextTemp != 0 {
			oneOfOdd = temp ^ v
		}
	}
	// 到这里最终temp 会等于 其中一个奇数
	// 求出另外一个
	return oneOfOdd, temp ^ oneOfOdd
}

// 一个数组中 有个数出现了k次 其他数出现了m次,  k < m && m > 1 求出现了K次的数
func findKNumber(arr []int, k, m int) int {
	var tmp = [32]int{}
	for _, v := range arr {
		for j := range tmp {
			if v>>j&1 != 0 {
				tmp[j]++
			}
		}
	}
	var result = 0
	for t, v := range tmp {
		// 说明存在出现了K次的数
		if v%m != 0 {
			result |= 1 << t
		}
	}
	return result
}

func main() {
	arr := []int{0, 22, 22, 33, 33, 22, 33}
	res := findKNumber(arr, 1, 3)
	fmt.Println(res)
}
