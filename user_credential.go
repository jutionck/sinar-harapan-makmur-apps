package main

type UserCredential struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
}
