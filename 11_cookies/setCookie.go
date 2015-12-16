package main

import (
	"fmt"
	"net/http"
	"github.com/nu7hatch/gouuid"
)

func setCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + " email=robertsturner52@gmail.com" + " JSON data" + " Whatever",
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}

func main() {
	http.HandleFunc("/", setCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}