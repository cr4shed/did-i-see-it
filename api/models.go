package main

import "github.com/cr4shed/did-i-see-it/data"

type Returnable interface {
	data.Collection | []data.Collection |
	data.MediaType | []data.MediaType |
	data.Media | []data.Media |
	data.View | []data.View |
	IdResposne
}

type UserDto struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IdResposne struct {
	Id int
}