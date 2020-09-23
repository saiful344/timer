package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"github.com/saiful344/timer/models/auth"
)

func Handlefunc(){
	r := mux.NewRouter()
	r.HandleFunc("/",HomePage)
	r.HandleFunc("/home",auth.Redirect)
	r.HandleFunc("/login",auth.Login)
	r.HandleFunc("/logout",auth.Logout)
	log.Fatal(http.ListenAndServe(":9000",r))
}
func HomePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World")
}


func main(){
	fmt.Println("Server starta t port :9000")
	Handlefunc()
}
// https://itnext.io/encrypt-data-with-a-password-in-go-b5366384e291