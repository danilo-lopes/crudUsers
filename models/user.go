package models

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name      string `json:"name"`
	Directory string `json:"directory"`
	Group     string `json:"group"`
	Shell     string `json:"shell"`
}
