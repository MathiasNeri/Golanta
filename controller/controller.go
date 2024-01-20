package controller

import (
	inittemplate "Golanta/templates"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Adventurer struct {
	CharId     int    `json:"id"`
	Name       string `json:"name"`
	Equipe     string `json:"equipe"`
	Lvl_survie string `json:"level"`
	HP         int    `json:"hp"`
	Res        int    `json:"res"`
}

var adventurers []Adventurer

const Port = "localhost:8080"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	LoadDataFromJSON()
	inittemplate.Temp.ExecuteTemplate(w, "index", adventurers)
}

func AdventurerHandler(w http.ResponseWriter, r *http.Request) {
	//Charger le JSO
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	LoadDataFromJSON()
	var selectedAdventurer Adventurer
	for _, adventurer := range adventurers {
		fmt.Println(adventurer)
		if adventurer.CharId == id {
			selectedAdventurer = adventurer
		}
	}
	inittemplate.Temp.ExecuteTemplate(w, "adventurer", selectedAdventurer)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LoadDataFromJSON()

		name := r.FormValue("name")
		equipe := r.FormValue("equipe")
		level := r.FormValue("level")
		hp, errHp := strconv.Atoi(r.FormValue("hp"))
		res, errRes := strconv.Atoi(r.FormValue("res"))

		if errHp != nil || errRes != nil {
			fmt.Println("Error parsing HP or Res from form:", errHp, errRes)
			http.Error(w, "Invalid HP or Res values", http.StatusBadRequest)
			return
		}

		if hp+res > 10 {
			fmt.Println("La somme des HP et Res ne peut pas dépasser 10")
			http.Error(w, "La somme des HP et Res ne peut pas dépasser 10", http.StatusBadRequest)
			return
		}

		charid := getNextCharID()

		adventurers = append(adventurers, Adventurer{CharId: charid, Name: name, Equipe: equipe, Lvl_survie: level, HP: hp, Res: res})
		SaveDataToJSON()
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	inittemplate.Temp.ExecuteTemplate(w, "create", nil)
}

func SaveDataToJSON() {
	data, err := json.MarshalIndent(adventurers, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde des données JSON:", err)
		return
	}

	err = os.WriteFile("characters.json", data, 0644)
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

func FindAdventurerByID(id int) (*Adventurer, error) {
	LoadDataFromJSON()

	for _, adv := range adventurers {
		if adv.CharId == id {
			return &adv, nil
		}
	}
	return nil, fmt.Errorf("aventurier avec l'ID %d non trouvé", id)
}

func getNextCharID() int {
	highestID := 0

	// Loop through the adventurers to find the highest CharId
	for _, adv := range adventurers {
		if adv.CharId > highestID {
			highestID = adv.CharId
		}
	}

	// Increment the highest ID by 1 to get the next available ID
	return highestID + 1
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LoadDataFromJSON()

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid character ID", http.StatusBadRequest)
			return
		}
		for i, adventurer := range adventurers {
			if adventurer.CharId == id {
				adventurers[i].Name = r.FormValue("name")
				adventurers[i].Equipe = r.FormValue("equipe")
				adventurers[i].Lvl_survie = r.FormValue("level")
				adventurers[i].HP, err = strconv.Atoi(r.FormValue("hp"))
				adventurers[i].Res, err = strconv.Atoi(r.FormValue("res"))
				if err != nil {
					http.Error(w, "Invalid HP or Res values", http.StatusBadRequest)
					return
				}
				if adventurers[i].HP+adventurers[i].Res > 10 {
					http.Error(w, "La somme des HP et Res ne peut pas dépasser 10", http.StatusBadRequest)
					return
				}

				SaveDataToJSON()
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
		http.Error(w, "Aventurier non trouvé", http.StatusNotFound)
	} else {
		http.Error(w, "Mauvaise méthode HTTP", http.StatusMethodNotAllowed)
	}
}
