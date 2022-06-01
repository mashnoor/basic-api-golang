package models

//
//import (
//	"github.com/mashnoor/basic-api-golang/pkg/registry"
//	"log"
//)
//
//type BaseModel interface {
//	Insert() error
//}
//
//func (b *BaseModel) Save(model interface{}) {
//	reg := registry.GlobalRegistry{}
//	err = reg.GetDB().Create(&model).Error
//	if err != nil {
//		log.Fatal("Couldn't create user")
//	}
//}
