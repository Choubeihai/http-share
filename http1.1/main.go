package main

import (
	"html/template"
	"log"
	"net/http"
)

// go http
func main() {
	svr := &http.Server{
		Addr: ":9092",
	}
	svr.SetKeepAlivesEnabled(true)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../static/js/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../static/img/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHTML := template.Must(template.ParseGlob("../static/html/index.html"))
		tmplHTML.Execute(w, nil)
		log.Println("remote addr:", r.RemoteAddr)
		log.Println("header:", r.Header)
		log.Println(w.Header())

	})
	svr.ListenAndServe()

}

//// gin http
//func main() {
//	router := gin.Default()
//	router.LoadHTMLGlob("../static/html/*")
//	router.Static("/img", "../static/img")
//	router.Static("/js", "../static/js")
//	router.GET("/", Index)
//	router.Run(":9091")
//}
//
//func Index(c *gin.Context) {
//	c.HTML(http.StatusOK, "index-back2.html", nil)
//}
