package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type attrs struct {
	HoleLenght int64
	Hp         int64
	Respect    int64
	Weight     float64
}

//	func (a attrs) GetHp() int64 {
//		return a.Hp
//	}
func New(HoleLenght int64, hp int64, respect int64, weight float64) *attrs {
	return &attrs{
		HoleLenght,
		hp,
		respect,
		weight}
}

func (a *attrs) GoodNight() {
	a.HoleLenght -= 2
	a.Hp += 20
	a.Respect -= 2
	a.Weight -= 5
}

func (a *attrs) DoDay(choice string) {
	switch choice {
	case "1":
		a.Dig()
	case "2":
		a.Eat()
	case "3":
		a.Fight()
	case "4":
		a.Sleep()
	}
}
func (a *attrs) Dig() {
	fmt.Println("Your choice is: Копать нору")
	fmt.Println("А как ты хочешь копать нору?")
	fmt.Println("Твои действия:\n 1 - Интенсивно\n 2 - Лениво")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	switch choice {
	case "1":
		a.HoleLenght += 5
		a.Hp -= 30
	case "2":
		a.HoleLenght += 2
		a.Hp -= 10
	default:
		fmt.Println("Не тупи, в наказание ты идёшь в кровать")
	}
}

func (a *attrs) Eat() {
	fmt.Println("Your choice is: Поесть травки")
	fmt.Println("А какую травку ты хочешь испробовать?")
	fmt.Println("Твой выбор:\n 1 - Жухлой\n 2 - Зелёной")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	switch choice {
	case "1":
		a.Weight += 15
		a.Hp += 10
	case "2":
		if a.Respect < 30 {
			a.Hp -= 30
		} else if a.Respect >= 30 {
			a.Hp += 30
			a.Weight += 30
		} else {
			fmt.Println("Не тупи, в наказание ты идёшь в кровать")
		}
	default:
		fmt.Println("Не тупи, в наказание ты идёшь в кровать")
	}
}

func Roll(a *attrs, EnemyWeight float64) float64 {
	propability := a.Weight / (EnemyWeight + a.Weight)
	return propability
}

func Defeat(a *attrs, EnemyWeight float64) {
	fmt.Println("Боже, ты просто нулёвый чел")
	fmt.Println("К твоему сожалению, ты проиграл")
	if a.Weight < EnemyWeight {
		a.Hp -= 50
	} else if a.Weight == EnemyWeight {
		a.Hp -= 35
	} else if a.Weight > EnemyWeight {
		a.Hp -= 20

	}

}
func Pobeda(a *attrs, EnemyWeight float64) {
	fmt.Println("Вау, какой ты крутой и классный")
	fmt.Println("Ты победил, вп!")

	if a.Weight < EnemyWeight {
		a.Respect += 40
	} else if a.Weight == EnemyWeight {
		a.Respect += 20
	} else if a.Weight > EnemyWeight {
		a.Respect += 5
	}

}

func (a *attrs) Fight() {
	easy := 30.0
	medium := 50.0
	hard := 70.0
	result := rand.Float64()
	fmt.Println("Your choice is: Пойти махаться")
	fmt.Println("С кем пойдёшь драться?")
	fmt.Println("Соперник:\n 1 - Лёгенький чел\n 2 - Среднячок\n 3 - На массе")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	if choice == "1" {
		fmt.Println("Ты выбрал лёгкую жертву")
		propability := Roll(a, easy)
		if result <= propability {
			Defeat(a, easy)
		} else if result > propability {
			Pobeda(a, easy)
		}
	} else if choice == "2" {
		fmt.Println("Ты выбрал среднячковую жертву")
		propability := Roll(a, easy)
		if result <= propability*100 {
			Defeat(a, medium)
		} else if result > propability*100 {
			Pobeda(a, medium)
		}
	} else if choice == "3" {
		fmt.Println("Ты выбрал массивного паренька")
		propability := Roll(a, easy)
		if result <= propability*100 {
			Defeat(a, hard)
		} else if result > propability*100 {
			Pobeda(a, hard)
		}
	} else {
		fmt.Println("Не тупи, в наказание ты идёшь в кровать")
	}

}

func (a *attrs) Sleep() {
	a.HoleLenght -= 2
	a.Hp += 20
	a.Respect -= 2
	a.Weight -= 5
}
