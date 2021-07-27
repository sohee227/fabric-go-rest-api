package resources

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"fabric-go-rest-api/hyperledger"
	ResourcesModel "fabric-go-rest-api/models/v1/resources"
)

func Show(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		rawResource, err := ResourcesModel.Show(clients, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(rawResource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
