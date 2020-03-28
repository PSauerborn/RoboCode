package engine

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type event interface {
	apply(game GameState) GameState
	TickCost() int 
}

type DelayedEvent struct{ Delay int; Event event }

// ##################################################
// # Define interface used to assign damage to robots
// ##################################################

type DamageEvent struct { RobotName string; Damage int }

// Function used to define how a DamageEvent is applied
// to the GameState. Damage Events are used to subtract
// Durability from Robots
func (e DamageEvent) apply(game GameState) GameState {

	robots := []Robot{}

	// loop over robots to find target robot and subtract damage
	for _, robot := range game.Robots {

		// decrease damage from robot health
		if (robot.RobotName == e.RobotName) { robot.robotChasis.Durability -= e.Damage }

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

type FireWeaponEvent struct { robot, target Robot; weapon Weapon }

func (e FireWeaponEvent) apply(game GameState) GameState { return game }

func (e FireWeaponEvent) TickCost() int { return e.weapon.FireRate }

// ######################################
// # Define interface used to Move Robots
// ######################################

type MoveEvent struct { robot Robot; target Position }

func (e MoveEvent) apply(game GameState) GameState { 

	robots := []Robot{}

	// loop over robots to find target robot and subtract damage
	for _, robot := range game.Robots {

		if (robot.RobotName == e.robot.RobotName) { robot.robotPosition = e.target }

		robots = append(robots, robot)
	}

	game.Robots = robots

	return game
}

func (e MoveEvent) TickCost() int { return 0 }





