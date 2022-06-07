package router

import (
	"mpg-tracker/api/controllers/fill_up_controller"
	"mpg-tracker/api/repositories/fill_up_repository"
)

func MakeFillUpRoutes(repo *fill_up_repository.FillUpRepository) Routes {
	controller := fill_up_controller.FillUpController{Repository: repo}
	return Routes{
		Route{Name: "PostFillUp", Method: "POST", Pattern: "/fillups", HandlerFunc: controller.PostFillUp},
		Route{Name: "GetFillUps", Method: "GET", Pattern: "/fillups", HandlerFunc: controller.GetListFillUps},
		Route{Name: "GetFillUpById", Method: "GET", Pattern: "/fillups/{fillupId}", HandlerFunc: controller.GetFillUpById},
		Route{Name: "PutFillUp", Method: "PUT", Pattern: "/fillups/{fillupId}", HandlerFunc: controller.PutFillUp},
		Route{Name: "DeleteFillUpById", Method: "DELETE", Pattern: "/fillups/{fillupId}", HandlerFunc: controller.DeleteFillUpById},
	}
}
