// 11.2 객체 초기화 시에 DI하는 방법

package main

import (
	"fmt"
)

// ServiceImpl 구조체 정의
type ServiceImpl struct{}

// ServiceImpl에 대한 Apply 메서드 구현
func (s *ServiceImpl) Apply(id int) error {
    fmt.Printf("ServiceImpl.Apply called with id: %d\n", id)
    return nil
}

// OrderService 인터페이스 정의
type OrderService interface {
    Apply(int) error
}

// Application 구조체 정의
type Application struct {
    os OrderService
}

// NewApplication 함수 (생성자) 정의
func NewApplication(os OrderService) *Application {
    fmt.Println("NewApplication called, OrderService implementation injected")
    return &Application{os: os}
}

// Application 구조체의 Apply 메서드 구현
func (app *Application) Apply(id int) error {
    fmt.Printf("Application.Apply called with id: %d\n", id)
    return app.os.Apply(id)
}

// main 함수
func main() {
    // NewApplication 함수를 호출하여 의존성 주입
    app := NewApplication(&ServiceImpl{})
    // 주입된 의존성을 사용하여 Apply 메서드 호출
    app.Apply(19)
}