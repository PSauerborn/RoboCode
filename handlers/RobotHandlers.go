package handlers

import (
	"../engine"
)


func KieranHandler(robot engine.Robot) {

}

func PascalHandler(robot engine.Robot) {

	enemyRobots := engine.ScanBattleField(robot)

	if len(enemyRobots) > 0 {

		for _, enemy := range enemyRobots { engine.FireWeapon(robot, enemy) }
	
	} else {
		engine.MoveRobot(robot, engine.GetRandomPosition())
	}

}

func CallumHandler(robot engine.Robot) {

}