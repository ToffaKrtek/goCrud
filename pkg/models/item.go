package models

import (
  "github.com/jinzhu/gorm"
)

var db *gorm.DB

type Item struct {
  gorm.Model
  Title string `gorm:"" json:"title"`  
  Body string `json:"body"`
}


func init(){}
func (i *Item) Create() *Item{
  db.NewRecord(i)
  db.Create(&i)
  return i
}

func GetAll() []Item{
  var Items []Item
  db.Find(&Items)
  return Items
}
