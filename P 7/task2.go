package main

import (
	"fmt"
	"sync"
)

// Код уходит в "deadlock"
// Изначально создается небуферизированный канал. Затем создается "WaitGroup", которая ожидает три горутины
// Затем созадется горутина в цикле, в котором идет вычесление "v * v" и попытка отправить результат в "ch"
// но все 3 ггорутины заблокируются на этапе cg <- v * v, потому что канал небуферизированный и нет читателя "<- ch"
// "wg.Wait()" зависнет, потому что горутины не могут завершиться. "defer wg.Done()" не выполниться
// цикл "for v := range ch" не начинается, так как "wg.Wait()" заблокировал функцию "main"
// В итоге все горутины ждут читателя, а функция "main" ждет завершения горутин
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	wg.Wait()
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}

// Исправлененый вариант
// Результат будет "5"
// Чтение из канала вынесено в отдельную горутину, теперь, то что читает наши записи
// "wg.Wait()" теперь не блокирует функцию "main". После "wg.Wait()" канал закрывается "close(ch)", чтобы завершить цикл в горутине
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}

	var sum int
	go func() {
		for v := range ch {
			sum += v
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Printf("result: %d\n", sum)
}
