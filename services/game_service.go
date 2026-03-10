package services //✍ Меняем объявление пакета на services

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"duel/config" //✅ Импорт пакета config
	"duel/models" //✅ Импорт пакета models
)

// GameService представляет игровую сессию
type GameService struct {
	enemies []*models.Character //✍ Приватное поле - доступно только внутри пакета services
	scanner *bufio.Scanner      //✍ Приватное поле - доступно только внутри пакета services
}

// NewGameService создает новую игровую сессию
func NewGameService() *GameService {
	//✍ Используем константы из пакета config
	enemies := []*models.Character{
		models.NewCharacter(config.KnightName, config.KnightHealth, config.KnightAttack),
		models.NewCharacter(config.FighterName, config.FighterHealth, config.FighterAttack),
		models.NewCharacter(config.AssassinName, config.AssassinHealth, config.AssassinAttack),
	}

	//✍ Заполняем приватные поля
	return &GameService{
		enemies: enemies,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (g *GameService) Run() {
	//✍ Обращение к приватному полю g.enemies (вместо g.Enemies)
	enemyID := rand.Intn(len(g.enemies))
	enemy := g.enemies[enemyID]

	//✍ Используем константы из пакета config
	actor := models.NewActor(
		config.ActorName,
		config.ActorHealth,
		config.ActorAttack,
		config.ActorArmor,
	)

	//✍ Getter-методы вместо прямого доступа к полям (actor.Name → actor.GetName(), enemy.Name → enemy.GetName())
	fmt.Printf("\n%s, добро пожаловать в игру Duel!\n", actor.GetName())

	fmt.Printf("\n%s, твой соперник: %s\n\n", actor.GetName(), enemy.GetName())

	var isPlaying bool = true
	for isPlaying {
		g.printStatus(actor, enemy)

		action := g.getUserAction()
		if !g.processAction(action, actor, enemy) {
			continue
		}

		fmt.Println()

		// Проверка окончания игры
		battleResult := g.checkBattleResult(actor, enemy)

		//✍ Для switch/case используем константы из пакета models
		switch battleResult {
		case models.BattleResultWin:
			//✍ Используем getter-методы
			fmt.Printf("%s! Ты победил %s\n\n", actor.GetName(), enemy.GetName())

			//✍ Приватное поле g.enemies (вместо g.Enemies)
			g.enemies[enemyID], g.enemies[len(g.enemies)-1] = g.enemies[len(g.enemies)-1], g.enemies[enemyID]
			g.enemies = g.enemies[:len(g.enemies)-1]

			if len(g.enemies) == 0 {
				//✍ Используем getter-метод
				fmt.Printf("%s, поздравляем! Ты победил всех врагов!\n\n", actor.GetName())
				isPlaying = false
			} else {
				enemyID = rand.Intn(len(g.enemies))
				enemy = g.enemies[enemyID]

				//✍ Используем getter-методы
				fmt.Printf("%s, твой следующий соперник: %s\n\n", actor.GetName(), enemy.GetName())
			}
		case models.BattleResultLose:
			//✍ Используем getter-метод
			fmt.Printf("%s, к сожалению, ты проиграл :(\n\n", actor.GetName())
			isPlaying = false
		case models.BattleResultDraw:
			fmt.Println("Ничья")
			isPlaying = false
		}
	}

	fmt.Printf("*нажать клавишу Enter для выхода")
	g.scanner.Scan()
}

// ✍ Корректируем параметры
func (g *GameService) printStatus(actor *models.Actor, enemy *models.Character) {
	//✍ Используем getter-методы вместо прямого доступа к полям
	fmt.Printf("Твое здоровье: %d, твоя атака: %d, твоя броня: %d\n",
		actor.GetHealth(), actor.GetDamage(), actor.GetArmor())
	fmt.Printf("Здоровье противника: %d, атака противника: %d\n",
		enemy.GetHealth(), enemy.GetDamage())
}

func (g *GameService) getUserAction() string {
	fmt.Print("\nВыберите действие:\n1 - увеличить здоровье\n2 - атаковать\n")

	//✍ Приватное поле g.scanner (вместо g.Scanner)
	g.scanner.Scan()
	return g.scanner.Text()
}

// processAction обрабатывает выбранное действие
func (g *GameService) processAction(action string, actor *models.Actor, enemy *models.Character) bool {
	switch action {
	case "1":
		restoreHealth, wasRestored := actor.Heal()
		if wasRestored {
			fmt.Printf("Ты восстановил %d здоровья\n", restoreHealth)
		} else {
			fmt.Println("Невозможно восстановить здоровье")
		}

		//✍ Используем getter-метод для получения урона врага
		enemyAttack := enemy.GetDamage()
		reductionDamage := actor.CalculateReductionDamage(enemyAttack)

		actor.TakeDamage(enemyAttack)
		fmt.Printf("Тебе нанесено %d урона, броня защитила, здоровье уменьшилось на %d\n", enemyAttack, reductionDamage)
	case "2":
		//✍ Используем getter-метод для получения урона игрока
		enemyDamage := actor.GetDamage()
		enemy.TakeDamage(enemyDamage)
		fmt.Printf("Ты нанес %d урона противнику\n", enemyDamage)

		// Враг атакует в ответ, если не умер
		if !enemy.IsDied() {
			//✍ Используем getter-метод для получения урона врага
			enemyAttack := enemy.GetDamage()
			reductionDamage := actor.CalculateReductionDamage(enemyAttack)

			actor.TakeDamage(enemyAttack)
			fmt.Printf("Тебе нанесено %d урона, броня защитила, здоровье уменьшилось на %d\n", enemyAttack, reductionDamage)
		}
	default:
		fmt.Println("Неверный ввод. Выберите 1 или 2.")
		return false
	}

	return true
}

// ✍ Изменяем тип параметра actor с *Character на *Actor
func (g *GameService) checkBattleResult(actor *models.Actor, enemy *models.Character) models.BattleResultEnum {
	actorDied := actor.IsDied()
	enemyDied := enemy.IsDied()

	if actorDied && enemyDied {
		return models.BattleResultDraw
	}
	if enemyDied {
		return models.BattleResultWin
	}
	if actorDied {
		return models.BattleResultLose
	}

	return models.BattleResultUnknown
}
