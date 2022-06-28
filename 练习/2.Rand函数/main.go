/**
 *@Author luojunying
 *@Date 2022-02-21 23:59
 */
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randNumber() {
	number := 3
	count := 0
	for i := 0; i < 10000; i++ {
		if int(xToXPower2()) < number {
			count++
		}
	}
	fmt.Println(count)
	fmt.Printf("%.2f" , float64(count / 10000))
}

func xToXPower2() float64 {

	rand.Seed(time.Now().UnixNano())
	return math.Max(float64(rand.Intn(10)), float64(rand.Intn(10) ))
}

func main() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println(xToXPower2())
	}
	//randNumber()
}
