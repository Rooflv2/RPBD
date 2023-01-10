package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type stats struct {
	HoleLenght int64
	Hp         int64
	Rep        int64
	Weight     float64
}

func New(HoleLenght int64, hp int64, rep int64, weight float64) *stats {
	return &stats{HoleLenght, hp, rep, weight}
}

func (a *stats) Night() {
	a.HoleLenght -= 2
	a.Hp += 20
	a.Rep -= 2
	a.Weight -= 5
}

func (a *stats) Day(choice string) {
	switch choice {
	case "1":
		a.Dig()
	case "2":
		a.Eat()
	case "3":
		a.Fight()
	case "4":
		a.Night()
	}
}

func (a *stats) Dig() {
	fmt.Println("Ты выбрал копать нору, но не выбрал как её копать \n Есть два варианта: \n1 - Интенсивно \n2 - Лениво")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	choice := scan.Text()
	switch choice {
	case "1":
		a.HoleLenght += 5
		a.Hp -= 30
	case "2":
		a.HoleLenght += 2
		a.Hp -= 10
	default:
		fmt.Println("Ахтунг! Такого выбора нет, проверь что ты ввёл")
	}
}

func (a *stats) Eat() {
	fmt.Println("Ты выбрал пожевать травки, но не выбрал какую \n Есть два варианта: \n1 - Жухлую \n2- Зелёную")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	choice := scan.Text()
	switch choice {
	case "1":
		a.Hp += 10
		a.Weight += 15
	case "2":
		if a.Rep < 30 {
			a.Hp -= 30
		} else {
			a.Hp += 30
			a.Weight += 30
		}
	default:
		fmt.Println("Ахтунг! Такого выбора нет, проверь что ты ввёл")
	}
}

func Generate(a *stats, EnemyWeight float64) float64 {
	chance := a.Weight / (EnemyWeight + a.Weight)
	return chance
}

func Lose(a *stats, EnemyWeight float64) {
	fmt.Println("Увы, но ты проиграл :(")
	if a.Weight < EnemyWeight {
		a.Hp -= 50
	} else if a.Weight == EnemyWeight {
		a.Hp -= 35
	} else if a.Weight > EnemyWeight {
		a.Hp -= 20
	}
}

func Win(a *stats, EnemyWeight float64) {
	fmt.Println("Ого, кто это тут? Правильно, наш победитель! :D")
	if a.Weight < EnemyWeight {
		a.Rep += 40
	} else if a.Weight == EnemyWeight {
		a.Rep += 20
	} else if a.Weight > EnemyWeight {
		a.Rep += 5
	}
}

func (a *stats) Fight() {
	eazy := 30.0
	medium := 50.0
	hard := 70.0
	result := rand.Float64()
	fmt.Println("Ты выбрал пойти подраться, но не выбрал противника. Варианты махыча: \n 1 - Легко \n 2 - Придётся потрудиться \n 3 - Нереально сложно")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	choice := scan.Text()
	switch choice {
	case "1":
		fmt.Println("Ты выбрал лёгкий бой \n Бой начинается...")
		chance := Generate(a, eazy)
		if result <= chance {
			Lose(a, eazy)
		} else if result > chance {
			Win(a, eazy)
		}
	case "2":
		fmt.Println("О, ты добрался и выбрал средний бой \n Бой начинается...")
		chance := Generate(a, medium)
		if result <= chance {
			Lose(a, medium)
		} else if result > chance {
			Win(a, medium)
		}
	case "3":
		fmt.Println("Ого, ты добрался до сложного боя! Или ты просто спидранер -_-? В любом случае уже поздно :) \n Бой начинается...")
		chance := Generate(a, hard)
		if result <= chance {
			Lose(a, hard)
		} else if result > chance {
			Win(a, hard)
		}
	default:
		fmt.Println("Ахтунг! Такого выбора нет, проверь что ты ввёл")
	}
}
