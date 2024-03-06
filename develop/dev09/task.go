package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")

	cmd := exec.Command("wget", args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	cmd.Run()
}
