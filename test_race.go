package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := make(chan int)
	results := make(chan int)

	go makeRandNums(nums)
	go powerNums(nums, results)

	// Имитируем медленное чтение
	for num := range results {
		fmt.Print(num, " ")
		time.Sleep(time.Millisecond * 10) // Задержка при чтении
	}
	fmt.Println()
}

func makeRandNums(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(10)
	}
	fmt.Println(nums)
	for _, num := range nums {
		ch <- num
	}
	close(ch)
}

func powerNums(in chan int, out chan int) {
	for num := range in {
		out <- num * num
		// time.Sleep(time.Millisecond * 1) // Можно добавить задержку для демонстрации
	}
	close(out)
	fmt.Println("powerNums: channel closed")
}