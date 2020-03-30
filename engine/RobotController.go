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

