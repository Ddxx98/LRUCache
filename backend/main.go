package main

import (
	"net/http"
	"fmt"
    "backend/router"
    "backend/middleware"
)



func main() {
	r := routers.Router()
	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", middleware.GetCorsConfig().Handler(r))
}
