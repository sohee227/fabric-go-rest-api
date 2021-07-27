package router

import (
	"fabric-go-rest-api/models"
	HomeHandler "fabric-go-rest-api/routes/home"
	StatusHandler "fabric-go-rest-api/routes/status"
)

func GetRoutes() models.Routes {
	return models.Routes{
		models.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: HomeHandler.Index},
		models.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: StatusHandler.Index},
	}
}
