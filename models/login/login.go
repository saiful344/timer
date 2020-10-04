package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// MY Package
	"github.com/saiful344/timer/helper/auth"
	"github.com/saiful344/timer/helper/helper"

	// Package Mysql
	"github.com/jinzhu/gorm"
	// "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

type Message struct {
	Value  string `json:"value"`
	Status string `json:"status"`
}

func dbconn() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "awc:root@/timer?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "root:root@/timer?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
		fmt.Println("can't connect to db")

	}
	return db
}

func Sign_up(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
	db := dbconn()
	defer db.Close()
	r.ParseForm()
	// Get value fromuser
	var password string
	var data User
	if len(r.FormValue("email")) != 0 {
		password = helper.HashAndSalt([]byte(r.FormValue("password")))
		username := r.FormValue("username")
		email := r.FormValue("email")
		data.Password = password
		data.Username = username
		data.Email = email
		fmt.Println("ok")
	} else {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		err = json.Unmarshal(b, &data)
		password = helper.HashAndSalt([]byte(data.Password))
		data.Password = password
	}
	var output []byte
	if len(data.Username) != 0 {
		db.NewRecord(data)
		db.Create(&data)
		w.WriteHeader(200)
		output = helper.Message_Json("Cie Berhasil Daftar")
	} else {
		output = helper.Message_Json("Cie Berhasil Kosong")
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
	db := dbconn()
	defer db.Close()
	var data User
	var forms User
	db.AutoMigrate(User{})
	var password []byte
	var username string
	if len(r.FormValue("username")) != 0 {
		password = []byte(r.FormValue("password"))
		username = r.FormValue("username")
	} else {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		err = json.Unmarshal(b, &forms)
		password = []byte(forms.Password)
		username = forms.Username
	}

	// code in below for compare password
	db.Model(&data).Where("username = ? OR email = ?", username, username).Find(&data)

	if len(data.Username) != 0 {
		fmt.Println(data)
		if helper.ComparePasswords(data.Password, password) {
			auth.SetCookies(username, w)
			w.WriteHeader(200)
			output, _ := json.Marshal(&Message{Value: "Login Success", Status: "200"})
			w.Header().Set("content-type", "application/json")
			w.Write(output)
		} else {
			w.WriteHeader(400)
			output, _ := json.Marshal(&Message{Value: "Login Failed", Status: "400"})
			w.Header().Set("content-type", "application/json")
			w.Write(output)
		}
	} else {
		w.WriteHeader(400)
		output, _ := json.Marshal(&Message{Value: "Login Failed", Status: "400"})
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	auth.ClearCookie(w)
	w.WriteHeader(200)
	output, _ := json.Marshal(&Message{Value: "Logout Success", Status: "200"})
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
