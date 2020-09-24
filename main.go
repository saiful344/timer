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

// https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
// scure
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72