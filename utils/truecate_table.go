package utils

import (
	"gorm.io/gorm"
)

type TruncateTableExecutor struct {
	db *gorm.DB
}

func InitTruncateTableExecutor(db *gorm.DB) TruncateTableExecutor {
	return TruncateTableExecutor{db}
}

func (executor *TruncateTableExecutor) TruncateTable() error {

	err := executor.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec("DELETE FROM products")
	if err != nil {
		panic(err)
	}
	return err.Error
}
