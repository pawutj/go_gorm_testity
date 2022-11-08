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

func (executor *TruncateTableExecutor) TruncateTable(db string) error {

	command := "DELETE FROM " + db + ";"
	err := executor.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Exec(command)
	if err != nil {
		return err.Error
	}
	return err.Error
}
