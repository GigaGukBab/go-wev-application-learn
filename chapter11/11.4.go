// 11.4 메서드(함수) 호출 시에 DI 하는 방법

package main

import (
	"fmt"
)

// ServiceImpl 구조체 정의
type ServiceImpl struct{}

// ServiceImpl 구조체에 대한 Apply 메서드 구현
func (s *ServiceImpl) Apply(id int) error {
    fmt.Printf("ServiceImpl.Apply called with id: %d\n", id)
    return nil
}

// OrderService 인터페이스 정의
type OrderService interface {
    Apply(int) error
}

// Application 구조체 정의
type Application struct{}

// Application 구조체에 대한 Apply 메서드 구현
// 이 메서드는 OrderService 인터페이스와 id를 매개변수로 받아, 
// OrderService 인터페이스의 Apply 메서드를 호출합니다.
func (app *Application) Apply(os OrderService, id int) error {
    fmt.Println("Application.Apply is called with OrderService and id")
    return os.Apply(id)
}

func main() {
    fmt.Println("Creating Application and calling Apply with ServiceImpl and id 19...")
    app := &Application{}
    // ServiceImpl의 인스턴스와 id 값을 Apply 메서드에 전달하여 호출
    app.Apply(&ServiceImpl{}, 19)
}