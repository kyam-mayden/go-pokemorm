package middleware

import (
	"fmt"
	"net/http"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		user, pass, ok := request.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)
		if !ok || !checkUsernameAndPassword(user, pass) {
			responseWriter.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			responseWriter.WriteHeader(401)
			responseWriter.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(responseWriter, request)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "cheeseman" && password == "cheese-cheese-cheese"
}
