package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	myHandler := HttpServerHandler{}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        loggingHandler(myHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type HttpServerHandler struct {
}

func (handler HttpServerHandler) ServeHTTP(reponse http.ResponseWriter, request *http.Request) {
	select {
	case <-time.After(5 * time.Second):
		reponse.Write([]byte("{'msg': 'Return before server timeout'}"))
	case <-request.Context().Done():
		fmt.Println("Cancel in backend")
		return
	}
	time.Sleep(10 * time.Second)
	reponse.Write([]byte("{'msg': 'Return after server timeout'}"))
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		t1 := time.Now()
		log.Println("Receiving incoming request -- log by interceptor")
		next.ServeHTTP(res, req)
		t2 := time.Now()
		log.Println("Finish request")
		log.Printf("%v", t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}
