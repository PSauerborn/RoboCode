package engine

import (
	"fmt"
	"time"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// define struct used to keep track of the game state
type GameState struct{ Robots []Robot; EventQueue []event; DelayedEventQueue []DelayedEvent}

var robotGame GameState
var websocketConnection *websocket.Conn

// Function used to process event queue and filter
// queue into execute and delayed events
func processEventQueue(events []event) ([]event, []DelayedEvent) {

	executeEvents := []event{}
	delayedEvents := []DelayedEvent{}

	for _, e := range events {

		log.Info(fmt.Sprintf("Processing Event %#v", e))
		
		if e.TickCost() > 1 {
			log.Debug(fmt.Sprintf("Adding Event %+v to Delay Queue", e))

			// if tick cost is larger than 0, add event to delay slice
			delayedEvents = append(delayedEvents, DelayedEvent{Delay: e.TickCost(), Event: e})
		
			// else add to execution queue to execute immediately
		} else { executeEvents = append(executeEvents, e) }	
	}

	log.Debug(fmt.Sprintf("Returning Execute Queue %+v", executeEvents))
	
	return executeEvents, delayedEvents
}

func updateGameState() {

	// iterate over robots and execute move functions
	for _, robot := range robotGame.Robots { robot.Execute(robot) }

	// get events for immediate execution and for delayed execution from event queue
	executeEvents, delayedEvents := processEventQueue(robotGame.EventQueue)

	for _, e := range robotGame.DelayedEventQueue {

		// put event in execute list if tick delay has passed
		if e.Delay <= 1 {
			executeEvents = append(executeEvents, e.Event)
		// else lower tick delay by 1
		} else {
			delayedEvents = append(delayedEvents, DelayedEvent{Delay: e.Delay - 1, Event: e.Event})
		}
	}

	// iterate over execute events and update game board state
	for _, e := range executeEvents { 

		log.Debug(fmt.Sprintf("Sending Event %+v", e))

		robotGame = e.Apply(robotGame) 
	}

	robotGame.DelayedEventQueue = delayedEvents
	robotGame.EventQueue = []event{}

	// send game update to UI over socket
	// websocketConnection.WriteJSON(robotGame.Robots)

	log.Debug(fmt.Sprintf("Updated Game State to %+v", robotGame))
}

// Function used to run a game and update
// the game board based on robots and their moves
func RunGame(robots []Robot) {

	// set global variable to point at socket
	// websocketConnection = connection

	log.Info("Starting New Robot Game")

	// fill gamestate variables with new robots
	robotGame = GameState{Robots: robots, EventQueue: []event{}, DelayedEventQueue: []DelayedEvent{}}

	// Update Game state until all but one robot remain
	for len(robotGame.Robots) > 1 { updateGameState(); time.Sleep(1e8) }

	if len(robotGame.Robots) > 0 {
		winningRobot := robotGame.Robots[0]

		log.Info(fmt.Sprintf("Success! Robot %s has Won!", winningRobot.RobotName))
	}	
}