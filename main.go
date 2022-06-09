package main

import (
	"mvc-go/app"
	"mvc-go/db"
)

func main() {
	db.StartDbEngine() //inicializa base de datos
	app.StartRoute()   //Inicializa app
}
