package router

import (
	"mpg-tracker/api/controllers/health_controller"
)

func MakeHealthRoutes() Routes {
	controller := health_controller.HealthController{}
	return Routes{
		Route{Name: "HealthIndex", Method: "GET", Pattern: "/health", HandlerFunc: controller.Health},
	}
}
