package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Handler(command string) {
	switch strings.Split(command, " ")[0] {

	case "cd":
		cd(command)
	case "pwd":
		pwd()
	case "echo":
		echo(command)
	case "kill":
		kill(command)
	case "Exit":
		exit()
	default:
		fmt.Println("Command not found")
	}
}

func cd(command string) {
	err := os.Chdir(strings.Replace(command, "cd", "", 1))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Directory changed")
	}
}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

func echo(command string) {
	fmt.Println(strings.Replace(command, "echo", "", 1))
}

func kill(command string) {
	pid, err := strconv.Atoi(strings.Replace(command, "kill", "", 1))
	if err != nil {
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	proc.Kill()
}

func exit() {
	os.Exit(0)
}

func main() {
	fmt.Print("Shell: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		Handler(scanner.Text())
		fmt.Print("Shell: ")
	}
}
