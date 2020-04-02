package engine

import (

)

type WebsocketEvent struct{ EventType, EventMessage, EventSource string }

type GamestateEvent struct{ EventType string; Robots []RobotUpdate }

type RobotPosition struct{PositionX, PositionY float64 }

type RobotUpdate struct{ Name string; Health int; Coordinates RobotPosition }

// Define function used to generate robot update messages
// that can be Marshalled into JSON format to send to frontend
func createUpdate(robots []Robot) (updates []RobotUpdate) {
	
	for _, robot := range robots {

		update := RobotUpdate{ 
			Name: robot.RobotName, 
			Health: robot.robotChasis.Durability,  
			Coordinates: RobotPosition{PositionX: robot.robotPosition.x, PositionY: robot.robotPosition.y},
		}

		updates = append(updates, update)
	}
	return
}