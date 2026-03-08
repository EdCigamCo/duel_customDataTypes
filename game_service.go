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

	//✍ Создаем пользователя как Actor вместо Character
	actor := NewActor("Лабубу", 230, 16, 25) //✍ Заменяем NewCharacter на NewActor, добавляем параметр брони (25)

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

		// Проверка окончания игры
		battleResult := g.checkBattleResult(actor, enemy)
		switch battleResult {
		case BattleResultWin:
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
		case BattleResultLose:
			fmt.Printf("%s, к сожалению, ты проиграл :(\n\n", actor.Name)
			isPlaying = false
		case BattleResultDraw:
			fmt.Println("Ничья")
			isPlaying = false
		}
	}

	fmt.Printf("*нажать клавишу Enter для выхода")
	g.Scanner.Scan()
}

// ✍ Изменяем тип параметра actor с *Character на *Actor
func (g *GameService) printStatus(actor *Actor, enemy *Character) {
	fmt.Printf("Твое здоровье: %d, твоя атака: %d, твоя броня: %d\n", actor.Health, actor.Damage, actor.Armor) //✍ Добавляем вывод брони
	fmt.Printf("Здоровье противника: %d, атака противника: %d\n", enemy.Health, enemy.Damage)
}

func (g *GameService) getUserAction() string {
	fmt.Print("\nВыберите действие:\n1 - увеличить здоровье\n2 - атаковать\n") // ✍ Добавляем новое действие "1 - увеличить здоровье"
	g.Scanner.Scan()
	return g.Scanner.Text()
}

// processAction обрабатывает выбранное действие
func (g *GameService) processAction(action string, actor *Actor, enemy *Character) bool {
	switch action {
	case "1": //✍ Изменяем обработку действия "1" - теперь это восстановление здоровья
		restoreHealth, wasRestored := actor.Heal() //✅ Вызываем метод Heal() из Actor
		if wasRestored {
			fmt.Printf("Ты восстановил %d здоровья\n", restoreHealth)
		} else {
			fmt.Println("Невозможно восстановить здоровье")
		}

		//✅ Враг атакует в ответ при восстановлении здоровья
		enemyAttack := enemy.Damage
		reductionDamage := actor.CalculateReductionDamage(enemyAttack) //✅ Вычисляем урон с учетом брони

		actor.TakeDamage(enemyAttack) //✅ Применяем урон (метод Actor учитывает броню)
		fmt.Printf("Тебе нанесено %d урона, броня защитила, здоровье уменьшилось на %d\n", enemyAttack, reductionDamage)
	case "2": //✍ Изменяем номер действия атаки с "1" на "2"
		// Пользователь атакует врага
		enemyDamage := actor.Damage
		enemy.TakeDamage(enemyDamage)
		fmt.Printf("Ты нанес %d урона противнику\n", enemyDamage)

		// Враг атакует в ответ, если не умер
		if !enemy.IsDied() {
			enemyAttack := enemy.Damage
			reductionDamage := actor.CalculateReductionDamage(enemyAttack) //✍ Добавляем вычисление урона с учетом брони

			actor.TakeDamage(enemyAttack)                                                                                    //✅ Применяем урон (метод Actor учитывает броню)
			fmt.Printf("Тебе нанесено %d урона, броня защитила, здоровье уменьшилось на %d\n", enemyAttack, reductionDamage) //✍ Обновляем сообщение для отображения информации о броне
		}
	default:
		fmt.Println("Неверный ввод. Выберите 1 или 2.") //✍ Обновляем сообщение об ошибке
		return false
	}

	return true
}

// ✍ Изменяем тип параметра actor с *Character на *Actor
func (g *GameService) checkBattleResult(actor *Actor, enemy *Character) BattleResultEnum {
	actorDied := actor.IsDied()
	enemyDied := enemy.IsDied()

	if actorDied && enemyDied {
		return BattleResultDraw
	}
	if enemyDied {
		return BattleResultWin
	}
	if actorDied {
		return BattleResultLose
	}

	return BattleResultUnknown
}
