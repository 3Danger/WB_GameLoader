package model

type Model struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsCustomer bool   `json:"is_customer"`

	Tasks []struct {
		Id     string  `json:"id"`
		Name   string  `json:"name"`
		Weight float32 `json:"weight"`
	} `json:"tasks"`

	JWT struct {
		RefreshToken string   `json:"refresh_token"`
		AcceptToken  []string `json:"accept_token"`
	} `json:"jwt"`
}
