package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(request.Method + "\n"))
		_, _ = writer.Write([]byte(request.RequestURI + "\n"))
		_, _ = writer.Write([]byte(time.Now().String() + "\n"))
	})

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
