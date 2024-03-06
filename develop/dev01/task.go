package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting NTP time: ", err)
		os.Exit(1)
	}
}

func GetTime(server string) error {
	time, err := ntp.Time(server)
	if err != nil {
		return err
	}
	fmt.Printf("Current time: %s\n", time)
	return nil
}
