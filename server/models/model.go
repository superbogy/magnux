package models

import "time"

type BaseStruct struct {
    id uint `gorm:"primary_key"`
    createdAt time.Time
    updatedAt time.Time
}

type Table struct {}

type Result struct {}

type Filter map[string]interface{}

func FindById(table string, id int, result BaseStruct) {
    DB.Table(table).First(&result)
}

func Get(table string, where Filter, result Result) {

}
