package models

import (
    "github.com/jinzhu/gorm"
)

import _ "github.com/go-sql-driver/mysql"


var DB *gorm.DB


func InitConnect(conn string) (db *gorm.DB, err error) {
    db, err = gorm.Open("mysql", conn)
    db.LogMode(true)
    DB = db
    return
}
