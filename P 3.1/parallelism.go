package main

import (
	"fmt"
	"sync"
	"time"
)

// Задача 1
// При каждом запуске кода выдаются разные числа, у меня после 4 запусков выдавались числа: 930, 990, 990, 982
// Код запускает 1000 горутин , каждая из которых пытается увеличить переменную "counter" на 1
// Все это просходит из-за "data race", то есть все горутины одновременно чистают и записывают в "counter" без синхронизации,
// также все горутины выполняются конкурентно
func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	fmt.Println(counter)
}

// Вариант для исправления
func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// Задача 2
// Код падает в deadlock после строки "ch <- 1"
// Так происходит потому что канал не буфферизирован и он блокирует отправку, пока другая горутина не начнет ее читать
// но так как горутина "go func()" стоит после строчки "ch <- 1" она запускается после блокировки и из-за этого происходит зависание
func main() {
	ch := make(chan int)
	ch <- 1
	go func() {
		fmt.Println(<-ch)
	}()
}

// Исправленный вариант
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	ch <- 1
	wg.Wait()

}

// Задача 3
// Код отрабатывает и выдает "x[1] = 2", но это неккоректно
// Здесь создается map b запускает 10 горутин, которые одновременно пытаются записать в нее данные, без синхронизации, возникает "data race"
// также решение в виде "time.Sleep" не гарантирует, что все записи успеют отработать
// Вообще у меня за 2 запуска не выдалась "panic" так как версия Go не всегда вызывает "panic" при конкуретной записи, так что это просто везение
func main() {
	x := make(map[int]int, 1)
	go func() { x[1] = 2 }()
	go func() { x[3] = 7 }()
	go func() { x[123] = 10 }()
	go func() { x[1] = 2 }()
	go func() { x[34] = 7 }()
	go func() { x[1432] = 10 }()
	go func() { x[1] = 2 }()
	go func() { x[100] = 7 }()
	go func() { x[34] = 10 }()
	go func() { x[1] = 2 }()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("x[1] =", x[1])
}

//Вариант исправления

func main() {
	x := make(map[int]int, 1)
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func(key, value int) {
			mu.Lock()
			x[key] = value
			mu.Unlock()
		}(i, i*10)
	}
	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	fmt.Println("x[1] =", x[1])
	mu.Unlock()
}
