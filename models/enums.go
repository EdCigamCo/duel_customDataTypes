package models //✍ Меняем объявление пакета на models

type BattleResultEnum = string

const (
	BattleResultUnknown BattleResultEnum = "Unknown"
	BattleResultWin     BattleResultEnum = "Победа"
	BattleResultLose    BattleResultEnum = "Поражение"
	BattleResultDraw    BattleResultEnum = "Ничья"
)
