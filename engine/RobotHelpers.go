package engine 

import (
	"math"
)

// Define function used to retrieve distance between 
// two robots based on their positions
func getDistance(robot Robot, target Robot) float64 {

	xDiff := robot.robotPosition.x - target.robotPosition.x
	yDiff := robot.robotPosition.y - target.robotPosition.y

	return math.Sqrt(math.Pow(xDiff, 2) - math.Pow(yDiff, 2))
}

// Helper function used to calculate the distane between two
// Position struct objects
func getPositionDistance(positionA, positionB Position) float64 {
	
	xDiff := positionA.x - positionB.x
	yDiff := positionA.y - positionB.y

	return math.Sqrt(math.Pow(xDiff, 2) - math.Pow(yDiff, 2))
}

// Helper function used to determine if a target robot
// is in range based on the weapon used by the hostile
// robot
func isInRange(robot Robot, target Robot) bool {

	distance := getDistance(robot, target)

	return distance < robot.robotWeapon.Range
}

// Helper function used to determine the tick price
// associated with moving across a board
func getTravelTime(robot Robot, distance float64) int {

	ticks := float64(robot.robotChasis.Weight) * (distance / 1000)

	return int(math.Floor(ticks))
}

func GetRandomPosition() Position { return generateStartingCoordinates() }

// Helper function used to determine if a robot has
// any move commands queued
func isMoving(robot Robot) bool { return false }

func isNearEdge(robot Robot) bool { return false }

func isOverEdge(robot Robot) bool { return false }
