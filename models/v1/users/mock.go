package users

import (
	"fabric-go-rest-api/models"
)

var mockUsers models.Users

func init() {
	usr, _ := models.NewUser("Nick", "Kotenberg", "nick@mail.com", "1234")

	mockUsers = models.Users{
		*usr,
	}
}
