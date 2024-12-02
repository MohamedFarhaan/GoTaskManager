package services

import (
	"project2/config/logger"
)

func TaskStatusCheckCron() {
	logger.Trace("Entering TaskStatusCheckCron")
	InvalidateTasks()
}
