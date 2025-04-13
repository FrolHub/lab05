package sweet

import (
	"fmt"
"example.com/food"
)

type Sweet struct {
    food.Food
    Sugar uint64
}

func NewSweet(name string, ccal uint64, price float64, sugar uint64) *Sweet {
    return &Sweet{food.Food{name, ccal, price}, sugar}
}

func (s *Sweet) Eaten() {
    fmt.Printf("%s: меня съели !\n", s.Name)
}

func (s *Sweet) Fall() {
    fmt.Printf("%s упала на пол. \n", s.Name)
}

func (s *Sweet) GetInfo() {
    fmt.Printf("%s содержит %d калорий, %d граммов сахара и стоит %.2f.\n", s.Name, s.Ccal, s.Sugar, s.Price)
}

func (s *Sweet) AddSugar(amount uint64) {
    s.Sugar += amount
    fmt.Printf("%d граммов сахара добавлено в %s.\n", amount, s.Name)
}

func (s *Sweet) DecreaseSugar(amount uint64) {
    s.Sugar -= amount
    fmt.Printf("%d граммов сахара убрано из %s.\n", amount, s.Name)
}

func (s *Sweet) UpdateSugar(amount uint64) {
    s.Sugar = amount
    fmt.Printf("Содержание сахара обновлено до %d граммов в %s.\n", amount, s.Name)
}
