package main

type Character struct {
	Name   string
	Health int
	Damage int
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
		Name:   name,
		Health: health,
		Damage: damage,
	}
}

func (c *Character) TakeDamage(value int) {
	if value < 0 {
		return
	}

	c.Health -= value
	if c.Health < 0 {
		c.Health = 0
	}
}

func (c *Character) IsDied() bool {
	return c.Health <= 0
}
