package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Результат выполнения после нескольких запусков: "Ошибка: Время таймаута вышло" и "174, 4540, 4648, 448"
// "predictableFunc()" делаем такой, чтобы возвращала, помимо "int64",  "error"
// Создаем контекст с таймаутом в секунду, создаем "defer cancel()", чтобы освободить ресурсы, создаем небуферизированный канал, который принимает тип "int64"
// "unpredictableFunc()" используем в горутине и пишем в наш канал. Сначала какое-то время выполняется функция "unpredictableFunc" и затем она пишет какое-то число в канал
// Затем используем "select" и тут получаем либо результат из канала, либо отмену от контекста
// Числа случайные, так как "unpredictableFunc" спит от 0 до 5 секунд (5000 * 0,001) и таймаут всего в 1 секунду
func main() {
	fmt.Println("started")
	res, err := predictableFunc()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)
	}
	fmt.Println(unpredictableFunc())
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

func predictableFunc() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	chanResult := make(chan int64)

	go func() {
		chanResult <- unpredictableFunc()
	}()

	select {
	case result := <-chanResult:
		return result, nil
	case <-ctx.Done():
		return 0, fmt.Errorf("Время таймаута вышло")
	}
}
