package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

var (
	svr *http.Server
)

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {

	if svr == nil {
		mux := http.NewServeMux()
		mux.HandleFunc("/", Dispatch)
		svr = &http.Server{
			Addr:           fmt.Sprintf(":%d", 9990),
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
	}
	svr.ListenAndServe()
}

func Dispatch(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w)
	//r.ParseForm()
	//r.FormValue("name")
	name := r.FormValue("name")
	age := r.FormValue("age")

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(name + ":" + age)
}
