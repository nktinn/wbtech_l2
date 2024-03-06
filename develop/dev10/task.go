package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут на подключение к серверу")
	flag.Parse()

	// Получаем адрес хоста и порт
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: [--timeout=10s] host port")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]

	var conn net.Conn
	var err error

	// Подключаемся к серверу
	start := time.Now()
	for time.Since(start) < *timeout {
		conn, err = net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
		if err != nil {
			continue
		}
	}
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Горутина для чтения данных из сокета и вывода их в STDOUT
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != io.EOF {
				os.Exit(1)
			}
			if err != nil {
				fmt.Println("Error reading from server:", err)
				continue
			}
			fmt.Println(msg)
		}
	}()

	// Горутина для чтения данных из STDIN и записи их в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			_, err := fmt.Fprintf(conn, "%s\n", msg)
			if err != nil {
				fmt.Println("Error writing to server:", err)
				os.Exit(1)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
