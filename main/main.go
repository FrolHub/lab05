package main

import (
	"fmt"
	"example.com/chocolate"
	"example.com/food"
	"example.com/sweet"
)

func main() {
	eat1 := food.NewFood("ClassFood", 5600, 2.5)
	sweet1 := sweet.NewSweet("ClassSweet", 10, 5.7, 56)
	chocolate1 := chocolate.NewChocolate("ClassChocolate", 230, 3.75, 100, 99.9)

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Показать информацию о еде")
		fmt.Println("2. Показать информацию о сладостях")
		fmt.Println("3. Показать информацию о шоколаде")
		fmt.Println("4. Добавить калорий")
		fmt.Println("5. Добавить цену")
		fmt.Println("6. Обновить название еды")
		fmt.Println("7. Применить скидку к еде")
		fmt.Println("8. Съесть что-то")
		fmt.Println("9. Добавить сахар в сладости")
		fmt.Println("10. Уменьшить сахар в сладостях")
		fmt.Println("11. Обновить сахар в сладостях")
		fmt.Println("12. Уронить сладость")
		fmt.Println("13. Увеличить содержание какао в шоколаде")
		fmt.Println("14. Уменьшить содержание какао в шоколаде")
		fmt.Println("15. Обновить содержание какао в шоколаде")
		fmt.Println("16. Сломать шоколадку пополам")
		fmt.Println("17. Выйти")

		var choice int
		fmt.Print("Введите ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			eat1.GetInfo()
		case 2:
			sweet1.GetInfo()
		case 3:
			chocolate1.GetInfo()
		case 4:
			var amount uint64
			fmt.Print("Введите количество калорий: ")
			fmt.Scan(&amount)
			fmt.Println("1. Добавить калории в еду")
			fmt.Println("2. Добавить калории в сладости")
			fmt.Println("3. Добавить калории в шоколад")
			fmt.Print("Выберите опцию: ")
			var option int
			fmt.Scan(&option)
			switch option {
			case 1:
				eat1.AddCalories(amount)
			case 2:
				sweet1.AddCalories(amount)
			case 3:
				chocolate1.AddCalories(amount)
			default:
				fmt.Println("Неверный выбор.")
			}
		case 5:
			var amount float64
			fmt.Print("Введите количество цены: ")
			fmt.Scan(&amount)
			fmt.Println("1. Добавить цену в еду")
			fmt.Println("2. Добавить цену в сладости")
			fmt.Println("3. Добавить цену в шоколад")
			fmt.Print("Выберите опцию: ")
			var option int
			fmt.Scan(&option)
			switch option {
			case 1:
				eat1.AddPrice(amount)
			case 2:
				sweet1.AddPrice(amount)
			case 3:
				chocolate1.AddPrice(amount)
			default:
				fmt.Println("Неверный выбор.")
			}
		case 6:
			var name string
			fmt.Print("Введите новое название для еды: ")
			fmt.Scan(&name)
			eat1.UpdateName(name)
		case 7:
			var percentage float64
			fmt.Print("Введите процент скидки для еды: ")
			fmt.Scan(&percentage)
			eat1.DiscountPrice(percentage)
		case 8:
			fmt.Println("1.Съесть еду")
			fmt.Println("2.Съесть сладость")
			fmt.Println("3.Съесть шоколад")
			fmt.Print("Выберите опцию: ")
			var option int
			fmt.Scan(&option)
			switch option {
			case 1:
				eat1.Eaten()
			case 2:
				sweet1.Eaten()
			case 3:
				chocolate1.Eaten()
			default:
				fmt.Println("Неверный выбор.")
			}
		case 9:
			var amount uint64
			fmt.Print("Введите количество сахара: ")
			fmt.Scan(&amount)
			sweet1.AddSugar(amount)
		case 10:
			var amount uint64
			fmt.Print("Введите количество сахара: ")
			fmt.Scan(&amount)
			sweet1.DecreaseSugar(amount)
		case 11:
			var amount uint64
			fmt.Print("Введите новое количество сахара: ")
			fmt.Scan(&amount)
			sweet1.UpdateSugar(amount)
		case 12:
			sweet1.Fall()
		case 13:
			var percentage float64
			fmt.Print("Введите процент увеличения какао: ")
			fmt.Scan(&percentage)
			chocolate1.IncreaseCacao(percentage)
		case 14:
			var percentage float64
			fmt.Print("Введите процент уменьшения какао: ")
			fmt.Scan(&percentage)
			chocolate1.DecreaseCacao(percentage)
		case 15:
			var percentage float64
			fmt.Print("Введите новое содержание какао: ")
			fmt.Scan(&percentage)
			chocolate1.UpdateCacao(percentage)
		case 16:
			chocolate1.Break()
		case 17:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
