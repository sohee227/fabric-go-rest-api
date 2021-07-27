package users

import "fabric-go-rest-api/models"

func Index() (users *models.Users, err error) {
	users = &mockUsers

	return
}
