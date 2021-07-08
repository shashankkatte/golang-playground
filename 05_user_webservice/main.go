package main

import (
	"net/http"

	"github.com/shashankkatte/golang-playground/05_user_webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
