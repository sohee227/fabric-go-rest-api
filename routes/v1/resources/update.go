package resources

import (
	"encoding/json"
	"net/http"

	"fabric-go-rest-api/models"

	"github.com/gorilla/mux"

	"fabric-go-rest-api/hyperledger"
	ResourcesModel "fabric-go-rest-api/models/v1/resources"
)

func Update(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts ResourcesModel.UpdateOpts
		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&resource); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updatedResource, err := ResourcesModel.Update(clients, id, &resource, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updatedResource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
