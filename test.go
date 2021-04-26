package main

import (
	"fmt"
	"net/http"
	"time"
)



func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, "hello.html")
	})
	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprint(w, "About Page")
    })
    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprint(w, "Contact Page") 
	})
	
	http.HandleFunc("/postform", func (w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("username")
        age := r.FormValue("userage")
         
        fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
	})

	
	http.HandleFunc("/hello", greet)
	fmt.Println(" 🚀 Server is listening on port 5000...")
	http.ListenAndServe(":5000", nil)
}


