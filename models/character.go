package models //✍ Меняем объявление пакета на models

// Character представляет персонажа в игре
type Character struct {
	name   string //✍ Приватное поле - доступно только внутри пакета models
	health int    //✍ Приватное поле - доступно только внутри пакета models
	damage int    //✍ Приватное поле - доступно только внутри пакета models
}

func NewCharacter(name string, health, damage int) *Character {
	if health <= 0 {
		health = 1
	}
	if damage < 0 {
		damage = 0
	}
	if name == "" {
		name = "Неизвестный"
	}

	return &Character{
		name:   name,
		health: health,
		damage: damage,
	}
}

//✅ GetName возвращает имя персонажа
func (c *Character) GetName() string {
	return c.name
}

//✅ GetHealth возвращает текущее здоровье персонажа
func (c *Character) GetHealth() int {
	return c.health
}

//✅ GetDamage возвращает урон персонажа
func (c *Character) GetDamage() int {
	return c.damage
}

func (c *Character) TakeDamage(value int) {
	if value < 0 {
		return
	}

	c.health -= value
	if c.health < 0 {
		c.health = 0
	}
}

func (c *Character) IsDied() bool {
	return c.health <= 0
}
