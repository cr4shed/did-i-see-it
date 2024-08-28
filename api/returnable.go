package main

import "github.com/cr4shed/did-i-see-it/data"

type Returnable interface {
	data.Collection | []data.Collection |
	data.MediaType | []data.MediaType |
	data.Media | []data.Media |
	data.View | []data.View
}