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


type KieranController struct{}

func(controller KieranController) OnEnemyDetection(robot engine.Robot, enemyRobots []engine.Robot) {
	
	engine.StopRobot(robot)

	for _, enemy := range enemyRobots { engine.FireWeapon(robot, enemy) }
	
}

func(controller KieranController) OnMapEdgeWarning(robot engine.Robot) {}

func(controller KieranController) OnRobotCollision(robot engine.Robot) {}

func(controller KieranController) OnDamage(robot engine.Robot) {}


type CallumController struct{}

func(controller CallumController) OnEnemyDetection(robot engine.Robot, enemyRobots []engine.Robot) {
	
	engine.StopRobot(robot)

	for _, enemy := range enemyRobots { engine.FireWeapon(robot, enemy) }
	
}

func(controller CallumController) OnMapEdgeWarning(robot engine.Robot) {}

func(controller CallumController) OnRobotCollision(robot engine.Robot) {}

func(controller CallumController) OnDamage(robot engine.Robot) {}


type KarlController struct{}

func(controller KarlController) OnEnemyDetection(robot engine.Robot, enemyRobots []engine.Robot) {
	
	engine.StopRobot(robot)

	for _, enemy := range enemyRobots { engine.FireWeapon(robot, enemy) }
	
}

func(controller KarlController) OnMapEdgeWarning(robot engine.Robot) {}

func(controller KarlController) OnRobotCollision(robot engine.Robot) {}

func(controller KarlController) OnDamage(robot engine.Robot) {}





