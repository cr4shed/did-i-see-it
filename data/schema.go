package data

type Media struct {
	Id		  		int
	MediaTypeId 	int
	Title			string
}

type MediaType struct {
	Id		  		int
	Title			string
}

type View struct {
	Id		  		int
	CollectionId	int
	MediaId		 	int
}

type Collection struct {
	Id		  		int
	Name			string
}

type User struct {
	Id int
	Username string
	Email string
	PassHash string
}