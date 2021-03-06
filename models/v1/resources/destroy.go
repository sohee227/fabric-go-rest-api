package resources

import (
	"fabric-go-rest-api/hyperledger"
)

func Destroy(clients *hyperledger.Clients, id string) error {
	res, err := Show(clients, id)
	if err != nil {
		return err
	}

	res.Active = false

	if _, err = Update(clients, id, res, &UpdateOpts{Replace: true}); err != nil {
		return err
	}

	return nil
}
