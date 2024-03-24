// 11.3 setter를 준비해서 DI하는 방법

package main

import (
	"fmt"
)

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error {
    fmt.Printf("ServiceImpl.Apply called with id: %d\n", id)
    return nil
}

type OrderService interface {
    Apply(int) error
}

type Application struct {
    os OrderService
}

func (app *Application) Apply(id int) error {
    fmt.Printf("Application.Apply called with id: %d\n", id)
    return app.os.Apply(id)
}

func (app *Application) SetService(os OrderService) {
    fmt.Println("SetService called, OrderService implementation set.")
    app.os = os
}

func main() {
    app := &Application{}
    app.SetService(&ServiceImpl{})
    app.Apply(19)
}