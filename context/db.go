package context

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	dsn := "user=postgres password=mongito2244 host=db.zdrljsxzqbboptonlqjq.supabase.co port=5432 dbname=postgres sslmode=disable TimeZone=Asia/Shanghai"
	initDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("la vida es una lenteja")
	}

	DB = initDB
	return nil
}
