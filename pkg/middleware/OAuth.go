package middleware

import (
	"fmt"
	"net/http"
)

func OAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		apiKey := request.Header.Get("x-api-key")
		fmt.Println("api key: ", apiKey)
		if !checkAuthCredentials(apiKey) {
			responseWriter.Header().Set("OAuth problem", "Please provide a valid API key")
			responseWriter.WriteHeader(401)
			responseWriter.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(responseWriter, request)
	}
}

func checkAuthCredentials(apiKey string) bool {
	return apiKey == "cheese-cheese-cheese"
}
