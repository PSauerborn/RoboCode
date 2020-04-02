package api


import (
	"../controllers"
	"../engine"
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func getGameConfig() (map[string]engine.RobotController, []map[string]string) {

	log.Debug("Getting Game configuration settings")

	// create config settings used to create robot
	robotConfigs := []map[string]string{
		{"name": "Kieran", "weapon": "Cannon", "chasis": "Tank"},
		{"name": "Pascal", "weapon": "Sniper", "chasis": "Tank"},
		{"name": "Callum", "weapon": "Rifle", "chasis": "Tank"},
		{"name": "Karl", "weapon": "Rifle", "chasis": "Tank"},
	}

	// create map of handler functions
	robotControllers := map[string]engine.RobotController{
		"Kieran": controllers.KieranController{},
		"Pascal": controllers.PascalController{},
		"Callum": controllers.CallumController{},
		"Karl": controllers.KarlController{},
	}

	return robotControllers, robotConfigs
}

func RunGameHandler(writer http.ResponseWriter, request *http.Request) {

	// upgrade connection to websocket type
	conn, err := upgradeHandler.Upgrade(writer, request, nil)

	if err != nil {
        log.Fatal(err)
        return
	}

	// get message from websocket
	_, p, err := conn.ReadMessage()

	if err != nil {
        log.Fatal(err)
        return
	}
	
	log.Info(fmt.Sprintf("Received Message %+v", p))

	// // convert message to bytes and return
	// err = conn.WriteMessage(websocket.TextMessage, []byte("Starting Robot Battle"))

	// get robot configuration settings and handlers from config
	robotControllers, robotConfigs := getGameConfig()

	// iterate over list of configs and generate robots
	robots, err := engine.CreateRobotsSlice(robotConfigs, robotControllers)

	if err != nil {
		log.Error("Unable To Create Robots with given Configuration")

	} else {
		log.Info("Successfully Created Robots. Starting Game Engine")

		engine.RunGame(robots, conn)
	}
}