package resources

import (
	"encoding/json"
	"net/http"

	"fabric-go-rest-api/hyperledger"
	ResourcesModel "fabric-go-rest-api/models/v1/resources"
)

func Index(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resources, err := ResourcesModel.Index(clients)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(resources)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
