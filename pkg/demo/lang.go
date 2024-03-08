package demo

import (
	"fmt"
	"reflect"
	"sync"
)

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("eat")
}

func invokeByName(name string) {
	a := &Animal{}
	reflect.ValueOf(a).MethodByName(name).Call([]reflect.Value{})
}

func goRoutineCommunication() {
	// a->b->c->a
	var a = make(chan struct{})
	var b = make(chan struct{})
	var c = make(chan struct{})
	var wg sync.WaitGroup

	fa := func() {
		<-c
		fmt.Println("a invoked")
		wg.Done()
		b <- struct{}{}
		fmt.Println("a done")
	}

	fb := func() {
		<-a
		fmt.Println("b invoked")
		wg.Done()
		c <- struct{}{}
		fmt.Println("b done")
	}

	fc := func() {
		<-b
		fmt.Println("c invoked")
		wg.Done()
		// 当a的消费者协程全部运行完退出后，这里会阻塞
		// 向无缓冲的channel发送也会阻塞
		a <- struct{}{}
		fmt.Println("c done")
	}
	wg.Add(3)
	go fa()
	go fb()
	go fc()
	a <- struct{}{}
	wg.Wait()
}

func Run() {
	invokeByName("Eat")
	goRoutineCommunication()
}
