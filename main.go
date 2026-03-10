package main

import "duel/services" //✅ Добавляем импорт пакета services

func main() {
	//✍ Вызываем конструктор структуры GameService из пакета services
	game := services.NewGameService()
	game.Run()
}
