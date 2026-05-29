package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Буферизированные каналы - потенциальная гонка!
	nums := make(chan int, 3)
	results := make(chan int, 3)

	go makeRandNums(nums)
	go powerNums(nums, results)

	// Имитируем очень медленное чтение
	for i := 0; i < 10; i++ {
		num := <-results
		fmt.Print(num, " ")
		time.Sleep(time.Millisecond * 100) // Длинная задержка
	}
	fmt.Println()
}

func makeRandNums(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(10)
	}
	fmt.Println("Generated:", nums)
	for _, num := range nums {
		ch <- num
	}
	close(ch)
	fmt.Println("makeRandNums: channel closed")
}

func powerNums(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
	fmt.Println("powerNums: channel closed")
}