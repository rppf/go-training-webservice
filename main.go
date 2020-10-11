package main

import (
	"net/http"

	"github.com/rppf/go-training-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
