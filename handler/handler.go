package handler

import (
	"github.com/MogLuiz/Gopportunities-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
