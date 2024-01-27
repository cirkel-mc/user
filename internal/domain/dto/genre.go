package dto

type RequestGenre struct {
	Id int `param:"id"`
}

type ResponseGenre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
