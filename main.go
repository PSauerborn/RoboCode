package main

import (
	"./engine"
	"./controllers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"./api"
)


func getGameConfig() (map[string]engine.RobotController, []map[string]string) {

	log.Debug("Getting Game configuration settings")

	// create config settings used to create robot
	robotConfigs := []map[string]string{
		{"name": "Kieran", "weapon": "Cannon", "chasis": "Ranger"},
		{"name": "Pascal", "weapon": "Cannon", "chasis": "Tank"},
	}

	// create map of handler functions
	robotControllers := map[string]engine.RobotController{
		"Kieran": engine.DefaultController{},
		"Pascal": controllers.PascalController{},
	}

	return robotControllers, robotConfigs
}

func main() {

	log.SetLevel(log.InfoLevel)

	// get router and start websocket server
	router := api.GetRouter()
	http.ListenAndServe(":10999", router)

	// handlers, configs := getGameConfig()

	// robots, err := engine.CreateRobotsSlice(configs, handlers)

	// if (err != nil) {
	// 	log.Fatal("Unable to Create robots with given configuration")

	// } else {
	// 	engine.RunGame(robots)
	// }
}