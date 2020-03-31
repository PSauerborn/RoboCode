package controllers

import (
	"../engine"
)


type PascalController struct{}

func(controller PascalController) OnEnemyDetection(robot engine.Robot, enemyRobots []engine.Robot) {
	
	engine.StopRobot(robot)

	for _, enemy := range enemyRobots { engine.FireWeapon(robot, enemy) }
	
}

func(controller PascalController) OnMapEdgeWarning(robot engine.Robot) {}

func(controller PascalController) OnRobotCollision(robot engine.Robot) {}

func(controller PascalController) OnDamage(robot engine.Robot) {}


