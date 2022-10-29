package controllers

import (
	"fmt"
	"net/http"
  "github.com/ToffaKrtek/goCrud/pkg/models"
)

func GetAll(w http.ResponseWriter, r *http.Request){
  items := models.GetAll()
}

func Create(w http.ResponseWriter, r *http.Request){}
func GetById(w http.ResponseWriter, r *http.Request){}
func Delete(w http.ResponseWriter, r *http.Request){}
func Update(w http.ResponseWriter, r *http.Request){}

