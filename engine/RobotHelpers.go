package engine 

import (
	"math"
	"math/rand"
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

// Write function used to determine if aa shot has hit a target
func shotHitTarget(robot, target Robot, weapon Weapon) bool {

	// get distance between robots
	distance := getDistance(robot, target)

	// calculate distance penalty to shot accuracy
	distancePenalty := (distance / weapon.Range) * weapon.Accuracy

	return weapon.Accuracy > (rand.Float64() + distancePenalty)
}

// Helper function used to determine the tick price
// associated with moving across a board
func getTravelTime(robot Robot, distance float64) int {

	ticks := float64(robot.robotChasis.Weight) * (distance / 100)

	return int(math.Floor(ticks))
}

func GetRandomPosition() Position { return generateStartingCoordinates() }

// Helper function used to determine if a robot has
// any move commands queued
func isMoving(robot Robot) bool { 

	// define condition used to filter events based on event source
	conditionDelayed := func(e DelayedEvent) bool { return (e.Event.EventSource() == robot.RobotName && e.Event.EventType() == "Move") }
	conditionExecutable := func(e event) bool { return (e.EventSource() == robot.RobotName && e.EventType() == "Move") }


	// determine if robot has queued events that match moving event types
	hasDelayed := len(filterDelayedEvents(robotGame.DelayedEventQueue, conditionDelayed)) != 0
	hasExecutable := len(filterEvents(robotGame.EventQueue, conditionExecutable)) != 0

	return hasDelayed && hasExecutable
}

func robotIsDestroyed(robot string) bool {
	
	for _, bot := range robotGame.Destroyed { 
		if bot.RobotName == robot { return true }
	}

	return false
}

func isNearEdge(robot Robot) bool { return false }

func isOverEdge(robot Robot) bool { return false }

// Define helper function used to filter robots based on arbitrary condition
func filterRobots(robots []Robot , condition func(robot Robot) bool) (filteredArray []Robot) {

	for _, robot := range robots {
		if condition(robot) { filteredArray = append(filteredArray, robot) }
	}
	return
}

// Define helper functions used to filter events based on arbitrary condition
func filterEvents(events []event , condition func(e event) bool) (filteredEvents []event) {

	for _, e := range events {
		if condition(e) { filteredEvents = append(filteredEvents, e) }
	}
	return
}

// Define helper functions used to filter delayed events based on arbitrary condition
func filterDelayedEvents(events []DelayedEvent , condition func(e DelayedEvent) bool) (filteredEvents []DelayedEvent) {

	for _, e := range events {
		if condition(e) { filteredEvents = append(filteredEvents, e) }
	}
	return
}

// define helper function used to shuffle event queue
func shuffleEvents(events []event) []event {

	// define function used to shuffle values
	shuffle := func(i, j int) { events[i], events[j] = events[j], events[i] }

	// shuffle events in slice
	rand.Shuffle(len(events), shuffle)

	return events
}