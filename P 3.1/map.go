package main

import "fmt"

//При запуске программы происходит паника, мапе присваивается значение nil
//Созадется переменная "m" с типом данных "map[string] int", но она не была проинициализирована и равна "nil"
//И при попытке записи в "nil" мапа вызывает панику
//В цикле программа пытается подсчитать частоту слова в списке, но сразу же падает
func main() {
	var m map[string]int
	for _, word := range []string{"hello", "world", "from", "the", "best",
		"language", "in", "the", "world"} {
		m[word]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

//Способ исправления

func main() {
	m := make(map[string]int)
	for _, word := range []string{"hello", "world", "from", "the", "best",
		"language", "in", "the", "world"} {
		m[word]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

//Это работает, так как при помощи "make(map[string]int)" создается и инициализируется пустая хэш-таблица и в нее теперь можно добавлять элементы
//Теперь "m[word]++" отрабатывает корректно, так как если ключа нет, то тогда ключ создается со значением 0 и при каждой следующей итерации значение увеличивается на 1
