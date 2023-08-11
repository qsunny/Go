package global

import (
	"base-showcase/config"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
	EWA_DB     *gorm.DB
)
