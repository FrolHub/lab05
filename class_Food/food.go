package food

import "fmt"

type Food struct {
    Name  string
    Ccal  uint64
    Price float64
}

func NewFood(name string, ccal uint64, price float64) *Food {
    return &Food{name, ccal, price}
}

func (f *Food) Eaten() {
    fmt.Printf("%s: меня съели !", f.Name)
}

func (f *Food) GetInfo() {
    fmt.Printf("%s содержит %d калорий и стоит %.2f.\n", f.Name, f.Ccal, f.Price)
}

func (f *Food) AddCalories(amount uint64) {
    f.Ccal += amount
    fmt.Printf("%d калорий добавлено в %s.\n", amount, f.Name)
}

func (f *Food) AddPrice(amount float64) {
    f.Price += amount
    fmt.Printf("Цена %s увеличена на %.2f.\n", f.Name, amount)
}

func (f *Food) UpdateName(name string) {
    f.Name = name
    fmt.Printf("Название обновлено до %s.\n", f.Name)
}

func (f *Food) DiscountPrice(percentage float64) {
    f.Price -= f.Price * (percentage / 100)
    fmt.Printf("Цена %s после скидки %.2f%% составляет %.2f.\n", f.Name, percentage, f.Price)
}
