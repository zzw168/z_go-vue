package initialize

import (
	"server/global"
	"server/model"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		model.Todo{},
	)
	if err != nil {
		return err
	}
	return nil
}
