package main

type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
	Native		string
}

func main() {
	u1 := User {
		ID:			1,
		FirstName: 	"Jerome",
		LastName: 	"Wilson",
		Native: 	"Kerala",
	}

	u2 := User {
		ID:			2,
		FirstName:	"Ajins",
		LastName: 	"G. Puthuran",
		Native: 	"Kerala",
	}

	if u1.ID == u2.ID {
		println("It's the same user")
	} else if(u1.FirstName == u2.FirstName) {
		println("It's a similar user")
	} else {
		println("Both are different users")
	}

	if(u1.Native == u2.Native) {
		println("User",u1.FirstName,u1.LastName,"and",u2.FirstName,u2.LastName,"are from same place -- ",u1.Native)
	}

}