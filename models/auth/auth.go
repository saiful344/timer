package auth

import(
	"net/http"
	"github.com/gorilla/securecookie"
)

// http://www.cihanozhan.com/building-login-and-register-application-with-golang/
var cookieHandler = securecookie new(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32)),
)

func Login(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != ""{
		// logic
		if bool {
			SetCookie(name,w)
			redirectTarget = "/home"
		} else{
			redirectTarget = "/daftar"
		}

	}

	http.Redirect(w,r,redirectTarget,302)
}

func Register(w http.ResponseWriter,r *http.Request){

}

func Logout(w http.ResponseWriter, r *http.Request){
	clearSession(r)
	http.Redirect(w,r,"/",302)
}

func setSession(username string, w http.ResponseWriter){
	value := map[string]string{
		"name" : username,
	}
	encoded, err := cookieHandler.Encode("session",value)
	if err == nil{
		cookie := &http.Cookie{
			Name: "session",
			Value : encoded,
			Path: "/",
		}
		http.SetCookie(w, cookie)
	}
}

func getUserName(request *http.Request) (userName string) {
     if cookie, err := request.Cookie("session"); err == nil {
         cookieValue := make(map[string]string)
         if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
             userName = cookieValue["name"]
        }
     }
     return userName
 }
 
 func clearSession(response http.ResponseWriter) {
     cookie := &http.Cookie{
         Name:   "session",
         Value:  "",
         Path:   "/",
         MaxAge: -1,
     }
     http.SetCookie(response, cookie)
 }

 func GetUserName(r *http.Request) (username string){
 	cookie,err := r.Cookie("cookie"); 
 	if err == nil{
 		cookieValue := mak
 	}
 }