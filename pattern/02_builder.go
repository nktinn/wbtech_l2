package main

import (
	"fmt"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Паттерн "строитель" необходим для создания экземпляров одной сущности с разными конфигурациями без необходимости
изменять код.
	Плюсы:
		- Разделение процесса создания экземпляра класса на несколько этапов, упрощая понимание и процесс отладки
		- Использование одной сущности без необходимости создания дополнительных
	Минусы:
		- Необходимость создания Builder для каждой конфигурации
	На практике паттерн можно использовать при разработке приложения для пиццерии. Пицца состоит из трех основных
компонентов: тесто, соус, начинка. Таким образом, можно использовать одну сущность, а конфигураций может быть несколько
десятков.
*/

type Wheel struct {
	Tire string
	Size string
	Rim  string
}

type Builder interface {
	BuildTire()
	BuildSize()
	BuildRim()
	GetWheel() *Wheel
}

type TOYOBuilder struct {
	Wheel *Wheel
}

func (m *TOYOBuilder) BuildTire() {
	m.Wheel.Tire = "TOYO PROXES R888"
}

func (m *TOYOBuilder) BuildSize() {
	m.Wheel.Size = "205/55 R16"
}

func (m *TOYOBuilder) BuildRim() {
	m.Wheel.Rim = "OZ Racing Estrema GT HLT"
}

func (m *TOYOBuilder) GetWheel() *Wheel {
	return m.Wheel
}

type MichelinBuilder struct {
	Wheel *Wheel
}

func (m *MichelinBuilder) BuildTire() {
	m.Wheel.Tire = "Michelin Pilot Sport 4"
}

func (m *MichelinBuilder) BuildSize() {
	m.Wheel.Size = "225/45 R17"
}

func (m *MichelinBuilder) BuildRim() {
	m.Wheel.Rim = "Neo 740"
}

func (m *MichelinBuilder) GetWheel() *Wheel {
	return m.Wheel
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.BuildTire()
	d.builder.BuildSize()
	d.builder.BuildRim()
}

func main() {
	builder1 := &TOYOBuilder{&Wheel{}}
	builder2 := &MichelinBuilder{&Wheel{}}

	director := &Director{}

	director.SetBuilder(builder1)
	director.Construct()
	product1 := builder1.GetWheel()

	fmt.Println(product1)

	director.SetBuilder(builder2)
	director.Construct()
	product2 := builder2.GetWheel()

	fmt.Println(product2)
}
