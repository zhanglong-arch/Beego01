package models

type User struct {
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Nick string `json:"nick"`
	Password string `json:"password"`
}
