package main

import (
	"fmt"
	"time"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saiful344/timer/helper/auth"
	"github.com/saiful344/timer/models/login"
	"github.com/saiful344/timer/models/timer"
)

func Handlefunc() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage)
	// r.HandleFunc("/home",auth.Redirect)
	r.HandleFunc("/create", timer.Create)
	r.HandleFunc("/login", login.Login)
	r.HandleFunc("/sign", login.Sign_up)
	r.HandleFunc("/logout", auth.Logout)
	r.HandleFunc("/msg", timer.Create_project)
	r.HandleFunc("/create_project", timer.Create_project)
	log.Fatal(http.ListenAndServe(":9000", r))
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	fmt.Println(userName)
}

func main() {
	fmt.Println("Server starta t port :9000")
	Handlefunc()

	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	// time := time.Now()
	fmt.Println(datetime)
}

// https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
// scure
// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
