package routers

import (
	"Go-API/controllers"
	"net/http"
)

func RegisterRouters() {
	http.HandleFunc("/getproduct", controllers.LogMiddleware(controllers.GetProduct))
	http.HandleFunc("/postproduct", controllers.LogMiddleware(controllers.PostProduct))
	http.HandleFunc("/updateproduct", controllers.Auth(controllers.LogMiddleware(controllers.PutProduct)))
	http.HandleFunc("/updateproduct/", controllers.LogMiddleware(controllers.PatchProduct))
	http.HandleFunc("/deleteproduct/", controllers.LogMiddleware(controllers.DeleteProduct))
}
