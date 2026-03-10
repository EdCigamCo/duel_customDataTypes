package models //✍ Меняем объявление пакета на models

// Actor представляет персонажа пользователя
type Actor struct {
	*Character
	maxHealth int //✍ Приватное поле - доступно только внутри пакета models
	armor     int //✍ Приватное поле - доступно только внутри пакета models
}

func NewActor(name string, health, damage, armor int) *Actor {
	if name == "" {
		name = "Игрок"
	}
	if armor < 0 {
		armor = 0
	}

	actor := &Actor{
		Character: NewCharacter(name, health, damage),
		armor:     armor,
	}

	actor.maxHealth = actor.health
	return actor
}

func (a *Actor) GetArmor() int {
	return a.armor
}

func (a *Actor) CalculateReductionDamage(damage int) int {
	if damage < 0 {
		return 0
	}

	reductionDamage := damage - a.armor
	if reductionDamage < 0 {
		reductionDamage = 0
	}

	return reductionDamage
}

func (a *Actor) Heal() (int, bool) {
	if a.health >= a.maxHealth {
		return 0, false
	}

	healAmount := a.maxHealth / 10
	a.health += healAmount
	if a.health > a.maxHealth {
		a.health = a.maxHealth
	}

	return healAmount, true
}

func (a *Actor) TakeDamage(value int) {
	if value < 0 {
		return
	}

	reductionDamage := a.CalculateReductionDamage(value)
	a.Character.TakeDamage(reductionDamage)

	if a.armor > 0 {
		a.armor -= value / 2
		if a.armor < 0 {
			a.armor = 0
		}
	}
}
