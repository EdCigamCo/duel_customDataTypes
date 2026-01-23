package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type GameService struct {
	Enemies []*Character
	Scanner *bufio.Scanner
}

func NewGameService() *GameService {
	enemies := []*Character{
		NewCharacter("Рыцарь", 330, 10),
		NewCharacter("Боец", 200, 18),
		NewCharacter("Убийца", 150, 25),
	}

	return &GameService{
		Enemies: enemies,
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

func (g *GameService) Run() {
	//✅ Случайный выбор противника
	enemyID := rand.Intn(len(g.Enemies))
	enemy := g.Enemies[enemyID]

	//✅ Создаем пользователя через конструктор
	actor := NewCharacter("Лабубу", 230, 16)

	fmt.Printf("\n%s, добро пожаловать в игру Duel!\n", actor.Name)

	fmt.Printf("\n%s, твой соперник: %s\n\n", actor.Name, enemy.Name)

	var isPlaying bool = true
	for isPlaying {
		//✍ Заменяем прямой вывод на вызов метода
		g.printStatus(actor, enemy)

		//✍ Заменяем прямой ввод на вызов метода
		action := g.getUserAction()
		if !g.processAction(action, actor, enemy) {
			continue
		}

		fmt.Println()

		//✍ Заменяем прямую проверку на вызов метода
		battleResult := g.checkBattleResult(actor, enemy)
		switch battleResult {
		case "Победа":
			fmt.Printf("%s! Ты победил %s\n\n", actor.Name, enemy.Name)
			g.Enemies[enemyID], g.Enemies[len(g.Enemies)-1] = g.Enemies[len(g.Enemies)-1], g.Enemies[enemyID]
			g.Enemies = g.Enemies[:len(g.Enemies)-1]

			if len(g.Enemies) == 0 {
				fmt.Printf("%s, поздравляем! Ты победил всех врагов!\n\n", actor.Name)
				isPlaying = false
			} else {
				enemyID = rand.Intn(len(g.Enemies))
				enemy = g.Enemies[enemyID]
				fmt.Printf("%s, твой следующий соперник: %s\n\n", actor.Name, enemy.Name)
			}
		case "Поражение":
			fmt.Printf("%s, к сожалению, ты проиграл :(\n\n", actor.Name)
			isPlaying = false
		case "Ничья":
			fmt.Println("Ничья")
			isPlaying = false
		}
	}

	fmt.Printf("*нажать клавишу Enter для выхода")
	g.Scanner.Scan()
}

func (g *GameService) printStatus(actor *Character, enemy *Character) {
	fmt.Printf("Твое здоровье: %d, твоя атака: %d\n", actor.Health, actor.Damage)
	fmt.Printf("Здоровье противника: %d, атака противника: %d\n", enemy.Health, enemy.Damage)
}

func (g *GameService) getUserAction() string {
	fmt.Print("\nВыберите действие:\n1 - атаковать\n")
	g.Scanner.Scan()
	return g.Scanner.Text()
}

func (g *GameService) processAction(action string, actor *Character, enemy *Character) bool {
	switch action {
	case "1":
		enemyDamage := actor.Damage
		enemy.TakeDamage(enemyDamage)
		fmt.Printf("Ты нанес %d урона противнику\n", enemyDamage)

		if !enemy.IsDied() {
			enemyAttack := enemy.Damage
			actor.TakeDamage(enemyAttack)
			fmt.Printf("Тебе нанесено %d урона\n", enemyAttack)
		}
	default:
		fmt.Println("Неверный ввод. Выберите 1 (атаковать).")
		return false
	}

	return true
}

func (g *GameService) checkBattleResult(actor *Character, enemy *Character) string {
	actorDied := actor.IsDied()
	enemyDied := enemy.IsDied()

	if actorDied && enemyDied {
		return "Ничья"
	}
	if enemyDied {
		return "Победа"
	}
	if actorDied {
		return "Поражение"
	}

	return "Unknown"
}
