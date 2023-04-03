package common

import (
	"gorm.io/gorm/logger"
)

var LogLevelMap = map[string]logger.LogLevel{
	"info":   logger.Info,
	"warn":   logger.Warn,
	"error":  logger.Error,
	"silent": logger.Silent,
}
