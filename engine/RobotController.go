package engine 

import (
	
)

// define interface used for robot controllers
type RobotController interface {
	OnEnemyDetection(robot Robot)
	OnWallCollision(robot Robot)
	OnDamage(robot Robot)
	OnRobotCollision(robot Robot)
}

type DefaultController struct {}

func(controller DefaultController) OnEnemyDetection(robot Robot) {}

func(controller DefaultController) OnWallCollision(robot Robot) {}

func(controller DefaultController) OnRobotCollision(robot Robot) {}

func(controller DefaultController) OnDamage(robot Robot) {}

// Generic handler function called each tick one every robot.
func tickHandler(robot Robot) {

	// scan battle field for robots and emit event
	for _, enemyRobot := range ScanBattleField(robot) { robot.controller.OnEnemyDetection(enemyRobot) }

	Roam(robot)
}

// Define function used to execute game ticks
func executeGameTick() { for _, robot := range robotGame.Robots { tickHandler(robot) }}