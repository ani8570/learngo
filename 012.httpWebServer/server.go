package main

import "net/http"

// Handler interface
type Handler interface {
	ServerHTTP(http.ResponseWriter, *http.Request)
}
type testHandler struct {
	http.Handler
}

func (h *testHandler) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

func main() {

	// http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Write([]byte("Hello world"))
	// })

	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)

}
