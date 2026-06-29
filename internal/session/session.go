package session

import "github.com/srajalnikhra/cli-login-system/internal/models"

var CurrentUser models.User
var IsLoggedIn bool

func Login(user models.User) {
	CurrentUser = user
	IsLoggedIn = true
}

func Logout() {
	CurrentUser = models.User{}
	IsLoggedIn = false
}

func GetCurrentUser() models.User {
	return CurrentUser
}
