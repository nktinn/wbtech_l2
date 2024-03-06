package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Паттерн "фасад" необходим для упрощения взаимодействия с системой. Например, клиент не знает о существовании отдельных
компонентов компьютера. Ему важно включить компьютер, поработать на нём и выключить.
	Примером реального использования паттерна Фасад может быть авторизация пользователя. Клиент вводит логин и пароль,
нажимает кнопку "Войти". За простым интерфейсом и одним действием от него скрываются многие процессы работы с
введенными данными - поиск аккаунта в базе дынных, хэширование пароля и проверка его с указанным в базе, создаение
куки или JWT, авторизация, перенаправление на следующую страницу.
	К плюсам можно отнести:
		- Упрощение взаимодействия с системой путем сокрытия сложной логики
		- Облегчение поддержки кода с возможностью изменения реализации системы без изменения интерфейса клиента
	К минусам можно отнести:
		- Может привести к нарушению принципа единой ответственности, если фасад выполняет слишком много функций
		- Могут возникнуть трудности с отладкой кода, т.к. основная логика спрятана за фасадом
*/

type ICPU interface {
	Start() error
	Stop() error
}

type CPU struct {
	model string
	state bool
}

func (c *CPU) Start() error {
	fmt.Printf("Starting %s CPU...\n", c.model)
	c.state = true
	return nil
}
func (c *CPU) Stop() error {
	fmt.Println("Shutting down CPU...")
	c.state = false
	return nil
}

type IGPU interface {
	Start() error
	Stop() error
}

type GPU struct {
	model string
	state bool
}

func (g *GPU) Start() error {
	fmt.Printf("Starting %s GPU...\n", g.model)
	g.state = true
	return nil
}
func (g *GPU) Stop() error {
	fmt.Println("Shutting down GPU...")
	g.state = false
	return nil
}

type IMemory interface {
	Start() error
	Stop() error
}

type Memory struct {
	model string
	state bool
}

func (m *Memory) Start() error {
	fmt.Printf("Starting %s memory...\n", m.model)
	m.state = true
	return nil
}
func (m *Memory) Stop() error {
	fmt.Println("Shutting down memory...")
	m.state = false
	return nil
}

type Facade struct {
	cpu    ICPU
	gpu    IGPU
	memory IMemory
}

func NewFacade(cpu ICPU, gpu IGPU, memory IMemory) *Facade {
	return &Facade{
		cpu:    cpu,
		gpu:    gpu,
		memory: memory,
	}
}

func (f *Facade) TurnOn() error {
	if err := f.cpu.Start(); err != nil {
		return err
	}
	if err := f.gpu.Start(); err != nil {
		return err
	}
	if err := f.memory.Start(); err != nil {
		return err
	}
	return nil
}

func (f *Facade) TurnOff() error {
	if err := f.cpu.Stop(); err != nil {
		return err
	}
	if err := f.gpu.Stop(); err != nil {
		return err
	}
	if err := f.memory.Stop(); err != nil {
		return err
	}
	return nil
}

func main() {
	cpu := &CPU{
		model: "Apple Silicon M1",
	}
	gpu := &GPU{
		model: "MSI 4080Ti",
	}
	memory := &Memory{
		model: "Hynix 16GB",
	}

	computer := NewFacade(cpu, gpu, memory)

	if err := computer.TurnOn(); err != nil {
		fmt.Println("Error turning on computer:", err)
		return
	}
	fmt.Println("Computer turned on")
	time.Sleep(2 * time.Second)

	if err := computer.TurnOff(); err != nil {
		fmt.Println("Error turning off computer:", err)
		return
	}
	fmt.Println("Computer turned off")
}
