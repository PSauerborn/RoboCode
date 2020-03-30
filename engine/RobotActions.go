package engine

import (
	"fmt"
	"math/rand"
	log "github.com/sirupsen/logrus"
	"errors"
)

// function used to emit a game event into the 
// Game Event Queue
func emitEvent(robot Robot, e event) {

	log.Debug(fmt.Sprintf("Emitting Event %#v", e))

	robotGame.EventQueue = append(robotGame.EventQueue, e)
}

func emitDelayedEvent(robot Robot, e DelayedEvent) {

	log.Debug(fmt.Sprintf("Emitting Delayed Event %#v", e))

	robotGame.DelayedEventQueue = append(robotGame.DelayedEventQueue, e)
}

// Function used to move a robot
func MoveRobot(robot Robot, target Position) {

	// calculate total distance that robot needs to move over entire event
	xDiff, yDiff := target.x - robot.robotPosition.x, target.y - robot.robotPosition.y

	// get total number of ticks required to 
	tickCount := getTravelTime(robot, getPositionDistance(robot.robotPosition, target))

	// get increment that needs to be added at each tick
	xIncrement, yIncrement := xDiff / float64(tickCount), yDiff / float64(tickCount)

	for i := 1; i < tickCount + 1; i++ { 

		// created delayed move event for later execution
		targetPosition := Position{x: robot.robotPosition.x + (float64(i) * xIncrement), y: robot.robotPosition.y + (float64(i) * yIncrement)}

		event := DelayedEvent{Delay: i, Event: MoveEvent{SourceRobot: robot, target: targetPosition}}

		emitDelayedEvent(robot, event) 
	}
}

// Function used to fire a weapon at a target 
// robot. Note that whether or not weapons hit
// is based on chance and weapon accuracy
func FireWeapon(robot Robot, target Robot) (bool, error) {

	log.Info(fmt.Sprintf("Robot %s Firing at %s", robot.RobotName, target.RobotName))

	// return false and error if target robot is not in range
	if !isInRange(robot, target) {
		return false, errors.New("Target Robot not in Range")
	}

	// determine whether or not shot hit based on accuracy
	if (robot.robotWeapon.Accuracy > rand.Float64()) {

		log.Debug("Weapon Shot Successfully. Emitting events")

		// emit event used to fire weapon
		fireEvent := FireWeaponEvent{SourceRobot: robot, target: target, weapon: robot.robotWeapon}

		emitEvent(robot, fireEvent)

		// emit event used to damage target robot
		damageEvent := DamageEvent{SourceRobot: robot, TargetRobot: target, Damage: robot.robotWeapon.Damage}

		emitDelayedEvent(robot, DelayedEvent{Delay: robot.robotWeapon.FireRate, Event: damageEvent})

		return true, nil
	}

	log.Debug("Weapon Shot Missed")

	return false, errors.New("Missed Target Robot")
}

func ScanBattleField(robot Robot) []Robot {

	robots := []Robot{}

	// loop over game board and find robots
	for _, inGameRobot := range robotGame.Robots {

		// if enemy robot is in range of scanner, add enemy to list
		if (inGameRobot.RobotName != robot.RobotName && getDistance(robot, inGameRobot) < 20) {
			robots = append(robots, inGameRobot)
		}
	}

	return robots
}

// define function used to roam robot across battlefield
func Roam(robot Robot) { if !isMoving(robot) { MoveRobot(robot, GetRandomPosition()) }} 

func ClearCommands(robot Robot) { 

	// filter out events based on robot name
	events := []event{}

	for _, event := range robotGame.EventQueue {
		if !(event.EventSource() == robot.RobotName) { events = append(events, event) }
	}

	robotGame.EventQueue = events

	// filter delayed events on robot name
	delayedEvents := []DelayedEvent{}

	for _, event := range robotGame.DelayedEventQueue {
		if !(event.Event.EventSource() == robot.RobotName) { delayedEvents = append(delayedEvents, event) }
	}

	robotGame.DelayedEventQueue = delayedEvents
}