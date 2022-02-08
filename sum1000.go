package main

import "fmt"

func SumNum(ch chan int, start int, finish int) {
	sum := 0
	for start <= finish {
		sum = sum + start
		start++
	}
	ch <- sum
}

func sum1000(goroutinesNum int) int {
	ch := make(chan int)
	sum := 0
	for i := 0; i < goroutinesNum; i++ {
		go SumNum(ch, i*1000/goroutinesNum+1, (i+1)*1000/goroutinesNum)
		sum = sum + <-ch
	}
	return sum
}

func main() {
	fmt.Println(sum1000(10))
}
