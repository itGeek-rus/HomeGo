package main

import (
	"fmt"
	"sync"
	"time"
)

// Результат: 10000 запросов сработало корректно и код выполняется не более чем за "9.3702ms"
// Создается константа со значением "10000", создается счетчик с типом данных "int" и создается "Mutex", чтобы блокировать доступ к переменной "count"
// Создается функция "networkRequest", в которой имитируется задержка в "1ms", "m.Lock() и m.Unlock()" сначала блокирует запись и потом когда запись произведена, разблокирует
// В функции "main" создается счетчик "WaitGroup" в этот счетчик добавляется 10000 запросов, затем создается счетчик на количество горутин, "defer wg.Done()" уменьшает счетчик горутин
// каждая горутина вызывает функцию "networkRequest()"
// "wg.Wait()" ожидает завершения всех 10000 горутин и после этого выдается результат
// "start := time.Now()", "finish := time.Since(start)" создал для подсчета скорости выполнения
const numRequests = 10000

var count int

var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	wg.Wait()
	finish := time.Since(start)
	fmt.Println(count)
	fmt.Println(finish)
}

func networkRequest() {
	time.Sleep(time.Millisecond)
	m.Lock()
	count++
	m.Unlock()
}
