package main

import (
	"fmt"
	"net/http"
)


func hanlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>my it library</h1>")
}
func main() {
	port := "8088"
	http.HandleFunc("/", hanlerFunc)
	fmt.Printf("Сервер стартовал на порту %s", port)
	
	http.ListenAndServe(":8088", nil)
}


