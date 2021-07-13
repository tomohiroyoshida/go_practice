package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello world!!!")
	// t := template.Must(template.ParseFiles("tmpl.html"))

	// ファイルが複数で、２個目のファイルを実行したい場合
	// t, _ := template.ParseFiles("tmpl.html", "tmpl2.html")
	// t.ExecuteTemplate(w,"tmpl2.html", "Hello world!!!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
