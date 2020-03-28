package main

import (
	"./api"
	"net/http"
	log "github.com/sirupsen/logrus"
)


func main() {

	log.SetLevel(log.InfoLevel)

	// get router and start websocket server
	router := api.GetRouter()
	http.ListenAndServe(":8844", router)
}