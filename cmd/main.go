package main

import (
	"animationtest/internal/character"
	"animationtest/internal/configuration"
)

func main() {
	//Iniziallizzo DB
	configuration.InitConfig()

	router := configuration.RouterInit()
	router.HandleFunc("/pg/{ID}", character.GetOne)

}
