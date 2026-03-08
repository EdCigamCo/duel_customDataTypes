package main

//✅ Actor представляет персонажа пользователя
type Actor struct {
	*Character     //✅ Встраивание - Actor получает все поля и методы Character
	MaxHealth  int //✅ Дополнительное поле только для Actor - максимальное здоровье
	Armor      int //✅ Дополнительное поле только для Actor - броня
}

//✅ NewActor создает нового пользователя с заданными характеристиками
func NewActor(name string, health, damage, armor int) *Actor {
	//✅ Валидация входных данных
	if name == "" {
		name = "Игрок"
	}
	if armor < 0 {
		armor = 0
	}

	//✅ Создаем Actor, инициализируя встроенную структуру Character через конструктор NewCharacter
	actor := &Actor{
		Character: NewCharacter(name, health, damage), //✅ Инициализация встроенной структуры
		Armor:     armor,
	}

	//✅ Устанавливаем MaxHealth равным начальному здоровью
	actor.MaxHealth = actor.Health
	return actor
}

//✅ CalculateReductionDamage применяет защиту брони к урону и возвращает финальный урон
func (a *Actor) CalculateReductionDamage(damage int) int {
	if damage < 0 {
		return 0
	}

	//✅ Вычисляем урон с учетом брони: урон минус броня
	reductionDamage := damage - a.Armor
	if reductionDamage < 0 {
		reductionDamage = 0 //✅ Урон не может быть отрицательным
	}

	return reductionDamage
}

//✅ Heal восстанавливает здоровье пользователя (10% от максимального здоровья)
// Возвращает количество восстановленного здоровья и успешно ли восстановлено
func (a *Actor) Heal() (int, bool) {
	//✅ Проверяем, нужно ли восстанавливать здоровье
	if a.Health >= a.MaxHealth {
		return 0, false //✅ Здоровье уже максимальное
	}

	//✅ Вычисляем количество восстановленного здоровья (10% от MaxHealth)
	healAmount := a.MaxHealth / 10
	a.Health += healAmount
	if a.Health > a.MaxHealth {
		a.Health = a.MaxHealth //✅ Здоровье не может превышать максимальное
	}

	return healAmount, true
}

// TakeDamage переопределяет метод получения урона с учетом брони
func (a *Actor) TakeDamage(value int) {
	if value < 0 {
		return
	}

	//✅ Вычисляем урон с учетом брони
	reductionDamage := a.CalculateReductionDamage(value)
	//✅ Вызываем метод встроенной структуры для применения урона
	a.Character.TakeDamage(reductionDamage)

	//✅ Броня уменьшается на половину атаки врага
	if a.Armor > 0 {
		a.Armor -= value / 2
		if a.Armor < 0 {
			a.Armor = 0 //✅ Броня не может быть отрицательной
		}
	}
}
