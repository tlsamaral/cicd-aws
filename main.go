 package main

 import "net/http"

 func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
     	w.Write([]byte("CI/CD funcionando")
    })

    http.ListenAndServe(":8080", nil)
 }
