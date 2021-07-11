package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	v := struct {
		Msg string `json: msg`
	} {
		Msg: "hello, world!",
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error", err)
	}
}


func main() {
	// http.HandleFunc("/", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// parameter
		fmt.Fprintln(w, "hello", "Gopherさんの運勢は" + r.FormValue("luck") + "です！")
		// conetnt-type
		contentType := r.Header.Get("Content-Type")
		fmt.Println(w,"content-type: ",  contentType)
		p := "request received!"
		// json
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&p); err != nil {
			log.Println("err", err)
		}
		fmt.Println(p)

	})
	http.ListenAndServe(":8080", nil)
}