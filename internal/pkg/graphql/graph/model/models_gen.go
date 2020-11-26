// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Account struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Token    *string `json:"token"`
}

type Ad struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
	Price int    `json:"price"`
}

type AdRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
	Price int    `json:"price"`
}
