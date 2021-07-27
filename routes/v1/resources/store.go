package resources

import (
	"encoding/json"
	"net/http"

	"fabric-go-rest-api/models"

	"fabric-go-rest-api/hyperledger"
	ResourcesModel "fabric-go-rest-api/models/v1/resources"
)

func Store(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newResource, err := ResourcesModel.Store(
			clients,
			resource.Name,
			resource.ResourceTypeID,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(newResource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
