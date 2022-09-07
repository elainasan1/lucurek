package attacks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Methods sk33d
var Methods map[string]*Attack = make(map[string]*Attack)

type toyota struct {
	Toyota []audi `json:"methods"`
}

type audi struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          uint16 `json:"ID"`
	Vip         bool   `json:"vip"`
	DefPort     string `json:"DefaultPort"`
	DefDuration     string `json:"DefaultDuration"`
	APIShit struct {
		API         string `json:"Apilink"`
	} `json:"Links"`
}

//Configure sk33d
func Configure() {
	jsonFile, err := os.Open("Methods.json")
	if err != nil {
		fmt.Print(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var toyotas toyota

	json.Unmarshal(byteValue, &toyotas)

	for i := 0; i < len(toyotas.Toyota); i++ {
		log.Println(" [Method Found] | Name \"" + toyotas.Toyota[i].Name + "\"")
		Methods[toyotas.Toyota[i].Name] = &Attack{
			ID:          toyotas.Toyota[i].ID,
			Name:        toyotas.Toyota[i].Name,
			Description: toyotas.Toyota[i].Description,
			Vip:         toyotas.Toyota[i].Vip,
			Port:        toyotas.Toyota[i].DefPort,
			Duration:        toyotas.Toyota[i].DefDuration,
			API:         toyotas.Toyota[i].APIShit.API,
		}
	}
}

//Get returns a method
func Get(name string) *Attack {
	return Methods[name]
}

//Flush The Methods (needed For Reload)
func Flush() *Attack {
	for _, k := range Methods {
		delete(Methods, k.Name)
	}
	return nil
}
