package routers

import (
	"JSON-Data/controllers"
	"fmt"
)

func RegistRoute () {
	controllers.EmployeeControllerDecode()

	fmt.Println()

	controllers.EmployeeControllerEncode()
}