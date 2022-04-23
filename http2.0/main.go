package main

import (
	"html/template"
	"log"
	"net/http"
)

// go http
func main() {
	svr := &http.Server{
		Addr:    ":9093",
		Handler: nil,
	}
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../static/js/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../static/img/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHTML := template.Must(template.ParseGlob("../static/html/index.html"))
		log.Println("remote addr:", r.RemoteAddr)
		log.Println("header:", r.Header)
		log.Println("response header:", w.Header())
		tmplHTML.Execute(w, nil)
	})
	svr.ListenAndServeTLS("server.crt", "server.key")
}
