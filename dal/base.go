package dal

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Base struct {
	gdb    *gorm.DB
	logger *zap.SugaredLogger
}
