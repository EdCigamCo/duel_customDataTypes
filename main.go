package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	actorName := "Лабубу"
	actorHealth := 230
	actorDamage := 16

	enemyNames := []string{"Рыцарь", "Боец", "Убийца"}
	enemyHealths := []int{330, 200, 150}
	enemyDamages := []int{10, 18, 25}

	enemyID := rand.Intn(len(enemyNames))
	enemyName := enemyNames[enemyID]
	enemyHealth := enemyHealths[enemyID]
	enemyDamage := enemyDamages[enemyID]

	fmt.Printf("\n%s, добро пожаловать в игру Duel!\n", actorName)

	fmt.Printf("\n%s, твой соперник: %s\n\n", actorName, enemyName)

	scanner := bufio.NewScanner(os.Stdin)

	var isPlaying bool = true
	for isPlaying {
		fmt.Printf("Твое здоровье: %d, твоя атака: %d\n", actorHealth, actorDamage)
		fmt.Printf("Здоровье противника: %d, атака противника: %d\n", enemyHealth, enemyDamage)

		fmt.Print("\nВыберите действие:\n1 - атаковать\n")
		scanner.Scan()
		action := scanner.Text()

		switch action {
		case "1":
			enemyHealth -= actorDamage
			if enemyHealth < 0 {
				enemyHealth = 0
			}
			fmt.Printf("Ты нанес %d урона противнику\n", actorDamage)

			if enemyHealth > 0 {
				actorHealth -= enemyDamage
				if actorHealth < 0 {
					actorHealth = 0
				}
				fmt.Printf("Тебе нанесено %d урона\n", enemyDamage)
			}
		default:
			fmt.Println("Неверный ввод. Выберите 1 (атаковать).")
			continue
		}

		fmt.Println()

		actorDied := actorHealth <= 0
		enemyDied := enemyHealth <= 0

		if actorDied && enemyDied {
			fmt.Println("Ничья")
			isPlaying = false
		} else if enemyDied {
			fmt.Printf("Победа! Ты победил %s\n\n", enemyName)
			enemyNames[enemyID], enemyNames[len(enemyNames)-1] = enemyNames[len(enemyNames)-1], enemyNames[enemyID]
			enemyHealths[enemyID], enemyHealths[len(enemyHealths)-1] = enemyHealths[len(enemyHealths)-1], enemyHealths[enemyID]
			enemyDamages[enemyID], enemyDamages[len(enemyDamages)-1] = enemyDamages[len(enemyDamages)-1], enemyDamages[enemyID]

			enemyNames = enemyNames[:len(enemyNames)-1]
			enemyHealths = enemyHealths[:len(enemyHealths)-1]
			enemyDamages = enemyDamages[:len(enemyDamages)-1]

			if len(enemyNames) == 0 {
				fmt.Printf("%s, поздравляем! Ты победил всех врагов!\n\n", actorName)
				isPlaying = false
			} else {
				enemyID = rand.Intn(len(enemyNames))
				enemyName = enemyNames[enemyID]
				enemyHealth = enemyHealths[enemyID]
				enemyDamage = enemyDamages[enemyID]
				fmt.Printf("%s, твой следующий соперник: %s\n\n", actorName, enemyName)
			}
		} else if actorDied {
			fmt.Printf("%s, к сожалению, ты проиграл :(\n\n", actorName)
			isPlaying = false
		}
	}

	fmt.Printf("*нажать клавишу Enter для выхода")
	scanner.Scan()
}
