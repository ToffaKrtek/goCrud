package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
  db *gorm.DB
)

func Connect(){
  dsn := "host=localhost user=gocrud password=gocrud_pass dbname=gocrud port=5432 sslmode=disable TimeZone=ALMST"
  d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  db = d
}
func GetDB() *gorm.DB{
  return db
}
