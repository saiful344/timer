package timer

import (
	"net/http"
	"fmt"
	_ "encoding/json"

	// MY Package
	 _ "github.com/saiful344/timer/helper/helper"

	// Package Mysql
	 "github.com/jinzhu/gorm"
  	 // "github.com/gorilla/mux"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Timer struct{
	Id string `json:"id"`
	Id_user string `json:"id_user"`
	date string `json:"date"`
	time_start string `json:"time_start"`
	time_run string `json:"time_run"`
	time_end string `json:"time_end"`
	id_project string `json:"id_project"`
}


func (Timer) TableName() string{
	return "timer"
}

func dbconn()(db *gorm.DB){
  db, err := gorm.Open("mysql", "root:root@/timer?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
     panic(err.Error())
     fmt.Println("can't connect to db")

  }
  return db
}

func Create(w http.ResponseWriter, r *http.Request){
	db := dbconn()
	defer db.Close()
	r.ParseForm()
}

func Create_project(w http.ResponseWriter, r *http.Request){
	
}

// var datetime = time.Now()
// datetime.Format(time.RFC3339)