package utils

import "github.com/sirupsen/logrus"

func HasError(err error, message string) bool {

	if err != nil {
		logrus.WithError(err).Error(message)
		return true
	}

	return false
}
