package login

import (
	"net/http"
	"fmt"
	"encoding/json"

	// MY Package
	"github.com/saiful344/timer/helper/helper"
	"github.com/saiful344/timer/helper/auth"
	// Package Mysql
	 "github.com/jinzhu/gorm"
  	 // "github.com/gorilla/mux"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


type User struct{
	gorm.Model
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string{
	return "user"
}

type Message struct {
	Value string `json:"value"`
	Status string`json:"status"`
}

func dbconn()(db *gorm.DB){
  db, err := gorm.Open("mysql", "root:root@/timer?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
     panic(err.Error())
     fmt.Println("can't connect to db")

  }
  return db
}

func Sign_up(w http.ResponseWriter,r *http.Request){
	db := dbconn()
	defer db.Close()
	r.ParseForm()
	// Get value fromuser
	password := helper.HashAndSalt([]byte(r.FormValue("password")))
	username := r.FormValue("username")
	var data User
	data.Password = password
	data.Username = username
	db.NewRecord(data)
	db.Create(&data)
	w.WriteHeader(200)
	  output, err := json.Marshal(&Message{Value:"Success Added Data",Status:"200"})
	  if err != nil {
	      http.Error(w, err.Error(), 500)
	      return
	  }
	  w.Header().Set("content-type", "application/json")
	  w.Write(output)
}

func Login(w http.ResponseWriter,r *http.Request){
	db := dbconn()
	defer db.Close()
	db.AutoMigrate(User{})
	password := []byte(r.FormValue("password"))
	username := r.FormValue("username")
	var data User
  	db.Model(&data).Where("username = ?",username).Find(&data)
  	if(helper.ComparePasswords(data.Password, password)){
  		  auth.SetCookies(username,w)
  		  w.WriteHeader(200)
  		  output, _ := json.Marshal(&Message{Value:"Login Success",Status:"200"})
		  w.Header().Set("content-type", "application/json")
		  w.Write(output)
  	}else{
  		  w.WriteHeader(400)
  		  output, _ := json.Marshal(&Message{Value:"Login Failed",Status:"400"})
		  w.Header().Set("content-type", "application/json")
		  w.Write(output)
  	}
}

func Logout(w http.ResponseWriter,r *http.Request){
	auth.ClearCookie(w)
	  w.WriteHeader(200)
	  output, _ := json.Marshal(&Message{Value:"Logout Success",Status:"200"})
	  w.Header().Set("content-type", "application/json")
	  w.Write(output)
}

