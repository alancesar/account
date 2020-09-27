package infra

import "github.com/sirupsen/logrus"

func ConfigureLogger(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}
