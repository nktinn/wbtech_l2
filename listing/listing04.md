Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Output:
0
1
2
3
4
5
6
7
8
9
deadlock

Объяснение:
    Deadlock - это ошибка, которая возникает, когда горутина блокируется, ожидая чтения или записи в канал.
В ситуации выше главная горутина ожидает чтения из канала ch, но вторая горутина, которая должна писать в канал, 
уже завершилась, а канал не был закрыт. Таким образом, главная горутина блокируется и программа завершается с ошибкой deadlock.

```
