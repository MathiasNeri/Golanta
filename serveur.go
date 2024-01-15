package main

import (
	"Golanta/manager"
	"Golanta/routeur"
	initTemplate "Golanta/templates"
	"fmt"
)

func main() {
	manager.PrintColorResult("purple", "server is running...")

	fmt.Println("")
	manager.PrintColorResult("yellow", "CLICK HERE to OPEN  PAGE--->")
	manager.PrintColorResult("blue", " http://localhost:8080/ \n")
	fmt.Println("")
	manager.PrintColorResult("green", "TO STOP THE SERVER , PRESS  'ctrl+C' ")
	initTemplate.InitTemplate()
	routeur.InitServe()
}
