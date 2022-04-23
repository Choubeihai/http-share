package main

import (
	"html/template"
	"log"
	"net/http"
)

// go http
func main() {
	svr := &http.Server{
		Addr: ":9091",
	}

	svr.SetKeepAlivesEnabled(false)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../static/js/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../static/img/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHTML := template.Must(template.ParseGlob("../static/html/index.html"))
		//w.Header().Set("Content-Length", "2049000")
		log.Println("remote addr:", r.RemoteAddr)
		log.Println("header:", r.Header)
		log.Println("response header:", w.Header())
		tmplHTML.Execute(w, nil)
	})
	svr.ListenAndServe()

}

//// gin http
//func main() {
//	router := gin.Default()
//	server := &http.Server{
//		Addr:    ":9092",
//		Handler: router,
//	}
//	router.LoadHTMLGlob("../static/html/*")
//	router.Static("/img", "../static/img")
//	router.Static("/js", "../static/js")
//	router.GET("/", Index)
//	server.SetKeepAlivesEnabled(false)
//	server.ListenAndServe()
//
//}
//
//func Index(c *gin.Context) {
//	c.HTML(http.StatusOK, "index-back2.html", nil)
//}
