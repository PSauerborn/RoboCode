package api


import (
	"../handlers"
	"../engine"
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func getGameConfig() (map[string]engine.RobotExecutionHandler, []map[string]string) {

	// create config settings used to create robot
	robotConfigs := []map[string]string{
		{"name": "Kieran", "weapon": "Cannon", "chasis": "Ranger"},
		{"name": "Pascal", "weapon": "Cannon", "chasis": "Tank"},
	}

	// create map of handler functions
	robotHandlers := map[string]engine.RobotExecutionHandler{
		"Kieran": handlers.KieranHandler,
		"Pascal": handlers.PascalHandler,
	}

	return robotHandlers, robotConfigs
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

	// convert message to bytes and return
	err = conn.WriteMessage(websocket.TextMessage, []byte("Starting Robot Battle"))

	// get robot configuration settings and handlers from config
	robotHandlers, robotConfigs := getGameConfig()

	// iterate over list of configs and generate robots
	robots, err := generateRobots(robotConfigs, robotHandlers)

	if err != nil {
		log.Error("Unable To Create Robots with given Configuration")

	} else {
		log.Info("Successfully Created Robots. Starting Game Engine")

		engine.RunGame(robots, conn)
	}
}