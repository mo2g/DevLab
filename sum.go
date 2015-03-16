package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
n 内 能被 2、3、5、6 整除的数 相加
*/

func main() {
	runtime.GOMAXPROCS(2)
	x := 0
	time1 := time.Now()
	x = _sum(1, 10000000)
	time2 := time.Now()
	x = go_sum(3, 10000000)
	time3 := time.Now()
	fmt.Println(x)
	fmt.Println(time2.Sub(time1))
	fmt.Println(time3.Sub(time2))
}

func _sum(start int, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%6 == 0 {
			sum += i
		}
	}
	return sum
}

func chan_sum(start int, end int, totle chan int) {
	sum := 0
	for i := start; i < end; i++ {
		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%6 == 0 {
			sum += i
		}
	}
	totle <- sum
}

/*123	456	789	10*/
func go_sum(num int, max int) int {
	sum := 0
	y := max / num
	start := 0
	end := 0

	if max%num != 0 {
		num += 1
	}
	totle := make(chan int, num)

	for i := 0; i < num; i++ {
		start = y * i
		end = y * (i + 1)

		if start == 0 {
			start = 1
		}

		if end >= max {
			end = max
		}
		go chan_sum(start, end, totle)
	}
	for i := 0; i < num; i++ {
		sum += <-totle
	}
	return sum
}
