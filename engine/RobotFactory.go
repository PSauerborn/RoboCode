package engine

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"errors"
	log "github.com/sirupsen/logrus"
)

// define struct used to represent coordinates
type Position struct{ x, y float64 }

// define struct used to define a robot chasis
type Chasis struct{ Durability, Weight int }

// define strict used to define robot weapons
type Weapon struct{
	Damage, FireRate int 
	Accuracy, Range float64
}

// define struct used to define robot
type Robot struct{
	RobotName string
	robotPosition Position
	robotChasis Chasis
	robotWeapon Weapon
	controller RobotController
}

// Function used to generate a random set of starting
// coordinates based on environment variables set
func generateStartingCoordinates() Position {

	// get combat zone length from environment variables
	CombatZoneLength, err := strconv.ParseFloat(os.Getenv("COMBAT_ZONE_LENGTH"), 64)

	// set default value if value has not been specified in environs
	if (err != nil || CombatZoneLength == 0) { CombatZoneLength = 100 }

	return Position{x: rand.Float64() * CombatZoneLength, y: rand.Float64() * CombatZoneLength}
}


// define mapping used to relate string inputs to chasis types
var allowedChasisTypes = map[string]Chasis{
	"Tank": Chasis{Durability: 100, Weight: 100},
	"Ranger": Chasis{Durability: 50, Weight: 60},
	"Buggy": Chasis{Durability: 20, Weight: 10},
}                                                            

// Function used to create the chasis for a robot. Chasis
// must be of type "Tank", "Ranger" or "Buggy"
func generateRobotChasis(ChasisType string) (Chasis, error) {

	log.Info(fmt.Sprintf("Creating Robot with Chasis %s", ChasisType))

	// return chasis if chasis type is valid
	if chasis, ok := allowedChasisTypes[ChasisType]; ok {return chasis, nil }

	return Chasis{}, errors.New("Invalid Chasis Configuration")
}


// define mapping used to relate string inputs to chasis types
var allowedWeaponTypes = map[string]Weapon{
	"Cannon": Weapon{Damage: 100, Range: 50, FireRate: 3, Accuracy: 0.3},
	"Sniper": Weapon{Damage: 50, Range: 100, FireRate: 2, Accuracy: 0.9},
	"Rifle": Weapon{Damage: 20, Range: 60, FireRate: 1, Accuracy: 0.7},
}         

// Function used to create the weapon for a robot. Weapons
// must be of type "Cannon", "Sniper" or "Rifle"
func generateRobotWeapon(WeaponType string) (Weapon, error) {

	log.Info(fmt.Sprintf("Creating Robot with Weapon %s", WeaponType))

	// return weapon if specified weapon type is valid
	if weapon, ok := allowedWeaponTypes[WeaponType]; ok { return weapon, nil }

	return Weapon{}, errors.New("Invalid Weapon Configuration")
}

// Function used to create Robot object using a map object.
// Maps contain the configuration settings used to define
// the robot, including Name, Weapon and Chasis
func CreateRobot(config map[string]string) (Robot, error) {

	// create robot chasis using configuration settings
	chasis, err := generateRobotChasis(config["chasis"])

	if err != nil {
		log.Error(err)

		return Robot{}, err
	}

	// create weapon using configuration settings
	weapon, err := generateRobotWeapon(config["weapon"])

	if err != nil {
		log.Error(err)

		return Robot{}, err
	}

	// create robot using the weapon and chasis generated from configuration settings
	robot := Robot{ 
		RobotName: config["name"], 
		robotPosition: generateStartingCoordinates(), 
		robotChasis: chasis, 
		robotWeapon: weapon,
		controller: DefaultController{},
	}

	return robot, nil
}

// Alternate creation function used to create robots that are stored in maps
// instead fo slices, which allows greater flexibility during game time. Robots
// are created from a list of configuration maps that can be parsed from JSON
// body sent with the request
func CreateRobotsDict(robotConfigs []map[string]string, robotControllers map[string]RobotController) (map[string]Robot, error) {

	robots := map[string]Robot{}

	// loop over given configuration objects and convert to robots
	for _, config := range robotConfigs {

		// convert configuration to robot using map
		robot, err := CreateRobot(config)

		if err != nil {
			log.Error("Cannot Create Robots with given Configuration")

			return map[string]Robot{}, err
			
		} else {
			log.Info(fmt.Sprintf("Successfully Created Robot %s", robot.RobotName))
			
			// inject handler method used to handle robot move
			robot.controller = robotControllers[robot.RobotName]
			
			// append robot to list of robots
			robots[config["name"]] = robot
		}
	} 

	return robots, nil
}

// Function used to generate a list of robots to play
// a game. Robots are created from a list of Map configuration
// settings, each of which contains the name, weapon and chasis
// types used to define the settings of the robot
func CreateRobotsSlice(robotConfigs []map[string]string, robotControllers map[string]RobotController) ([]Robot, error) {

	robots := []Robot{}

	// loop over given configuration objects and convert to robots
	for _, config := range robotConfigs {

		// convert configuration to robot using map
		robot, err := CreateRobot(config)

		if err != nil {
			log.Error("Cannot Create Robots with given Configuration")

			return []Robot{}, err
			
		} else {
			log.Info(fmt.Sprintf("Successfully Created Robot %s", robot.RobotName))
			
			// inject handler method used to handle robot move
			robot.controller = robotControllers[robot.RobotName]

			// append robot to list of robots
			robots = append(robots, robot)
		}
	} 

	return robots, nil
}