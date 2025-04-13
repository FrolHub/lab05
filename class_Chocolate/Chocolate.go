package chocolate

import ("fmt"
"example.com/sweet"
"example.com/food")

type Chocolate struct {
    sweet.Sweet
    PCacao float64
}

func NewChocolate(name string, ccal uint64, price float64, sugar uint64, pcacao float64) *Chocolate {
    return &Chocolate{sweet.Sweet{food.Food{name, ccal, price}, sugar}, pcacao}
}

func (c *Chocolate) Eaten() {
    fmt.Printf("%s: меня съели !\n", c.Name)
}

func (c *Chocolate) GetInfo() {
    fmt.Printf("%s содержит %d калорий, %d граммов сахара, %.1f%% какао и стоит %.2f.\n", c.Name, c.Ccal, c.Sugar, c.PCacao, c.Price)
}

func (c *Chocolate) IncreaseCacao(percentage float64) {
    c.PCacao += percentage
    fmt.Printf("Содержание какао в %s увеличено на %.1f%%.\n", c.Name, percentage)
}

func (c *Chocolate) DecreaseCacao(percentage float64) {
    c.PCacao -= percentage
    fmt.Printf("Содержание какао в %s уменьшено на %.1f%%.\n", c.Name, percentage)
}

func (c *Chocolate) UpdateCacao(percentage float64) {
    c.PCacao = percentage
    fmt.Printf("Содержание какао обновлено до %.1f%% в %s.\n", c.PCacao, c.Name)
}

func (c *Chocolate) Break() {
    fmt.Printf("Шоколад: меня поломали на пополам !\n")
}
