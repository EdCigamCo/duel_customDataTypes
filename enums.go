package main

// ✅ Создаем type alias для string - это будет наш enum-подобный тип
type BattleResultEnum = string

// ✅ Объявляем константы для возможных результатов боя
const (
	BattleResultUnknown BattleResultEnum = "Unknown"
	BattleResultWin     BattleResultEnum = "Победа"
	BattleResultLose    BattleResultEnum = "Поражение"
	BattleResultDraw    BattleResultEnum = "Ничья"
)
