package main

import (
	"github.com/sirupsen/logrus"
	"myProject/basic_learning/modules/contratar"
	"myProject/basic_learning/modules/despedir"
	"myProject/basic_learning/modules/saludar"
)

func main() {
	saludar.Hola()
	saludar.Adios()
	contratar.Contratar()
	despedir.Despedir()

	logger := logrus.New()

	logrus.Println("Hola desde logrus external package")
	logrus.Info("This is an info from logrus")
	logger.Info("This is an info from logger")
}
