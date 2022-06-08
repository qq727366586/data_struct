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


func main () {

}
