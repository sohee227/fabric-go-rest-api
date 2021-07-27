package resources

import (
	"net/http"

	"github.com/gorilla/mux"

	"fabric-go-rest-api/hyperledger"
	ResourcesModel "fabric-go-rest-api/models/v1/resources"
)

func Destroy(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := ResourcesModel.Destroy(clients, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
