package common

import (
	"bill_serve/global"
	"bill_serve/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{
			Logger: global.MysqlLog,
		})
	}

	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认排序
	}

	count = DB.Select("id").Find(&list).RowsAffected
	if count > 0 {
		offset := option.Limit * (option.Page - 1)
		if offset <= 0 {
			offset = 0
		}
		err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	}
	return list, count, err
}