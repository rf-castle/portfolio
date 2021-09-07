package models

type Page struct {
	Name string `json:"name"`
	Root string `json:"root"`
	Icon string `json:"icon"`
}

var SecretPage = []Page{}
