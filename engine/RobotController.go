package engine 

import (
	
)

// define interface used for robot controllers
type RobotController interface {
	OnEnemyDetection(robot Robot, enemyRobots []Robot)
	OnMapEdgeWarning(robot Robot)
	OnDamage(robot Robot)
	OnRobotCollision(robot Robot)
}

type DefaultController struct {}

func(controller DefaultController) OnEnemyDetection(robot Robot, enemyRobots []Robot) {}

func(controller DefaultController) OnMapEdgeWarning(robot Robot) {}

func(controller DefaultController) OnRobotCollision(robot Robot) {}

func(controller DefaultController) OnDamage(robot Robot) {}

