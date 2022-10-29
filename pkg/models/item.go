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

func GetById(Id int64) (*Item, *gorm.DB){
  var gItem Item 
  db:= db.Where("ID=?", Id).Find(&gItem)
  return &gItem, db
}

func Delete(Id int64) Item{
  var gItem Item 
  db.Where("ID=?", Id).Delete(gItem)
  return gItem
}
