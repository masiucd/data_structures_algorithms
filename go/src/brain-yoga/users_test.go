package brainyoga

import (
	"go-ds/src/lib/enums/gender"
	"strings"
	"testing"
)

func TestUserById(t *testing.T) {
	user := userById(2, users)
	if user == nil {
		t.Errorf("user with id 2 not found")
	}
	if user != nil && user.ID != 2 {
		t.Errorf("user id is not 2")
	}
}

func TestTransformedLadies(t *testing.T) {
	xs := transformedLadies(users)
	if len(xs) != 2 {
		t.Errorf("transformed ladies length is not 2")
	}
	for _, user := range xs {
		if user.Gender != gender.Female {
			t.Errorf("transform " + user.Name + " is not a lady")
		}
		if user.Name != strings.ToUpper(user.Name) {
			t.Errorf("transformed name is not uppercase")
		}
	}
}

func TestAccessibleUsers(t *testing.T) {
	ys := AccessibleUsers(users, bannedUsersIds)
	if len(ys) != 2 {
		t.Errorf("accessible users length is not 2")
	}
	for _, user := range ys {
		isBanned := false
		for _, bannedID := range bannedUsersIds {
			if user.ID == bannedID {
				isBanned = true
				break
			}
		}
		if isBanned {
			t.Errorf("user " + user.Name + " is banned")
		}
	}
}
