/**
 *@Author luojunying
 *@Date 2022-03-31 22:37
 */
package main

import (
	"fmt"
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []int{-1,0,3,4,6,10,13,14}
	out := FindIndex(arr, 14,0, len(arr) - 1)
	fmt.Println(out)
}

func TestFindIndex2(t *testing.T) {
	arr := []int{3,4,5,1,2,3}
	out := getPartMaxFromArr(arr)
	fmt.Println(out)
}