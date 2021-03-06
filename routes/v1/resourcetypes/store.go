package resourcetypes

import (
	"encoding/json"
	"net/http"

	"fabric-go-rest-api/models"

	ResourceTypesModel "fabric-go-rest-api/models/v1/resourcetypes"
)

func Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var resourcetype models.ResourceType
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&resourcetype)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newResourceType, err := ResourceTypesModel.Store(
			resourcetype.Name,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(newResourceType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
