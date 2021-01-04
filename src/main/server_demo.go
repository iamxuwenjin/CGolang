package main

import (
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("/Users/wenjin.xu/Golang/CGolang/src/main/index.html")
	_, _ = w.Write(content)
}

//func main() {
//	http.HandleFunc("/index", index)
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//}
