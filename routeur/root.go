package routeur

import (
	"Golanta/controller"
	"fmt"
	"log"
	"net/http"
)

func InitServe() {

	FileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/adventurer", controller.AdventurerHandler)
	http.HandleFunc("/create", controller.CreateHandler)
	http.HandleFunc("/update", controller.UpdateHandler)

	if err := http.ListenAndServe(controller.Port, nil); err != nil {

		fmt.Printf("ERREUR LORS DE L'INITIATION DES ROUTES %v \n", err)

		log.Fatal(err)

	}
}
