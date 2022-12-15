package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/RyabovNick/databasecourse_p2/tree/master/golang/tasks/console_game/functions"
)

func main() {
	hero := all_functions.New(10, 100, 20, 30.0)

	fmt.Print("Твои параметры:\n Длина норы = ", hero.HoleLenght, "\n Хп = ", hero.Hp, "\n Уважение = ", hero.Respect, "\n Вес = ", hero.Weight, "\n")

	// easy := 30.0
	// medium := 50.0
	// hard := 70.0

	for {
		fmt.Println("\nТвои действия\n 1 - Копать нору\n 2 - Поесть травки\n 3 - Пойти махаться\n 4 - Пойти баеньки\n ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()
		hero.DoDay(choice)
		hero.GoodNight()
		fmt.Print("\nТвои параметры:\n Длина норы = ", hero.HoleLenght, "\n Здоровье = ", hero.Hp, "\n Уважение = ", hero.Respect, "\n Вес = ", hero.Weight, "\n")
		if hero.HoleLenght <= 0 || hero.Hp <= 0 || hero.Weight <= 0 || hero.Respect <= 0 {
			fmt.Println("\n YOU DIED")
			break
		} else if hero.Respect > 100 {
			fmt.Println("\n WINNER WINNER, CHICKEN DINNER")
			break
		} else if hero.Weight >= 150 {
			fmt.Println("\n YOU ARE TOO FAT")
			break
		}
	}
}
