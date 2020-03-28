package api

import (
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiResponse struct {success bool; httpCode int}

// create upgrade handler use to upgrade HTTP connection to websocket
var upgradeHandler = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func GetRouter() *mux.Router {

	// define handler function used to check origins
	upgradeHandler.CheckOrigin = func(request *http.Request) bool { return true }

	router := mux.NewRouter()
	router.HandleFunc("/run", RunGameHandler).Methods("GET")

	return router
}