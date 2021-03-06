package resources

import (
	"encoding/json"
	"errors"
	"os"

	"fabric-go-rest-api/hyperledger"
	"fabric-go-rest-api/models"
)

func Show(clients *hyperledger.Clients, id string) (resource *models.Resource, err error) {

	resources := new(models.Resources)

	MSPID := os.Getenv("HYPERLEDGER_MSP_ID")
	if len(MSPID) == 0 {
		MSPID = "ibm"
	}

	res, err := clients.Query(MSPID, "mainchannel", "resources", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"id\": { \"$eq\":\"" + id + "\" } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, resources); err != nil {
		return
	}

	list := *resources

	if len(list) == 0 {
		err = errors.New("unable to find resource with id " + id)
		return
	}

	resource = &list[0]

	return
}
