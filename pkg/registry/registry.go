package registry

import (
	"github.com/mashnoor/basic-api-golang/pkg/database"
	"gorm.io/gorm"
)

//type GlobalRegistry struct {
//}

func GetDB() *gorm.DB {
	return database.GetDB()
}
