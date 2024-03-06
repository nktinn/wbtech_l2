package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Паттерн «фабричный метод» используется для создания объектов. Он определяет интерфейс для создания объекта,
но делегирует создание объекта наследникам. Таким образом, классы-наследники могут изменять тип создаваемого объекта.
	Плюсы:
		- Позволяет избавиться от привязки к конкретным типам создаваемых объектов.
		- Позволяет создавать объекты с различными параметрами на основе одного интерфейса.
	Минусы:
		- Усложняет код программы.
		- Необходимо создавать большую фабрику даже для одного объекта.
	Пример использования паттера "фабричный метод"

*/

type Product interface {
	Use()
}

type Creator interface {
	FactoryMethod(string) Product
}

type ConcreteProduct struct {
	name string
}

func (p *ConcreteProduct) Use() {
	fmt.Println("Using " + p.name)
}

type ConcreteCreator struct {
}

func (c *ConcreteCreator) FactoryMethod(name string) Product {
	return &ConcreteProduct{name: name}
}

func main() {
	var creator Creator
	creator = &ConcreteCreator{}

	product := creator.FactoryMethod("product")
	product.Use()
}
