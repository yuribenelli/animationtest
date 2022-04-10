package main

import (
	v1 "animationtest/internal/api/v1"
	"animationtest/internal/configuration"
	"net/http"
)

func main() {
	//Iniziallizzo DB
	configuration.InitConfig()
	router := v1.RouterInit()
	http.ListenAndServe("5800", router)

}
