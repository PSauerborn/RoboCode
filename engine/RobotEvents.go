package engine

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type event interface {
	Apply(game GameState) GameState
	TickCost() int 
	EventSource() string
	EventType() string
}

type DelayedEvent struct{ Delay int; Event event }

// ##################################################
// # Define interface used to assign damage to robots
// ##################################################

type DamageEvent struct { SourceRobot, TargetRobot Robot; Damage int }

func (e DamageEvent) EventSource() string { return e.SourceRobot.RobotName }

func (e DamageEvent) EventType() string { return "Damage" }
 
// Function used to define how a DamageEvent is applied
// to the GameState. Damage Events are used to subtract
// Durability from Robots
func (e DamageEvent) Apply(game GameState) GameState {

	robots := []Robot{}

	// loop over robots to find target robot and subtract damage
	for _, robot := range game.Robots {

		// decrease damage from robot health
		if (robot.RobotName == e.TargetRobot.RobotName) { robot.robotChasis.Durability -= e.Damage }

		if (robot.robotChasis.Durability > 0) {
			robots = append(robots, robot)
		} else {
			log.Info(fmt.Sprintf("Destroyed Robot %s", robot.RobotName))
		}
	}

	game.Robots = robots

	return game
}

func (e DamageEvent) TickCost() int { return 0 }

// #######################################
// # Define interface used to Fire Weapons
// #######################################

type FireWeaponEvent struct { SourceRobot, target Robot; weapon Weapon }

func (e FireWeaponEvent) EventSource() string { return e.SourceRobot.RobotName }

func (e FireWeaponEvent) EventType() string { return "FireWeapon" }

func (e FireWeaponEvent) Apply(game GameState) GameState { return game }

func (e FireWeaponEvent) TickCost() int { return e.weapon.FireRate }

// ######################################
// # Define interface used to Move Robots
// ######################################

type MoveEvent struct { SourceRobot Robot; target Position }

func (e MoveEvent) EventSource() string { return e.SourceRobot.RobotName }

func (e MoveEvent) EventType() string { return "Move" }

func (e MoveEvent) Apply(game GameState) GameState { 

	robots := []Robot{}

	// loop over robots to find target robot and subtract damage
	for _, robot := range game.Robots {

		if (robot.RobotName == e.SourceRobot.RobotName) { robot.robotPosition = e.target }

		// fire wall collision event if robot is close to edge of map
		if isNearEdge(robot) { robot.controller.OnWallCollision(robot) }

		// remove robot if gone over edge
		if !isOverEdge(robot) { robots = append(robots, robot) } 
	}

	game.Robots = robots

	return game
}

func (e MoveEvent) TickCost() int { return 0 }





