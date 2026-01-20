package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	actor := Character{
		Name:   "Лабубу",
		Health: 230,
		Damage: 16,
	}

	game := GameService{
		Enemies: []*Character{
			{Name: "Рыцарь", Health: 330, Damage: 10},
			{Name: "Боец", Health: 200, Damage: 18},
			{Name: "Убийца", Health: 150, Damage: 25},
		},
		Scanner: bufio.NewScanner(os.Stdin),
	}

	enemyID := rand.Intn(len(game.Enemies))
	enemy := game.Enemies[enemyID]

	fmt.Printf("\n%s, добро пожаловать в игру Duel!\n", actor.Name)

	fmt.Printf("\n%s, твой соперник: %s\n\n", actor.Name, enemy.Name)

	var isPlaying bool = true
	for isPlaying {
		fmt.Printf("Твое здоровье: %d, твоя атака: %d\n", actor.Health, actor.Damage)
		fmt.Printf("Здоровье противника: %d, атака противника: %d\n", enemy.Health, enemy.Damage)

		fmt.Print("\nВыберите действие:\n1 - атаковать\n")
		game.Scanner.Scan()
		action := game.Scanner.Text()

		switch action {
		case "1":
			enemy.Health -= actor.Damage
			if enemy.Health < 0 {
				enemy.Health = 0
			}
			fmt.Printf("Ты нанес %d урона противнику\n", actor.Damage)

			if enemy.Health > 0 {
				actor.Health -= enemy.Damage
				if actor.Health < 0 {
					actor.Health = 0
				}
				fmt.Printf("Тебе нанесено %d урона\n", enemy.Damage)
			}
		default:
			fmt.Println("Неверный ввод. Выберите 1 (атаковать).")
			continue
		}

		fmt.Println()

		actorDied := actor.Health <= 0
		enemyDied := enemy.Health <= 0

		if actorDied && enemyDied {
			fmt.Println("Ничья")
			isPlaying = false
		} else if enemyDied {
			fmt.Printf("Победа! Ты победил %s\n\n", enemy.Name)
			game.Enemies[enemyID], game.Enemies[len(game.Enemies)-1] = game.Enemies[len(game.Enemies)-1], game.Enemies[enemyID]

			game.Enemies = game.Enemies[:len(game.Enemies)-1]

			if len(game.Enemies) == 0 {
				fmt.Printf("%s, поздравляем! Ты победил всех врагов!\n\n", actor.Name)
				isPlaying = false
			} else {
				enemyID = rand.Intn(len(game.Enemies))
				enemy = game.Enemies[enemyID]
				fmt.Printf("%s, твой следующий соперник: %s\n\n", actor.Name, enemy.Name)
			}
		} else if actorDied {
			fmt.Printf("%s, к сожалению, ты проиграл :(\n\n", actor.Name)
			isPlaying = false
		}
	}

	fmt.Printf("*нажать клавишу Enter для выхода")
	game.Scanner.Scan()
}
