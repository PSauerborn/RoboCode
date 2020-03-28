package api

import (
	"fmt"
	"../engine"
	log "github.com/sirupsen/logrus"
)


// Function used to generate a list of robots to play
// a game. Robots are created from a list of Map configuration
// settings, each of which contains the name, weapon and chasis
// types used to define the settings of the robot
func generateRobots(robotConfigs []map[string]string, robotHandlers map[string]engine.RobotExecutionHandler) ([]engine.Robot, error) {

	robots := []engine.Robot{}

	// loop over given configuration objects and convert to robots
	for _, config := range robotConfigs {

		// convert configuration to robot using map
		robot, err := engine.CreateRobot(config)

		if err != nil {
			log.Error("Cannot Create Robots with given Configuration")

			return []engine.Robot{}, err
			
		} else {
			log.Info(fmt.Sprintf("Successfully Created Robot %s", robot.RobotName))
			
			// inject handler method used to handle robot move
			robot.Execute = robotHandlers[robot.RobotName]
			
			// append robot to list of robots
			robots = append(robots, robot)
		}
	} 

	return robots, nil
}
