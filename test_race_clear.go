package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := make(chan int, 5)      // Буферизированный!
	results := make(chan int, 5)   // Буферизированный!

	go makeRandNums(nums)
	go powerNums(nums, results)

	// Читаем с проверкой на закрытие канала
	for {
		num, ok := <-results
		if !ok {
			fmt.Println("\nMain: channel closed!")
			break
		}
		fmt.Print(num, " ")
		time.Sleep(time.Millisecond * 200) // Очень медленно
	}
	fmt.Println("Main: done")
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
	fmt.Println("makeRandNums: finished and closed nums")
}

func powerNums(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
	fmt.Println("powerNums: finished and closed results")
}