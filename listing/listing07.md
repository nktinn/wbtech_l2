Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Output:
    Программа выведет сначала 1, 2, а потом последовательность оставшихся чисел в произвольном порядке. После этого,
программа будет постоянно выводить 0.

Объяснение:
    Программа создает два канала a и b, которые заполняются числами 1, 3, 5, 7 и 2, 4, 6, 8 соответственно.
Затем программа создает канал c, который объединяет каналы a и b. В цикле for v := range c программа выводит числа из канала c.
После того, как каналы a и b заканчиваются, программа начинает выводить 0, так как в канале c больше нет значений, а
он не был закрыт.

```
