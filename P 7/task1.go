package main

import (
	"fmt"
	"sync"
)

// Код при каждом запуске выдает разный результат
// Создаются небуферизированые каналы, затем создаются горутины, которые отправляют булевое значение "true" в свой канал
// Затем "select" выбирает первый канал, который успел отправить данные и выдает сообщение, порядок выполнения горутин не контролируется
func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		ch <- true
	}()

	go func() {
		ch2 <- true
	}()

	go func() {
		ch3 <- true
	}()

	select {
	case <-ch:
		fmt.Println("val from ch")
	case <-ch2:
		fmt.Println("val from ch2")
	case <-ch3:
		fmt.Println("val from ch3")
	}
}

// Исправленный вариант
func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	var mu sync.Mutex
	mu.Lock()

	go func() {
		ch <- true
	}()

	go func() {
		mu.Lock()
		ch2 <- true
	}()

	go func() {
		mu.Lock()
		ch3 <- true
	}()

	select {
	case <-ch:
		fmt.Println("val from ch")
	case <-ch2:
		fmt.Println("val from ch2")
	case <-ch3:
		fmt.Println("val from ch3")
	}
}

//Подключаем пакет "sync.Mutex", блокируем ненужыне горутины при помощи "mu.Lock()" и в результате выводится только первый канал
