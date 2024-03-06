package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Паттерн "состояние" используется когда объект меняет свое поведение в зависимости от состояния. Это делает его
похожим на конечный автомат.
	Плюсы:
		- Избавляет от множества больших условных операторов, связанных с различными состояниями объекта.
		- Позволяет переключаться между состояниями объекта во время выполнения.
		- При добавлении нового состояния может потребоваться изменение только предыдущего состояния.
	Минусы:
		- Может привести к созданию большого количества классов.
		- Усложняет отладку программы из-за зависимости от текущего состояния.
	Примером использования паттерна "состояние" может быть видео плеер. У видео будет два состояния "воспроизведение" и "пауза".
Когда видео находится в состоянии "воспроизведение", нажатие на кнопку "пробел" переведет его в состояние "пауза", и наоборот.
*/

type State interface {
	Handle(context *Context)
}

type PlayVideo struct {
}

func (PlayVideo) Handle(context *Context) {
	fmt.Println("Setting pause")
	context.SetState(PauseVideo{})
}

type PauseVideo struct {
}

func (PauseVideo) Handle(context *Context) {
	fmt.Println("Starting video")
	context.SetState(PlayVideo{})
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	context := Context{
		state: PlayVideo{},
	}

	context.Request() // Имитация нажатия пробела
	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
