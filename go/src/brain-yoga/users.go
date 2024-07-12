package brainyoga

import (
	"go-ds/src/lib/enums/gender"
	"strings"
)

type User struct {
	ID     int
	Name   string
	Email  string
	Gender gender.Gender
}

var users = []User{
	{ID: 1, Name: "John", Email: "j@ex.com", Gender: gender.Male},
	{ID: 2, Name: "Doe", Email: "doe@ex.com", Gender: gender.Male},
	{ID: 3, Name: "Jane", Email: "jane@ex.com", Gender: gender.Female},
	{ID: 4, Name: "Alice", Email: "alice@ex.com", Gender: gender.Female},
}

var bannedUsersIds []int = []int{3, 4}

func userById(id int, users []User) *User {
	var foundUser *User
	for _, user := range users {
		if user.ID == id {
			foundUser = &user
			return foundUser
		}
	}
	return foundUser
}

// like a filter and map
func transformedLadies(users []User) []User {
	var result []User
	for _, user := range users {
		if user.Gender == gender.Female {
			u := User{ID: user.ID, Name: strings.ToUpper(user.Name), Gender: user.Gender}
			result = append(result, u)
		}
	}
	return result
}

func AccessibleUsers(users []User, bannedUsersIds []int) []User {
	var result []User
	for _, user := range users {
		isBanned := false
		for _, bannedID := range bannedUsersIds {
			if user.ID == bannedID {
				isBanned = true
				break
			}
		}
		if !isBanned {
			result = append(result, user)
		}
	}
	return result
}
