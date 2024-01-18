package controller

import (
	inittemplate "Golanta/templates"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var id = 1

type Adventurer struct {
	CharId  int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Level   int    `json:"level"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

var adventurers []Adventurer

const Port = "localhost:8080"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	LoadDataFromJSON()

	inittemplate.Temp.ExecuteTemplate(w, "index", adventurers)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		charid := id
		id += 1
		name := r.FormValue("name")
		class := r.FormValue("class")
		adventurers = append(adventurers, Adventurer{CharId: charid, Name: name, Class: class})
		SaveDataToJSON()
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	inittemplate.Temp.ExecuteTemplate(w, "create", nil)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	adventurer, err := FindAdventurerByName(name)
	if err != nil {
		http.Error(w, "Aventurier non trouvé", http.StatusNotFound)
		return
	}

	// Affiche le profil de l'aventurier
	inittemplate.Temp.ExecuteTemplate(w, "profile", adventurer)
}

func SaveDataToJSON() {
	data, err := json.MarshalIndent(adventurers, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde des données JSON:", err)
		return
	}

	err = ioutil.WriteFile("characters.json", data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON:", err)
	}
}

func LoadDataFromJSON() {
	file, err := os.Open("characters.json")
	if err != nil {
		fmt.Println("Aucun fichier JSON existant. Création d'un nouveau.")
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON:", err)
		return
	}

	err = json.Unmarshal(data, &adventurers)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}
}

func AdventurerHandler(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	name := segments[len(segments)-1]
	adventurer, err := FindAdventurerByName(name)
	if err != nil {
		http.Error(w, "Aventurier non trouvé", http.StatusNotFound)
		return
	}

	inittemplate.Temp.ExecuteTemplate(w, "adventurer", adventurer)
}

func FindAdventurerByName(name string) (*Adventurer, error) {
	LoadDataFromJSON()

	for _, adv := range adventurers {
		if adv.Name == name {
			return &adv, nil
		}
	}
	return nil, fmt.Errorf("Aventurier avec le nom %s non trouvé", name)
}
