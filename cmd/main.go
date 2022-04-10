package main

import (
	v1 "animationtest/internal/api/v1"
	"animationtest/internal/configuration"
	"net/http"
)

func main() {
	//Iniziallizzo DB
	configuration.InitConfig()
	http.ListenAndServe(":80", v1.RouterInit())

}
