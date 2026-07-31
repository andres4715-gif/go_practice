package main

import (
	"github.com/sirupsen/logrus"
	"myProject/basic_learning/modules/hire"
	"myProject/basic_learning/modules/fire"
	"myProject/basic_learning/modules/sayHello"
)

func main() {
	sayHello.Hi()
	sayHello.GodBy()
	hire.Hire()
	fire.Fire()

	logger := logrus.New()

	logrus.Println("Hola desde logrus external package")
	logrus.Info("This is an info from logrus")
	logger.Info("This is an info from logger")
}
