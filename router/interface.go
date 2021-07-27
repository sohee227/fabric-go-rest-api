package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"fabric-go-rest-api/hyperledger"
	V1Router "fabric-go-rest-api/routes/v1"
)

const (
	staticDir = "/static/"
)

type Service interface {
	GetRawRouter() *mux.Router
}

func GetRouter() Service {
	r := Router{
		RawRouter: mux.NewRouter().StrictSlash(true),
	}

	configPath := os.Getenv("HYPERLEDGER_CONFIG_PATH")
	MSPID := os.Getenv("HYPERLEDGER_MSP_ID")
	if len(MSPID) == 0 {
		MSPID = "ibm"
	}

	if len(configPath) == 0 {
		panic("ENV var 'HYPERLEDGER_CONFIG_PATH' is not set. unable to connect to network")
	}

	fmt.Println("configPath")
	fmt.Println(configPath)
	fmt.Println("MSPID")
	fmt.Println(MSPID)

	clients := hyperledger.NewClientMap(
		"test-network",
		configPath,
	)
	fmt.Println("clients")
	fmt.Println(clients)

	_, err := clients.AddClient(
		"Admin",
		MSPID,
		"mychannel",
	)
	if err != nil {
		panic(err)
	}

	r.RawRouter.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	for _, route := range GetRoutes() {
		r.RawRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	for name, pack := range V1Router.GetRoutes(clients) {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return r
}
