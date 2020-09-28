package auth

import(
	"fmt"
	"net/http"
	"github.com/gorilla/securecookie"
)


var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))


func SetCookies(username string,w http.ResponseWriter){
	SetCookie(username,w)
}

func Redirect(w http.ResponseWriter,r *http.Request){
	userName := GetUserName(r)
    if !(len(userName) <= 0) {
    	fmt.Println("Goal")
    }else{
    	fmt.Println("failed")
    }
}

func Logout(response http.ResponseWriter, request *http.Request) {
    ClearCookie(response)
    http.Redirect(response, request, "/", 302)
}

func SetCookie(username string, w http.ResponseWriter){
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

func GetUserName(request *http.Request) (userName string) {
     if cookie, err := request.Cookie("session"); err == nil {
         cookieValue := make(map[string]string)
         if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
             userName = cookieValue["name"]
        }
     }
     return userName
 }
 
 func ClearCookie(response http.ResponseWriter) {
     cookie := &http.Cookie{
         Name:   "session",
         Value:  "",
         Path:   "/",
     }
     http.SetCookie(response, cookie)
 }
