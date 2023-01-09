package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hero := functions.New(10, 100, 20, 30)

	fmt.Print("\n Твои параметры:\n Длина норы = ", hero.HoleLenght, "\n Здоровье = ", hero.Hp, "\n Уважение = ", hero.Rep, "\n Вес = ", hero.Weight, "\n")

	for {
		fmt.Println("\n Выбери действие:\n 1 - Копать нору\n 2 - Пожевать траву\n 3 - Подраться\n 4 - Пойти спать\n ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()
		hero.DoDay(choice)
		hero.GoodNight()
		fmt.Print(" Твои новые параметры:\n Длина норы = ", hero.HoleLenght, "\n Здоровье = ", hero.Hp, "\n Уважение = ", hero.Rep, "\n Вес = ", hero.Weight, "\n")
		if hero.HoleLenght <= 0 || hero.Hp <= 0 || hero.Weight <= 0 || hero.Rep <= 0 {
			fmt.Println("\n YOU DIED \n")
			break
		} else if hero.Rep > 100 {
			fmt.Println("\n WINNER WINNER, CHICKEN DINNER \n")
			break
		} else if hero.Weight >= 150 {
			fmt.Println("\n YOU ARE TOO FAT \n")
			break
		}
	}
}
