package main 

import (
	"Gin-Framework/routers"
)

func main() {
	var PORT = ":8080"

	routers.RegistRoutesServer().Run(PORT)
}